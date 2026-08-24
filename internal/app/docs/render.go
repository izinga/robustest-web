package docs

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// Page is a rendered documentation page.
type Page struct {
	Title       string
	Path        string // URL path under /docs, e.g. "admin/healthpage"
	Description string // meta description, derived from the first paragraph
	Content     template.HTML
	TOC         []TOCItem
}

// TOCItem is one heading in a page's table of contents.
type TOCItem struct {
	Level int // 2 or 3
	ID    string
	Text  string
}

// Nav is the sidebar tree parsed from _sidebar.md.
type Nav struct {
	Sections []NavSection
}

// NavSection groups links under a bold heading in _sidebar.md.
type NavSection struct {
	Title string
	Links []NavLink
}

// NavLink is one sidebar entry.
type NavLink struct {
	Title string
	Path  string // URL path under /docs ("" = docs home)
}

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
	goldmark.WithParserOptions(parser.WithAutoHeadingID()),
	goldmark.WithRendererOptions(html.WithUnsafe()), // docs repo is trusted, team-authored
)

// Nav parses and caches the sidebar for the current synced tree.
func (s *Store) Nav() *Nav {
	s.mu.RLock()
	cached := s.navCache
	dir := s.dir
	s.mu.RUnlock()
	if cached != nil {
		return cached
	}
	if dir == "" {
		return &Nav{}
	}
	nav := parseSidebar(filepath.Join(dir, "_sidebar.md"))
	s.mu.Lock()
	s.navCache = nav
	s.mu.Unlock()
	return nav
}

var sidebarLink = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)

func parseSidebar(path string) *Nav {
	raw, err := os.ReadFile(path)
	if err != nil {
		return &Nav{}
	}
	nav := &Nav{}
	var current *NavSection
	for _, line := range strings.Split(string(raw), "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		// Section headings look like: * **Getting Started**
		if strings.HasPrefix(trimmed, "*") && strings.Contains(trimmed, "**") {
			title := strings.TrimSpace(strings.Trim(strings.TrimLeft(trimmed, "*- "), "*"))
			nav.Sections = append(nav.Sections, NavSection{Title: title})
			current = &nav.Sections[len(nav.Sections)-1]
			continue
		}
		m := sidebarLink.FindStringSubmatch(trimmed)
		if m == nil {
			continue
		}
		link := NavLink{Title: m[1], Path: docPathFromLink(m[2])}
		if current == nil {
			nav.Sections = append(nav.Sections, NavSection{})
			current = &nav.Sections[len(nav.Sections)-1]
		}
		current.Links = append(current.Links, link)
	}
	return nav
}

// absLinkRe matches targets that already address somewhere else entirely
// (http://, https://, mailto:, protocol-relative //host). These must never be
// rewritten into /docs routes.
var absLinkRe = regexp.MustCompile(`^(?:[a-zA-Z][a-zA-Z0-9+.-]*:|//)`)

// docPathFromLink converts a sidebar/markdown link target to a /docs URL path.
// It returns "" for the docs root and for anything that is not a repo-relative
// page reference.
func docPathFromLink(target string) string {
	target = strings.TrimSpace(target)
	if target == "" || absLinkRe.MatchString(target) {
		return ""
	}
	target = strings.TrimPrefix(target, "/")
	target = strings.TrimSuffix(target, ".md")
	// Collapse "./" and "../" segments so a stray relative prefix cannot leak a
	// literal ".." into the URL. path.Clean keeps leading ".." on escaping
	// paths, so drop those explicitly rather than emitting a broken link.
	target = path.Clean(target)
	for strings.HasPrefix(target, "../") {
		target = strings.TrimPrefix(target, "../")
	}
	if target == "." || target == ".." || target == "" || target == "README" {
		return ""
	}
	return target
}

// IndexEntry is one searchable page in the docs index.
type IndexEntry struct {
	Title    string   `json:"title"`
	Section  string   `json:"section"`
	Path     string   `json:"path"`
	Headings []string `json:"headings,omitempty"`
}

