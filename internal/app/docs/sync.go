// Package docs syncs the RobusTest documentation repository at runtime and
// renders it under /docs. The docs repo (izinga/robustest_documentation_md)
// stays the single source of truth: the site downloads GitHub's tarball on
// startup and on an interval, so pushes to the repo go live without a site
// deploy and nothing is ever copied into this codebase.
package docs

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const defaultRepo = "izinga/robustest_documentation_md"

// Store holds the synced docs tree and rendered-page cache.
type Store struct {
	mu       sync.RWMutex
	dir      string // directory containing the current synced tree
	sha      string // commit SHA of the current tree
	syncedAt time.Time
	lastErr  error

	repo   string
	branch string
	token  string
	root   string // parent dir under which synced trees are extracted
	local  string // non-empty: serve a bundled docs dir, never call GitHub

	pageCache sync.Map // key string -> *Page (invalidated on new SHA)
	navCache  *Nav
}

// NewStore configures the docs store from the environment.
func NewStore() *Store {
	repo := os.Getenv("DOCS_REPO")
	if repo == "" {
		repo = defaultRepo
	}
	branch := os.Getenv("DOCS_BRANCH")
	if branch == "" {
		branch = "main"
	}
	root := os.Getenv("DOCS_DIR")
	if root == "" {
		root = "./data/docs"
	}
	// Preferred mode: docs bundled at build time (make docs-fetch) and
	// shipped with the deploy — no GitHub credentials on the server.
	local := os.Getenv("DOCS_LOCAL_DIR")
	if local == "" {
		local = "./docs-content"
	}
	if fi, err := os.Stat(local); err != nil || !fi.IsDir() {
		local = ""
	}
	return &Store{
		repo:   repo,
		branch: branch,
		token:  os.Getenv("DOCS_GITHUB_TOKEN"),
		root:   root,
		local:  local,
	}
}

// Start performs an initial sync. Publishing is a manual step from there:
// the docs repo is private, so there is no background polling — after
// pushing docs, hit /docs/refresh (or `make docs-refresh`). Periodic
// polling can be opted into by setting DOCS_SYNC_INTERVAL (e.g. "10m").
func (s *Store) Start() {
	if err := s.Sync(); err != nil {
		log.Printf("docs: initial sync failed (refresh manually via /docs/refresh): %v", err)
	}
	if s.local != "" {
		log.Printf("docs: serving bundled content from %s", s.local)
		return
	}
	v := os.Getenv("DOCS_SYNC_INTERVAL")
	if v == "" {
		log.Printf("docs: periodic sync disabled; publish via /docs/refresh")
		return
	}
	d, err := time.ParseDuration(v)
	if err != nil || d < time.Minute {
		log.Printf("docs: invalid DOCS_SYNC_INTERVAL %q; periodic sync disabled", v)
		return
	}
	go func() {
		for range time.Tick(d) {
			if err := s.Sync(); err != nil {
				log.Printf("docs: periodic sync failed: %v", err)
			}
		}
	}()
}

// Ready reports whether a docs tree is available to serve.
func (s *Store) Ready() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.dir != ""
}

// Status returns the current sync state for the refresh endpoint.
func (s *Store) Status() (sha string, syncedAt time.Time, lastErr error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.sha, s.syncedAt, s.lastErr
}

// Dir returns the current synced tree directory (empty until first sync).
func (s *Store) Dir() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.dir
}

func (s *Store) apiRequest(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	if s.token != "" {
		req.Header.Set("Authorization", "Bearer "+s.token)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "robustest-web-docs")
	client := &http.Client{Timeout: 60 * time.Second}
	return client.Do(req)
}

// headSHA asks GitHub for the branch tip so unchanged trees skip the download.
func (s *Store) headSHA() (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/commits/%s", s.repo, s.branch)
	resp, err := s.apiRequest(http.MethodGet, url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return "", fmt.Errorf("github commits API: %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}
	var out struct {
		SHA string `json:"sha"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	return out.SHA, nil
}

// Sync refreshes the served tree: in bundled mode it re-fingerprints the
// local directory (picking up docs shipped by a deploy or docs-publish);
// otherwise it fetches the repo tarball from GitHub.
func (s *Store) Sync() error {
	if s.local != "" {
		return s.syncLocal()
	}
	s.mu.Lock()
	current := s.sha
	s.mu.Unlock()

	sha, err := s.headSHA()
	if err != nil {
		s.recordErr(err)
		return err
	}
	if sha == current {
		s.recordErr(nil)
		return nil
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/tarball/%s", s.repo, sha)
	resp, err := s.apiRequest(http.MethodGet, url)
	if err != nil {
		s.recordErr(err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		err := fmt.Errorf("github tarball API: %s: %s", resp.Status, strings.TrimSpace(string(body)))
		s.recordErr(err)
		return err
	}

	dest := filepath.Join(s.root, sha[:12])
	if err := extractTarball(resp.Body, dest); err != nil {
		os.RemoveAll(dest)
		s.recordErr(err)
		return err
	}

	s.mu.Lock()
	old := s.dir
	s.dir = dest
	s.sha = sha
	s.syncedAt = time.Now()
	s.lastErr = nil
	s.navCache = nil
	s.mu.Unlock()
	s.pageCache = sync.Map{}

	if old != "" && old != dest {
		os.RemoveAll(old)
	}
	log.Printf("docs: synced %s@%s", s.repo, sha[:12])
	return nil
}

// syncLocal fingerprints the bundled docs dir so caches invalidate when a
// deploy or docs-publish replaces the content.
func (s *Store) syncLocal() error {
	h := sha256.New()
	err := filepath.WalkDir(s.local, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		fmt.Fprintf(h, "%s|%d|%d;", p, info.Size(), info.ModTime().UnixNano())
		return nil
	})
	if err != nil {
		s.recordErr(err)
		return err
	}
	sha := hex.EncodeToString(h.Sum(nil))

	s.mu.Lock()
	changed := sha != s.sha
	s.dir = s.local
	s.sha = sha
	s.syncedAt = time.Now()
	s.lastErr = nil
	if changed {
		s.navCache = nil
	}
	s.mu.Unlock()
	if changed {
		s.pageCache = sync.Map{}
	}
	return nil
}

func (s *Store) recordErr(err error) {
	s.mu.Lock()
	s.lastErr = err
	s.mu.Unlock()
}

// extractTarball unpacks a GitHub tarball (which nests everything under a
// single top-level directory) into dest, stripping that first component.
func extractTarball(r io.Reader, dest string) error {
	if err := os.MkdirAll(dest, 0o755); err != nil {
		return err
	}
	gz, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gz.Close()
	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		parts := strings.SplitN(hdr.Name, "/", 2)
		if len(parts) < 2 || parts[1] == "" {
			continue
		}
		rel := filepath.Clean(parts[1])
		if rel == "." || strings.HasPrefix(rel, "..") {
			continue
		}
		target := filepath.Join(dest, rel)
		if !strings.HasPrefix(target, filepath.Clean(dest)+string(os.PathSeparator)) {
			continue // path traversal guard
		}
		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return err
			}
			f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
			if err != nil {
				return err
			}
			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				return err
			}
			f.Close()
		}
	}
}
