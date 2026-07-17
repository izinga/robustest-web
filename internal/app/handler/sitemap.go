package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// marketingPages are the static site URLs with their sitemap priority.
var marketingPages = []struct {
	Path     string
	Priority string
}{
	{"/", "1.0"},
	{"/features", "0.9"},
	{"/platform/manual-testing", "0.9"},
	{"/platform/test-automation", "0.9"},
	{"/platform/performance-testing", "0.9"},
	{"/platform/tv-testing", "0.9"},
	{"/platform/network-capture", "0.9"},
	{"/platform/device-lab", "0.9"},
	{"/platform/integrations", "0.9"},
	{"/compare/robustest-vs-browserstack", "0.8"},
	{"/enterprise", "0.8"},
	{"/partners", "0.8"},
	{"/pricing", "0.8"},
	{"/security", "0.7"},
	{"/docs", "0.7"},
	{"/blog", "0.7"},
	{"/about", "0.6"},
	{"/contact", "0.5"},
	{"/legal", "0.3"},
}

// SitemapXML emits the sitemap dynamically: static marketing pages, every page
// in the currently published docs tree, and every published blog post.
func SitemapXML(c *gin.Context) {
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	b.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n")

	today := time.Now().UTC().Format("2006-01-02")
	docsDate := today
	if docsStore != nil {
		if _, syncedAt, _ := docsStore.Status(); !syncedAt.IsZero() {
			docsDate = syncedAt.UTC().Format("2006-01-02")
		}
	}

	write := func(path, lastmod, priority string) {
		fmt.Fprintf(&b, "  <url>\n    <loc>https://robustest.com%s</loc>\n    <lastmod>%s</lastmod>\n    <priority>%s</priority>\n  </url>\n",
			path, lastmod, priority)
	}

	for _, p := range marketingPages {
		write(p.Path, today, p.Priority)
	}
	if docsStore != nil && docsStore.Ready() {
		for _, entry := range docsStore.Index() {
			if entry.Path == "" {
				continue // /docs home already listed
			}
			write("/docs/"+entry.Path, docsDate, "0.6")
		}
	}
	// Posts carry their own lastmod: a revised post (the regulation pages get
	// refreshed at each compliance milestone) should say so.
	for _, p := range BlogPosts() {
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
