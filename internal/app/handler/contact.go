package handler

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/izinga/robustest-web/internal/app/views/components"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// rateLimiter implements a simple in-memory rate limiter for contact form submissions
type rateLimiter struct {
	mu       sync.RWMutex
	requests map[string][]time.Time
	limit    int           // max requests
	window   time.Duration // time window
}

var contactRateLimiter = &rateLimiter{
	requests: make(map[string][]time.Time),
	limit:    5,              // 5 requests
	window:   5 * time.Minute, // per 5 minutes
}

// isAllowed checks if a request from the given IP is allowed
func (rl *rateLimiter) isAllowed(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.window)

	// Clean up old entries
	if timestamps, exists := rl.requests[ip]; exists {
		var valid []time.Time
		for _, t := range timestamps {
			if t.After(windowStart) {
				valid = append(valid, t)
			}
		}
		rl.requests[ip] = valid
	}

	// Check if under limit
	if len(rl.requests[ip]) >= rl.limit {
		return false
	}

	// Add new request
	rl.requests[ip] = append(rl.requests[ip], now)
	return true
}

// emailRegex provides stricter email validation
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// phoneRegex validates phone numbers (allows +, digits, spaces, dashes, parentheses)
var phoneRegex = regexp.MustCompile(`^[\d\s\-\+\(\)]{0,20}$`)

// ContactFormRequest represents the contact form submission
type ContactFormRequest struct {
	Name    string `form:"name" binding:"required,max=100"`
	Email   string `form:"email" binding:"required,email,max=254"`
	Company string `form:"company" binding:"max=200"`
	Phone   string `form:"phone" binding:"max=20"`
	Message string `form:"message" binding:"max=2000"`
}

// sanitize cleans and validates the contact form request
func (req *ContactFormRequest) sanitize() error {
	// Trim whitespace from all fields
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Company = strings.TrimSpace(req.Company)
	req.Phone = strings.TrimSpace(req.Phone)
	req.Message = strings.TrimSpace(req.Message)

	// Validate email format with stricter regex
	if !emailRegex.MatchString(req.Email) {
		return fmt.Errorf("invalid email format")
	}

	// Validate phone format if provided
	if req.Phone != "" && !phoneRegex.MatchString(req.Phone) {
		return fmt.Errorf("invalid phone number format")
	}

	return nil
}

// SubmitContactForm handles the contact form submission
func SubmitContactForm(c *gin.Context) {
	// Rate limiting check
	clientIP := c.ClientIP()
	if !contactRateLimiter.isAllowed(clientIP) {
		log.Printf("Rate limit exceeded for IP: %s", clientIP)
		c.Status(http.StatusTooManyRequests)
		if err := components.ContactFormError("Too many requests. Please wait a few minutes before trying again.").Render(c.Request.Context(), c.Writer); err != nil {
			log.Printf("Error rendering rate limit response: %v", err)
		}
		return
	}

	var req ContactFormRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Printf("Contact form validation error: %v", err)
		c.Status(http.StatusBadRequest)
		if err := components.ContactFormError("Please fill in all required fields correctly.").Render(c.Request.Context(), c.Writer); err != nil {
			log.Printf("Error rendering validation error response: %v", err)
		}
		return
	}

	// Sanitize and validate input
	if err := req.sanitize(); err != nil {
		log.Printf("Contact form sanitization error: %v", err)
		c.Status(http.StatusBadRequest)
		if err := components.ContactFormError("Please check your input and try again.").Render(c.Request.Context(), c.Writer); err != nil {
			log.Printf("Error rendering sanitization error response: %v", err)
		}
		return
	}

	// Build email content with HTML-escaped values
	subject := fmt.Sprintf("New Contact Form Submission from %s",
		html.EscapeString(req.Name))

	htmlContent := buildEmailHTML(req)
	textContent := buildEmailText(req)

	// Send email via SendGrid
	if err := sendEmail(subject, htmlContent, textContent); err != nil {
		log.Printf("Failed to send contact email: %v", err)
		c.Status(http.StatusInternalServerError)
		if err := components.ContactFormError("Failed to send your request. Please try again or email us directly at hello@robustest.com").Render(c.Request.Context(), c.Writer); err != nil {
			log.Printf("Error rendering email error response: %v", err)
		}
		return
	}

	log.Printf("Contact form submitted successfully: %s <%s>",
		req.Name, req.Email)

	// Send confirmation email to the sender
	if err := sendConfirmationEmail(req); err != nil {
		log.Printf("Failed to send confirmation email to %s: %v", req.Email, err)
	}

	// Return success response
	c.Status(http.StatusOK)
	if err := components.ContactFormSuccess().Render(c.Request.Context(), c.Writer); err != nil {
		log.Printf("Error rendering success response: %v", err)
	}
}

