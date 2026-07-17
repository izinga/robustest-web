package handler

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/izinga/robustest-web/internal/app/blog"
	"github.com/izinga/robustest-web/internal/app/views/pages"
)

var blogStore *blog.Store

// InitBlog loads posts from blog-content/ (override with BLOG_CONTENT_DIR).
// A malformed post fails loudly at startup rather than silently vanishing from
// the index — a post that doesn't render is a bug, not a draft.
func InitBlog() {
	blogStore = blog.NewStore(os.Getenv("BLOG_CONTENT_DIR"))
	if err := blogStore.Load(); err != nil {
		log.Printf("blog: load failed: %v", err)
		return
	}
	log.Printf("blog: loaded %d posts", len(blogStore.Posts()))
}

// BlogIndex renders the post ledger at /blog.
func BlogIndex(c *gin.Context) {
	if blogStore == nil {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	renderPage(c, "blog", func() error {
		return pages.BlogIndex(blogStore.ByYear()).Render(c.Request.Context(), c.Writer)
	})
}

// BlogPost renders one post at /blog/:slug.
func BlogPost(c *gin.Context) {
	if blogStore == nil {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	slug := strings.Trim(c.Param("slug"), "/")
	if slug == "" {
		BlogIndex(c)
		return
	}
	p, ok := blogStore.Get(slug)
	if !ok {
		c.Status(http.StatusNotFound)
		_ = pages.NotFoundPage().Render(c.Request.Context(), c.Writer)
		return
	}
	renderPage(c, "blog post", func() error {
		return pages.BlogPost(p, blogStore.Related(slug, 3)).Render(c.Request.Context(), c.Writer)
	})
}

// BlogPosts exposes published posts to the sitemap builder.
func BlogPosts() []*blog.Post {
	if blogStore == nil {
		return nil
	}
	return blogStore.Posts()
}
