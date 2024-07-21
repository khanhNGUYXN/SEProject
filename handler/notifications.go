// File: handler/notification.go

package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
)

// Function to send email notifications
func SendEmailNotification(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "your-email@example.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// SMTP server configuration
	d := gomail.NewDialer("smtp.example.com", 587, "your-username", "your-password")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// Echo handler to trigger sending notifications
func NotifyCustomer(c echo.Context) error {
	email := c.QueryParam("email")
	subject := "Notification Subject"
	message := "<h1>This is your Notification</h1><p>Details about the notification.</p>"

	if err := SendEmailNotification(email, subject, message); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to send email"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Email sent successfully"})
}
