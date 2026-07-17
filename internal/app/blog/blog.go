// Package blog renders the posts under blog-content/ as pages beneath /blog.
//
// Unlike docs (synced from a separate repo), blog content lives in this repo:
// posts are reviewed alongside the site that publishes them, so a claim and the
// page making it move together.
package blog

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

// Category codes echo the instrument tags used across the site (LIVE, PERF,
// TV…). They are the post's subject, not a taxonomy for its own sake.
const (
	CatWhy      = "WHY"      // why we built it this way
	CatField    = "FIELD"    // what broke, and what it cost us
	CatBusiness = "BUSINESS" // the industry, the bet, the market
	CatPerf     = "PERF"
	CatTV       = "TV"
	CatLab      = "LAB"
)

// Post is one rendered blog post.
type Post struct {
	Slug        string
	Title       string
	Description string
	Category    string
	Author          string
	AuthorRole      string // job title only, e.g. "Co-founder"
	AuthorURL       string // the person's profile — the byline icon links here
	AuthorImage     string // headshot path; falls back to the initials monogram if empty
	AuthorCompany   string // e.g. "RobusTest" — rendered after the role
	AuthorCompanyURL string // the company page — the company name links here
	Date        time.Time
	Updated     time.Time // zero unless the post has been revised
	Draft       bool
	ReadingTime int // minutes, computed — never author-supplied
	WordCount   int
	Content     template.HTML
	TOC         []TOCItem
	Sources     []Source
	// Illustration is an inline schematic SVG loaded from
	// blog-content/illustrations/<slug>.svg, if present. Inlined (not <img>)
	// so it inherits the page's theme variables and recolours in light/dark.
	Illustration template.HTML
}

// TOCItem is one H2 in a post.
type TOCItem struct {
	ID   string
	Text string
}

// Source is a citation listed in the post's frontmatter. Every factual claim
// about a competitor or about our own code is expected to carry one.
type Source struct {
	Title string `yaml:"title"`
	URL   string `yaml:"url"`
	Note  string `yaml:"note"`
}

// Year groups posts for the index ledger.
type Year struct {
	Year  int
	Posts []*Post
}

type frontmatter struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Category    string   `yaml:"category"`
	Author          string `yaml:"author"`
	AuthorRole      string `yaml:"author_role"`
	AuthorURL       string `yaml:"author_url"`
	AuthorImage     string `yaml:"author_image"`
	AuthorCompany   string `yaml:"author_company"`
	AuthorCompanyURL string `yaml:"author_company_url"`
	Date        string   `yaml:"date"`
	Updated     string   `yaml:"updated"`
	Draft       bool     `yaml:"draft"`
	Sources     []Source `yaml:"sources"`
}

// Store loads posts from a directory and caches the parsed result.
type Store struct {
	mu    sync.RWMutex
	dir   string
	posts []*Post
	bySlug map[string]*Post
	loaded bool
}

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
	goldmark.WithParserOptions(parser.WithAutoHeadingID()),
	goldmark.WithRendererOptions(html.WithUnsafe()), // team-authored, reviewed in PR
)

// NewStore reads posts from dir (default "blog-content").
func NewStore(dir string) *Store {
	if dir == "" {
		dir = "blog-content"
	}
	return &Store{dir: dir, bySlug: map[string]*Post{}}
}

// Load parses every post. Called once at startup; Reload picks up edits.
func (s *Store) Load() error {
	entries, err := filepath.Glob(filepath.Join(s.dir, "*.md"))
	if err != nil {
		return err
	}

	posts := make([]*Post, 0, len(entries))
	bySlug := make(map[string]*Post, len(entries))

	for _, path := range entries {
		slug := strings.TrimSuffix(filepath.Base(path), ".md")
		if strings.HasPrefix(slug, "_") {
			continue // _template.md and friends
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("blog: read %s: %w", path, err)
		}
		p, err := parsePost(slug, raw)
		if err != nil {
			return fmt.Errorf("blog: parse %s: %w", path, err)
		}
		if p.Draft {
			continue
		}
		if svg, err := os.ReadFile(filepath.Join(s.dir, "illustrations", slug+".svg")); err == nil {
			p.Illustration = template.HTML(svg)
		}
		posts = append(posts, p)
		bySlug[slug] = p
	}

	// Newest first — the ledger reads top-down like a log.
	sort.Slice(posts, func(i, j int) bool { return posts[i].Date.After(posts[j].Date) })

	s.mu.Lock()
	s.posts, s.bySlug, s.loaded = posts, bySlug, true
	s.mu.Unlock()
	return nil
}

// Reload re-reads from disk.
func (s *Store) Reload() error { return s.Load() }