// Index returns the search index for the current tree: every sidebar page
// with its section and headings. Cached per synced SHA.
func (s *Store) Index() []IndexEntry {
	s.mu.RLock()
	sha := s.sha
	s.mu.RUnlock()
	if v, ok := s.pageCache.Load("index|" + sha); ok {
		return v.([]IndexEntry)
	}
	var idx []IndexEntry
	for _, section := range s.Nav().Sections {
		for _, link := range section.Links {
			entry := IndexEntry{Title: link.Title, Section: section.Title, Path: link.Path}
			if page, err := s.Load(link.Path); err == nil {
				for _, t := range page.TOC {
					entry.Headings = append(entry.Headings, t.Text)
				}
			}
			idx = append(idx, entry)
		}
	}
	s.pageCache.Store("index|"+sha, idx)
	return idx
}

// PrevNext returns the sidebar neighbours of a page, for footer navigation.
func (s *Store) PrevNext(path string) (prev, next *NavLink) {
	var flat []NavLink
	for _, section := range s.Nav().Sections {
		flat = append(flat, section.Links...)
	}
	for i, link := range flat {
		if link.Path == path {
			if i > 0 {
				p := flat[i-1]
				prev = &p
			}
			if i < len(flat)-1 {
				n := flat[i+1]
				next = &n
			}
			return prev, next
		}
	}
	return nil, nil
}

// Load returns the rendered page for a /docs URL path ("" = home/README).
func (s *Store) Load(urlPath string) (*Page, error) {
	s.mu.RLock()
	dir, sha := s.dir, s.sha
	s.mu.RUnlock()
	if dir == "" {
		return nil, fmt.Errorf("docs not synced yet")
	}

	urlPath = strings.Trim(urlPath, "/")
	if strings.Contains(urlPath, "..") || strings.HasPrefix(urlPath, "_") {
		return nil, os.ErrNotExist
	}
	cacheKey := sha + "|" + urlPath
	if v, ok := s.pageCache.Load(cacheKey); ok {
		return v.(*Page), nil
	}

	file, ok := resolveFile(dir, urlPath)
	if !ok {
		return nil, os.ErrNotExist
	}
	raw, err := os.ReadFile(file)
	if err != nil {
		return nil, os.ErrNotExist
	}

	page, err := renderPage(raw, urlPath)
	if err != nil {
		return nil, err
	}
	s.pageCache.Store(cacheKey, page)
	return page, nil
}

// resolveFile maps a docs URL path to the markdown file backing it, following
// the same rule Load uses: "a/b" is a/b.md, falling back to a/b/README.md for
// section links, and "" is the tree README.
func resolveFile(dir, urlPath string) (string, bool) {
	rel := strings.Trim(urlPath, "/")
	if rel == "" {
		rel = "README"
	}
	file := filepath.Join(dir, rel+".md")
	if _, err := os.Stat(file); err == nil {
		return file, true
	}
	alt := filepath.Join(dir, rel, "README.md")
	if _, err := os.Stat(alt); err == nil {
		return alt, true
	}
	return "", false
}

// ModTime is when the markdown behind a docs page last changed. The docs tree
// is extracted from a GitHub tarball, which stamps every entry with its commit
// time, so this is the content's real date and stays put across re-syncs of the
// same commit — unlike the sync timestamp, which moves every time we poll.
func (s *Store) ModTime(urlPath string) (time.Time, bool) {
	s.mu.RLock()
	dir := s.dir
	s.mu.RUnlock()
	if dir == "" {
		return time.Time{}, false
	}
	if strings.Contains(urlPath, "..") || strings.HasPrefix(strings.Trim(urlPath, "/"), "_") {
		return time.Time{}, false
	}
	file, ok := resolveFile(dir, urlPath)
	if !ok {
		return time.Time{}, false
	}
	fi, err := os.Stat(file)
	if err != nil {
		return time.Time{}, false
	}
	return fi.ModTime(), true
}

var (
	headingRe = regexp.MustCompile(`(?m)^(#{1,3})\s+(.+?)\s*$`)
	mdLinkRe  = regexp.MustCompile(`\]\(([^)#]+\.md)(#[^)]*)?\)`)
	imgSrcRe  = regexp.MustCompile(`(src="|\]\()(?:\.\./|\./)*(assets/[^")]+)`)
)

