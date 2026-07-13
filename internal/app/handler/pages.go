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

// ManualTestingPage renders the manual testing capability page
func ManualTestingPage(c *gin.Context) {
	renderPage(c, "manual-testing", func() error {
		return pages.ManualTestingPage().Render(c.Request.Context(), c.Writer)
	})
}

// TestAutomationPage renders the test automation capability page
func TestAutomationPage(c *gin.Context) {
	renderPage(c, "test-automation", func() error {
		return pages.TestAutomationPage().Render(c.Request.Context(), c.Writer)
	})
}

// PerformanceTestingPage renders the performance testing capability page
func PerformanceTestingPage(c *gin.Context) {
	renderPage(c, "performance-testing", func() error {
		return pages.PerformanceTestingPage().Render(c.Request.Context(), c.Writer)
	})
}

// TVTestingPage renders the Smart TV & OTT testing capability page
func TVTestingPage(c *gin.Context) {
	renderPage(c, "tv-testing", func() error {
		return pages.TVTestingPage().Render(c.Request.Context(), c.Writer)
	})
}

// NetworkCapturePage renders the network capture capability page
func NetworkCapturePage(c *gin.Context) {
	renderPage(c, "network-capture", func() error {
		return pages.NetworkCapturePage().Render(c.Request.Context(), c.Writer)
	})
}

// DeviceLabPage renders the device lab operations capability page
func DeviceLabPage(c *gin.Context) {
	renderPage(c, "device-lab", func() error {
		return pages.DeviceLabPage().Render(c.Request.Context(), c.Writer)
	})
}

// IntegrationsPage renders the integrations & enterprise capability page
func IntegrationsPage(c *gin.Context) {
	renderPage(c, "integrations", func() error {
		return pages.IntegrationsPage().Render(c.Request.Context(), c.Writer)
	})
}

// EnterprisePage renders the enterprise industry-scenarios page
func EnterprisePage(c *gin.Context) {
	renderPage(c, "enterprise", func() error {
		return pages.EnterprisePage().Render(c.Request.Context(), c.Writer)
	})
}

// PartnersPage renders the testing-services partners page
func PartnersPage(c *gin.Context) {
	renderPage(c, "partners", func() error {
		return pages.PartnersPage().Render(c.Request.Context(), c.Writer)
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
	leadType := ""
	if c.Query("type") == "partner" {
		leadType = "partner"
	}
	renderPage(c, "contact", func() error {
		return pages.ContactPage(leadType).Render(c.Request.Context(), c.Writer)
	})
}

// LegalPage renders the combined privacy and terms page
func LegalPage(c *gin.Context) {
	renderPage(c, "legal", func() error {
		return pages.LegalPage().Render(c.Request.Context(), c.Writer)
	})
}
