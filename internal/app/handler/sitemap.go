package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// docsFallbackDate is the lastmod used for the docs and blog listing pages
// before their real dates are known (docs not yet synced, no posts loaded). It
// is a fixed past date rather than "now" so an empty store can't advertise a
// fresh change that never happened.
const docsFallbackDate = "2026-07-18"

// marketingPages are the static site URLs with their sitemap priority and the
// date the page's content last actually changed.
//
// LastMod is hand-maintained on purpose: it must reflect a real content change,
// so bump it in the same commit that edits the page's template. Deriving it from
// time.Now() (as this once did) republishes every page as "modified today" on
// every request, which teaches Google the field is noise — it then ignores
// lastmod site-wide and wastes crawl budget re-fetching pages that never change.
var marketingPages = []struct {
	Path     string
	Priority string
	LastMod  string
}{
	{"/", "1.0", "2026-07-14"},
	{"/features", "0.9", "2026-07-14"},
	{"/platform/manual-testing", "0.9", "2026-07-14"},
	{"/platform/test-automation", "0.9", "2026-07-14"},
	{"/platform/performance-testing", "0.9", "2026-07-14"},
	{"/platform/tv-testing", "0.9", "2026-08-24"},
	{"/platform/network-capture", "0.9", "2026-07-14"},
	{"/platform/device-lab", "0.9", "2026-07-18"},
	{"/platform/integrations", "0.9", "2026-07-17"},
	{"/compare/robustest-vs-browserstack", "0.8", "2026-07-17"},
	{"/enterprise", "0.8", "2026-07-14"},
	{"/partners", "0.8", "2026-07-14"},
	{"/pricing", "0.8", "2026-07-17"},
	{"/security", "0.7", "2026-07-14"},
	{"/about", "0.6", "2026-07-14"},
	{"/contact", "0.5", "2026-07-14"},
	{"/legal", "0.3", "2026-07-14"},
}

// SitemapXML emits the sitemap dynamically: static marketing pages, every page
// in the currently published docs tree, and every published blog post.
func SitemapXML(c *gin.Context) {
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	b.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n")

	posts := BlogPosts()

	// The docs tree and the blog index are the two listing pages whose content
	// really does change on their own schedule, so they get a derived lastmod
	// rather than a hand-maintained one.
	//
	// Deliberately not the sync timestamp: we poll the docs repo on a timer, so
	// syncedAt advances even when nothing changed, and every docs URL would
	// claim a fresh edit on every poll.
	docsPageDate := func(path string) string {
		if docsStore == nil {
			return docsFallbackDate
		}
		if t, ok := docsStore.ModTime(path); ok {
			return t.UTC().Format("2006-01-02")
		}
		return docsFallbackDate
	}
	docsDate := docsPageDate("")
	blogDate := docsFallbackDate
	for _, p := range posts {
		d := p.Date
		if !p.Updated.IsZero() {
			d = p.Updated
		}
		if s := d.UTC().Format("2006-01-02"); s > blogDate {
			blogDate = s
		}
	}

	write := func(path, lastmod, priority string) {
		fmt.Fprintf(&b, "  <url>\n    <loc>https://robustest.com%s</loc>\n    <lastmod>%s</lastmod>\n    <priority>%s</priority>\n  </url>\n",
			path, lastmod, priority)
	}

	for _, p := range marketingPages {
		write(p.Path, p.LastMod, p.Priority)
	}
	write("/docs", docsDate, "0.7")
	write("/blog", blogDate, "0.7")
	if docsStore != nil && docsStore.Ready() {
		for _, entry := range docsStore.Index() {
			if entry.Path == "" {
				continue // /docs home already listed
			}
			write("/docs/"+entry.Path, docsPageDate(entry.Path), "0.6")
		}
	}
	// Posts carry their own lastmod: a revised post (the regulation pages get
	// refreshed at each compliance milestone) should say so.
	for _, p := range posts {
		lastmod := p.Date
		if !p.Updated.IsZero() {
			lastmod = p.Updated
		}
		write("/blog/"+p.Slug, lastmod.UTC().Format("2006-01-02"), "0.7")
	}

	b.WriteString("</urlset>\n")
	c.Header("Content-Type", "application/xml; charset=utf-8")
	c.Header("Cache-Control", "public, max-age=3600")
	c.String(http.StatusOK, b.String())
}