func renderPage(raw []byte, urlPath string) (*Page, error) {
	src := string(raw)

	// Rewrite .md links to /docs routes. The repo follows the docsify
	// convention of repo-root-relative targets ("guides/x.md").
	src = mdLinkRe.ReplaceAllStringFunc(src, func(m string) string {
		parts := mdLinkRe.FindStringSubmatch(m)
		target, frag := strings.TrimPrefix(parts[1], "./"), parts[2]
		// External targets that merely happen to end in ".md" are left alone;
		// rewriting them would produce "/docs/https://...".
		if absLinkRe.MatchString(strings.TrimSpace(target)) {
			return m
		}
		p := docPathFromLink(target)
		if p == "" {
			return "](/docs" + frag + ")"
		}
		return "](/docs/" + p + frag + ")"
	})

	// Rewrite asset references to the synced-assets route.
	src = imgSrcRe.ReplaceAllString(src, `$1/docs/$2`)

	// Title = first h1; TOC = h2/h3.
	title := ""
	var toc []TOCItem
	for _, m := range headingRe.FindAllStringSubmatch(src, -1) {
		level := len(m[1])
		text := strings.TrimSpace(m[2])
		if level == 1 && title == "" {
			title = stripMD(text)
			continue
		}
		if level == 2 || level == 3 {
			toc = append(toc, TOCItem{Level: level, ID: headingID(text), Text: stripMD(text)})
		}
	}
	if title == "" {
		title = "Documentation"
	}

	var buf bytes.Buffer
	if err := md.Convert([]byte(src), &buf); err != nil {
		return nil, err
	}
	return &Page{
		Title:       title,
		Path:        urlPath,
		Description: metaDescription(src, title),
		Content:     template.HTML(buf.String()),
		TOC:         toc,
	}, nil
}

// metaDescription derives a unique meta description from the page's first
// prose paragraph, falling back to the title when the page opens with
// non-prose (a table of contents, a code block, an image).
func metaDescription(src, title string) string {
	const maxLen = 155

	var para []string
	inFence := false
	for _, line := range strings.Split(src, "\n") {
		t := strings.TrimSpace(line)
		if strings.HasPrefix(t, "```") || strings.HasPrefix(t, "~~~") {
			inFence = !inFence
			continue
		}
		if inFence {
			continue
		}
		// A blank line ends the first paragraph once we have collected some.
		if t == "" {
			if len(para) > 0 {
				break
			}
			continue
		}
		// Skip headings, list items, tables, quotes, images and raw HTML —
		// none of them read as a summary.
		if strings.HasPrefix(t, "#") || strings.HasPrefix(t, ">") ||
			strings.HasPrefix(t, "|") || strings.HasPrefix(t, "<") ||
			strings.HasPrefix(t, "!") || strings.HasPrefix(t, "- ") ||
			strings.HasPrefix(t, "* ") || strings.HasPrefix(t, "+ ") ||
			listItemRe.MatchString(t) {
			if len(para) > 0 {
				break
			}
			continue
		}
		para = append(para, t)
	}

	desc := stripMD(strings.Join(para, " "))
	desc = mdLinkTextRe.ReplaceAllString(desc, "$1")
	desc = strings.Join(strings.Fields(desc), " ")
	if desc == "" {
		return "RobusTest documentation: " + title
	}
	// A paragraph that only introduces a following list ends in a colon, which
	// reads as truncated in a search result.
	desc = strings.TrimRight(desc, " :")
	return truncateWords(desc, maxLen)
}

// truncateWords trims s to at most max characters without splitting a word,
// appending an ellipsis when content was dropped.
func truncateWords(s string, max int) string {
	if len(s) <= max {
		return s
	}
	cut := s[:max-1]
	if i := strings.LastIndex(cut, " "); i > 0 {
		cut = cut[:i]
	}
	return strings.TrimRight(cut, " ,.;:—-") + "…"
}

var nonIDChars = regexp.MustCompile(`[^a-z0-9\- ]`)

// headingID mirrors goldmark's auto heading ID generation closely enough
// for TOC anchors (lowercase, spaces to hyphens, strip punctuation).
func headingID(text string) string {
	t := strings.ToLower(stripMD(text))
	t = nonIDChars.ReplaceAllString(t, "")
	t = strings.ReplaceAll(strings.TrimSpace(t), " ", "-")
	return t
}

var mdInline = regexp.MustCompile("[*_`]")

// listItemRe matches ordered list items ("1. ", "2) ").
var listItemRe = regexp.MustCompile(`^\d+[.)]\s`)

// mdLinkTextRe reduces "[text](target)" to just "text" for descriptions.
var mdLinkTextRe = regexp.MustCompile(`\[([^\]]*)\]\([^)]*\)`)

func stripMD(s string) string {
	return strings.TrimSpace(mdInline.ReplaceAllString(s, ""))
}
