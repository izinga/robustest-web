package handler

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/izinga/robustest-web/internal/app/docs"
	"github.com/izinga/robustest-web/internal/app/views/pages"
)

var docsStore *docs.Store

// InitDocs starts the documentation sync loop and enables /docs routes.
func InitDocs() {
	docsStore = docs.NewStore()
	docsStore.Start()
}

// DocsPage serves rendered documentation pages, synced assets, and the
// refresh endpoint.
func DocsPage(c *gin.Context) {
	path := strings.Trim(c.Param("path"), "/")

	switch {
	case path == "index.json":
		if !docsStore.Ready() {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "docs not synced"})
			return
		}
		c.Header("Cache-Control", "public, max-age=300")
		c.JSON(http.StatusOK, docsStore.Index())
		return

	case path == "refresh":
		if err := docsStore.Sync(); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"status": "error", "error": err.Error()})
			return
		}
		sha, syncedAt, _ := docsStore.Status()
		c.JSON(http.StatusOK, gin.H{"status": "ok", "sha": sha, "synced_at": syncedAt})
		return

	case strings.HasPrefix(path, "assets/"):
		dir := docsStore.Dir()
		if dir == "" {
			c.Status(http.StatusNotFound)
			return
		}
		full := filepath.Join(dir, filepath.Clean(path))
		if !strings.HasPrefix(full, filepath.Clean(dir)+string(os.PathSeparator)) {
			c.Status(http.StatusNotFound)
			return
		}
		c.File(full)
		return
	}

	if !docsStore.Ready() {
		c.Status(http.StatusServiceUnavailable)
		c.Header("Content-Type", "text/html; charset=utf-8")
		if err := pages.DocsUnavailable().Render(c.Request.Context(), c.Writer); err != nil {
			log.Printf("Error rendering docs-unavailable page: %v", err)
		}
		return
	}

	page, err := docsStore.Load(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			c.Redirect(http.StatusFound, "/docs")
			return
		}
		log.Printf("Error loading doc %q: %v", path, err)
		c.Status(http.StatusInternalServerError)
		return
	}

	prev, next := docsStore.PrevNext(path)
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := pages.DocsPage(page, docsStore.Nav(), path, prev, next).Render(c.Request.Context(), c.Writer); err != nil {
		log.Printf("Error rendering doc %q: %v", path, err)
	}
}
