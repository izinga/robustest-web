package handler

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/izinga/robustest-web/internal/app/views/components"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// ContactFormRequest represents the contact form submission
type ContactFormRequest struct {
	FirstName string `form:"firstName" binding:"required"`
	LastName  string `form:"lastName" binding:"required"`
	Email     string `form:"email" binding:"required,email"`
	Company   string `form:"company" binding:"required"`
	Phone     string `form:"phone"`
	Devices   string `form:"devices"`
	Message   string `form:"message"`
}

// SubmitContactForm handles the contact form submission
func SubmitContactForm(c *gin.Context) {
	var req ContactFormRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Printf("Contact form validation error: %v", err)
		components.ContactFormError("Please fill in all required fields correctly.").Render(c.Request.Context(), c.Writer)
		return
	}

	// Build email content
	subject := fmt.Sprintf("New Demo Request from %s %s - %s", req.FirstName, req.LastName, req.Company)

	htmlContent := buildEmailHTML(req)
	textContent := buildEmailText(req)

	// Send email via SendGrid
	if err := sendEmail(subject, htmlContent, textContent); err != nil {
		log.Printf("Failed to send contact email: %v", err)
		components.ContactFormError("Failed to send your request. Please try again or email us directly at hello@robustest.com").Render(c.Request.Context(), c.Writer)
		return
	}

	log.Printf("Contact form submitted successfully: %s %s <%s> from %s", req.FirstName, req.LastName, req.Email, req.Company)

	// Return success response
	components.ContactFormSuccess().Render(c.Request.Context(), c.Writer)
}

func buildEmailHTML(req ContactFormRequest) string {
	var sb strings.Builder

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
            <h1 style="margin:0;">New Demo Request</h1>
        </div>
        <div class="content">
            <div class="field">
                <div class="label">Name</div>
                <div class="value">` + req.FirstName + ` ` + req.LastName + `</div>
            </div>
            <div class="field">
                <div class="label">Email</div>
                <div class="value"><a href="mailto:` + req.Email + `">` + req.Email + `</a></div>
            </div>
            <div class="field">
                <div class="label">Company</div>
                <div class="value">` + req.Company + `</div>
            </div>`)

	if req.Phone != "" {
		sb.WriteString(`
            <div class="field">
                <div class="label">Phone</div>
                <div class="value">` + req.Phone + `</div>
            </div>`)
	}

	if req.Devices != "" {
		sb.WriteString(`
            <div class="field">
                <div class="label">Number of Devices</div>
                <div class="value">` + req.Devices + `</div>
            </div>`)
	}

	if req.Message != "" {
		sb.WriteString(`
            <div class="field">
                <div class="label">Message</div>
                <div class="value">` + strings.ReplaceAll(req.Message, "\n", "<br>") + `</div>
            </div>`)
	}

	sb.WriteString(`
        </div>
        <div class="footer">
            This demo request was submitted via the RobusTest website contact form.
        </div>
    </div>
</body>
</html>`)

	return sb.String()
}

func buildEmailText(req ContactFormRequest) string {
	var sb strings.Builder

	sb.WriteString("NEW DEMO REQUEST\n")
	sb.WriteString("================\n\n")
	sb.WriteString(fmt.Sprintf("Name: %s %s\n", req.FirstName, req.LastName))
	sb.WriteString(fmt.Sprintf("Email: %s\n", req.Email))
	sb.WriteString(fmt.Sprintf("Company: %s\n", req.Company))

	if req.Phone != "" {
		sb.WriteString(fmt.Sprintf("Phone: %s\n", req.Phone))
	}

	if req.Devices != "" {
		sb.WriteString(fmt.Sprintf("Number of Devices: %s\n", req.Devices))
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