// Posts returns every published post, newest first.
func (s *Store) Posts() []*Post {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.posts
}

// ByYear groups published posts into descending years for the index.
func (s *Store) ByYear() []Year {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var years []Year
	for _, p := range s.posts {
		y := p.Date.Year()
		if n := len(years); n > 0 && years[n-1].Year == y {
			years[n-1].Posts = append(years[n-1].Posts, p)
			continue
		}
		years = append(years, Year{Year: y, Posts: []*Post{p}})
	}
	return years
}

// Get returns one post by slug.
func (s *Store) Get(slug string) (*Post, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	p, ok := s.bySlug[slug]
	return p, ok
}

// Related returns up to n other posts, preferring the same category.
func (s *Store) Related(slug string, n int) []*Post {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cur, ok := s.bySlug[slug]
	if !ok {
		return nil
	}
	var same, other []*Post
	for _, p := range s.posts {
		if p.Slug == slug {
			continue
		}
		if p.Category == cur.Category {
			same = append(same, p)
		} else {
			other = append(other, p)
		}
	}
	out := append(same, other...)
	if len(out) > n {
		out = out[:n]
	}
	return out
}

var fmDelim = regexp.MustCompile(`(?s)^---\r?\n(.*?)\r?\n---\r?\n?`)

func parsePost(slug string, raw []byte) (*Post, error) {
	m := fmDelim.FindSubmatch(raw)
	if m == nil {
		return nil, fmt.Errorf("missing YAML frontmatter")
	}
	var fm frontmatter
	if err := yaml.Unmarshal(m[1], &fm); err != nil {
		return nil, fmt.Errorf("frontmatter: %w", err)
	}
	if fm.Title == "" {
		return nil, fmt.Errorf("frontmatter: title is required")
	}
	if fm.Date == "" {
		return nil, fmt.Errorf("frontmatter: date is required")
	}
	date, err := time.Parse("2006-01-02", fm.Date)
	if err != nil {
		return nil, fmt.Errorf("frontmatter: date %q: want YYYY-MM-DD", fm.Date)
	}
	var updated time.Time
	if fm.Updated != "" {
		updated, err = time.Parse("2006-01-02", fm.Updated)
		if err != nil {
			return nil, fmt.Errorf("frontmatter: updated %q: want YYYY-MM-DD", fm.Updated)
		}
	}

	body := raw[len(m[0]):]

	var buf bytes.Buffer
	if err := md.Convert(body, &buf); err != nil {
		return nil, err
	}

	author := fm.Author
	if author == "" {
		author = "RobusTest"
	}
	category := strings.ToUpper(fm.Category)
	if category == "" {
		category = CatField
	}

	words := len(strings.Fields(string(body)))
	minutes := words / 200
	if minutes < 1 {
		minutes = 1
	}

	return &Post{
		Slug:        slug,
		Title:       fm.Title,
		Description: fm.Description,
		Category:    category,
		Author:      author,
		AuthorRole:  fm.AuthorRole,
		AuthorURL:   fm.AuthorURL,
		AuthorImage: fm.AuthorImage,
		AuthorCompany:    fm.AuthorCompany,
		AuthorCompanyURL: fm.AuthorCompanyURL,
		Date:        date,
		Updated:     updated,
		Draft:       fm.Draft,
		WordCount:   words,
		ReadingTime: minutes,
		Content:     template.HTML(buf.String()),
		TOC:         extractTOC(body),
		Sources:     fm.Sources,
	}, nil
}

var (
	h2Line     = regexp.MustCompile(`(?m)^##\s+(.+?)\s*$`)
	nonIDChars = regexp.MustCompile(`[^a-z0-9\- ]`)
	mdInline   = regexp.MustCompile("[*_`]")
	fenceLine  = regexp.MustCompile("(?m)^```")
)

// extractTOC pulls H2s for the reading rail. Headings inside fenced code
// blocks are skipped — a shell comment is not a section.
func extractTOC(body []byte) []TOCItem {
	var out []TOCItem
	inFence := false
	for _, line := range strings.Split(string(body), "\n") {
		if fenceLine.MatchString(line) {
			inFence = !inFence
			continue
		}
		if inFence {
			continue
		}
		m := h2Line.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		text := stripMD(m[1])
		out = append(out, TOCItem{ID: headingID(text), Text: text})
	}
	return out
}

// headingID mirrors goldmark's auto heading IDs closely enough for anchors.
func headingID(text string) string {
	s := strings.ToLower(strings.TrimSpace(text))
	s = nonIDChars.ReplaceAllString(s, "")
	return strings.ReplaceAll(s, " ", "-")
}

func stripMD(s string) string {
	return strings.TrimSpace(mdInline.ReplaceAllString(s, ""))
}
