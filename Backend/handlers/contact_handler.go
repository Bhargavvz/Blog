package handlers

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"

	"blog-backend/db"
	"blog-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitContactHandler(client *mongo.Client) {
	// Remove this since it's handled in db package
}

func HandleContact(c *gin.Context) {
	log.Printf("Received contact form submission")

	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Contact form received with email: %s", contact.Email)

	// Validate required fields
	if contact.Name == "" || contact.Email == "" || contact.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name, email and message are required"})
		return
	}

	contact.CreatedAt = time.Now()

	// Save to database
	result, err := db.ContactCollection.InsertOne(context.Background(), contact)
	if err != nil {
		log.Printf("Error saving contact: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving contact"})
		return
	}

	// Send notification to admin with user's message
	go sendNotificationEmail(contact)

	// Send thank you email to user
	go sendThankYouEmail(contact)

	c.JSON(http.StatusCreated, gin.H{
		"id":      result.InsertedID,
		"message": "Thank you for your message. We'll get back to you soon!",
	})
}

func sendNotificationEmail(contact models.Contact) {
	subject := fmt.Sprintf("New Contact Message from %s", contact.Name)
	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<style>
			body { font-family: Arial, sans-serif; line-height: 1.6; }
			.container { max-width: 600px; margin: 0 auto; padding: 20px; }
			.header { background-color: #f8f9fa; padding: 20px; text-align: center; }
			.content { padding: 20px; }
			.message { background-color: #f5f5f5; padding: 15px; border-radius: 5px; }
			.footer { text-align: center; color: #666; font-size: 14px; }
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h2>New Contact Form Message</h2>
			</div>
			<div class="content">
				<p><strong>From:</strong> %s (%s)</p>
				<p><strong>Subject:</strong> %s</p>
				<div class="message">
					<p><strong>Message:</strong></p>
					<p>%s</p>
				</div>
				<p><strong>Received at:</strong> %s</p>
			</div>
			<div class="footer">
				<p>You can reply directly to this email to respond to the sender.</p>
			</div>
		</div>
	</body>
	</html>
	`, contact.Name, contact.Email, contact.Subject, contact.Message, contact.CreatedAt.Format("Jan 02, 2006 15:04:05"))

	// Send to both admin and user emails
	toEmails := []string{os.Getenv("ADMIN_EMAIL"), contact.Email}
	sendEmail(toEmails, subject, body)
}

func sendThankYouEmail(contact models.Contact) {
	subject := "Thank you for contacting us"
	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<style>
			body { font-family: Arial, sans-serif; line-height: 1.6; }
			.container { max-width: 600px; margin: 0 auto; padding: 20px; }
			.header { 
				background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
				padding: 30px;
				text-align: center;
				color: white;
			}
			.content { padding: 20px; }
			.footer { text-align: center; color: #666; font-size: 14px; }
			.button {
				display: inline-block;
				padding: 10px 20px;
				background-color: #4f46e5;
				color: white;
				text-decoration: none;
				border-radius: 5px;
				margin-top: 15px;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h1>Thank You for Contacting Us</h1>
			</div>
			<div class="content">
				<p>Dear %s,</p>
				<p>Thank you for reaching out to us. We have received your message and will get back to you as soon as possible.</p>
				<p>We typically respond within 24-48 hours.</p>
				<center>
					<a href="http://localhost:5173" class="button" style="color: white;">Visit Our Blog</a>
				</center>
			</div>
			<div class="footer">
				<p>This is an automated response. Please don't reply to this email.</p>
			</div>
		</div>
	</body>
	</html>
	`, contact.Name)

	// Send to both admin and user emails
	toEmails := []string{os.Getenv("ADMIN_EMAIL"), contact.Email}
	sendEmail(toEmails, subject, body)
}

func sendEmail(to []string, subject, body string) {
	from := os.Getenv("SMTP_FROM")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	headers := make(map[string]string)
	headers["From"] = fmt.Sprintf("Blog Contact <%s>", from)
	headers["To"] = strings.Join(to, ", ") // Join all email addresses
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"

	var message bytes.Buffer
	for key, value := range headers {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	message.WriteString("\r\n")
	message.WriteString(body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
		to, // Pass the slice of email addresses
		message.Bytes(),
	)

	if err != nil {
		log.Printf("Error sending email to %v: %v", to, err)
	} else {
		log.Printf("Email sent successfully to %v", to)
	}
}
