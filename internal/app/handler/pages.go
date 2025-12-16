package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izinga/robustest-web/internal/app/views/pages"
)

// renderPage is a helper function that handles template rendering with proper error handling
func renderPage(c *gin.Context, pageName string, renderFunc func() error) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := renderFunc(); err != nil {
		log.Printf("Error rendering %s page: %v", pageName, err)
		c.Status(http.StatusInternalServerError)
		return
	}
}

// HomePage renders the home page
func HomePage(c *gin.Context) {
	renderPage(c, "home", func() error {
		return pages.HomePage().Render(c.Request.Context(), c.Writer)
	})
}

// FeaturesPage renders the features page
func FeaturesPage(c *gin.Context) {
	renderPage(c, "features", func() error {
		return pages.FeaturesPage().Render(c.Request.Context(), c.Writer)
	})
}

// PricingPage renders the pricing page
func PricingPage(c *gin.Context) {
	renderPage(c, "pricing", func() error {
		return pages.PricingPage().Render(c.Request.Context(), c.Writer)
	})
}

// SecurityPage renders the security page
func SecurityPage(c *gin.Context) {
	renderPage(c, "security", func() error {
		return pages.SecurityPage().Render(c.Request.Context(), c.Writer)
	})
}

// AboutPage renders the about page
func AboutPage(c *gin.Context) {
	renderPage(c, "about", func() error {
		return pages.AboutPage().Render(c.Request.Context(), c.Writer)
	})
}

// ContactPage renders the contact page
func ContactPage(c *gin.Context) {
	renderPage(c, "contact", func() error {
		return pages.ContactPage().Render(c.Request.Context(), c.Writer)
	})
}

// LegalPage renders the combined privacy and terms page
func LegalPage(c *gin.Context) {
	renderPage(c, "legal", func() error {
		return pages.LegalPage().Render(c.Request.Context(), c.Writer)
	})
}