func buildEmailHTML(req ContactFormRequest) string {
	var sb strings.Builder

	// Escape all user input to prevent XSS in email clients
	name := html.EscapeString(req.Name)
	email := html.EscapeString(req.Email)
	company := html.EscapeString(req.Company)
	phone := html.EscapeString(req.Phone)
	message := html.EscapeString(req.Message)

	sb.WriteString(`<!DOCTYPE html>
<html>
<head>
    <style>
        body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .header { background: #2563eb; color: white; padding: 20px; border-radius: 8px 8px 0 0; }
        .content { background: #f9fafb; padding: 20px; border: 1px solid #e5e7eb; }
        .field { margin-bottom: 15px; }
        .label { font-weight: bold; color: #374151; }
        .value { margin-top: 5px; }
        .footer { background: #1f2937; color: #9ca3af; padding: 15px; border-radius: 0 0 8px 8px; font-size: 12px; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1 style="margin:0;">New Contact Form Submission</h1>
        </div>
        <div class="content">
            <div class="field">
                <div class="label">Name</div>
                <div class="value">` + name + `</div>
            </div>
            <div class="field">
                <div class="label">Email</div>
                <div class="value"><a href="mailto:` + email + `">` + email + `</a></div>
            </div>`)

	if req.Company != "" {
		sb.WriteString(`
            <div class="field">
                <div class="label">Company</div>
                <div class="value">` + company + `</div>
            </div>`)
	}

	if req.Phone != "" {
		sb.WriteString(`
            <div class="field">
                <div class="label">Phone</div>
                <div class="value">` + phone + `</div>
            </div>`)
	}

	if req.Message != "" {
		sb.WriteString(`
            <div class="field">
                <div class="label">Message</div>
                <div class="value">` + strings.ReplaceAll(message, "\n", "<br>") + `</div>
            </div>`)
	}

	sb.WriteString(`
        </div>
        <div class="footer">
            This message was sent from the contact form on robustest.com
        </div>
    </div>
</body>
</html>`)

	return sb.String()
}

func buildEmailText(req ContactFormRequest) string {
	var sb strings.Builder

	sb.WriteString("NEW CONTACT FORM SUBMISSION\n")
	sb.WriteString("===========================\n\n")
	sb.WriteString(fmt.Sprintf("Name: %s\n", req.Name))
	sb.WriteString(fmt.Sprintf("Email: %s\n", req.Email))

	if req.Company != "" {
		sb.WriteString(fmt.Sprintf("Company: %s\n", req.Company))
	}

	if req.Phone != "" {
		sb.WriteString(fmt.Sprintf("Phone: %s\n", req.Phone))
	}

	if req.Message != "" {
		sb.WriteString(fmt.Sprintf("\nMessage:\n%s\n", req.Message))
	}

	sb.WriteString("\n---\nSubmitted via RobusTest website contact form")

	return sb.String()
}

func sendEmail(subject, htmlContent, textContent string) error {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("SENDGRID_API_KEY environment variable not set")
	}

	fromEmail := os.Getenv("CONTACT_FROM_EMAIL")
	if fromEmail == "" {
		fromEmail = "noreply@robustest.com"
	}

	toEmail := os.Getenv("CONTACT_TO_EMAIL")
	if toEmail == "" {
		toEmail = "hello@robustest.com"
	}

	from := mail.NewEmail("RobusTest Website", fromEmail)
	to := mail.NewEmail("RobusTest Team", toEmail)

	message := mail.NewSingleEmail(from, subject, to, textContent, htmlContent)

	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)

	if err != nil {
		return fmt.Errorf("sendgrid error: %w", err)
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("sendgrid returned status %d: %s", response.StatusCode, response.Body)
	}

	log.Printf("Email sent successfully, status: %d", response.StatusCode)
	return nil
}

// sendConfirmationEmail sends a thank you email to the sender
func sendConfirmationEmail(req ContactFormRequest) error {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("SENDGRID_API_KEY environment variable not set")
	}

	fromEmail := os.Getenv("CONTACT_FROM_EMAIL")
	if fromEmail == "" {
		fromEmail = "noreply@robustest.com"
	}

	from := mail.NewEmail("RobusTest", fromEmail)
	to := mail.NewEmail(req.Name, req.Email)
	subject := "Thank you for contacting RobusTest"

	htmlContent := buildConfirmationHTML(req)
	textContent := buildConfirmationText(req)

	message := mail.NewSingleEmail(from, subject, to, textContent, htmlContent)

	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)

	if err != nil {
		return fmt.Errorf("sendgrid error: %w", err)
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("sendgrid returned status %d: %s", response.StatusCode, response.Body)
	}

	log.Printf("Confirmation email sent to %s, status: %d", req.Email, response.StatusCode)
	return nil
}

func buildConfirmationHTML(req ContactFormRequest) string {
	name := html.EscapeString(req.Name)
	company := html.EscapeString(req.Company)
	message := html.EscapeString(req.Message)

	var sb strings.Builder

	sb.WriteString(`<!DOCTYPE html>
<html>
<head>
    <style>
        body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .header { background: #2563eb; color: white; padding: 20px; border-radius: 8px 8px 0 0; }
        .content { background: #f9fafb; padding: 20px; border: 1px solid #e5e7eb; }
        .message-box { background: white; padding: 15px; border-radius: 4px; border: 1px solid #e5e7eb; margin-top: 15px; }
        .field { margin-bottom: 10px; }
        .label { font-weight: bold; color: #6b7280; font-size: 12px; text-transform: uppercase; }
        .value { margin-top: 4px; }
        .footer { padding: 15px; font-size: 12px; color: #6b7280; text-align: center; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h2 style="margin: 0;">Thank You for Contacting Us</h2>
        </div>
        <div class="content">
            <p>Hi ` + name + `,</p>
            <p>Thank you for reaching out to RobusTest. We have received your message and will get back to you soon.</p>
            <div class="message-box">
                <p style="margin-top: 0; font-weight: bold; color: #374151;">Your message:</p>`)

	if req.Company != "" {
		sb.WriteString(`
                <div class="field">
                    <div class="label">Company</div>
                    <div class="value">` + company + `</div>
                </div>`)
	}

	if req.Message != "" {
		sb.WriteString(`
                <div class="field">
                    <div class="label">Message</div>
                    <div class="value">` + strings.ReplaceAll(message, "\n", "<br>") + `</div>
                </div>`)
	}

	sb.WriteString(`
            </div>
        </div>
        <div class="footer">
            RobusTest<br>
            <a href="https://robustest.com">robustest.com</a>
        </div>
    </div>
</body>
</html>`)

	return sb.String()
}

func buildConfirmationText(req ContactFormRequest) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Hi %s,\n\n", req.Name))
	sb.WriteString("Thank you for reaching out to RobusTest. We have received your message and will get back to you soon.\n\n")
	sb.WriteString("--- Your message ---\n\n")

	if req.Company != "" {
		sb.WriteString(fmt.Sprintf("Company: %s\n", req.Company))
	}

	if req.Message != "" {
		sb.WriteString(fmt.Sprintf("Message: %s\n", req.Message))
	}

	sb.WriteString("\n---\nRobusTest\nhttps://robustest.com")

	return sb.String()
}
