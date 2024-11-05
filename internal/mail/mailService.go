package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jumayevgadam/go-mail-service/internal/model"
	"gopkg.in/gomail.v2"
)

// SendMail sends an email to the user
func SendMail(userData model.UserData) error {
	// Define SMTP configuration
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	// Create a new message
	m := gomail.NewMessage()
	m.SetHeader("From", "hypergadam@gmail.com")
	m.SetHeader("To", userData.Email)
	m.SetHeader("Subject", "Thank you for interest!")
	m.SetBody("text/html", generateHTMLMessage(userData.FullName))
	m.Attach("https://ramsesramva.com/pdf/ramsesramva_cv.pdf")

	// Set up the STMP dialer
	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Fatalf("invalid SMTP port: %v", err.Error())
		return err
	}

	d := gomail.NewDialer(smtpHost, port, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("error sending mail: %w", err)
	}

	log.Println("Email sent successfully!!")
	return nil
}

// generateHTMLMessage generates an HTML message for the email
func generateHTMLMessage(recipientName string) string {
	return fmt.Sprintf(
		`<!DOCTYPE html>
    <html>
    <head>
        <title>Email</title>
    </head>
    <body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; margin: 0; padding: 20px; background-color: #f9f9f9;">
        <div style="max-width: 600px; margin: 0 auto; background: #fff; padding: 20px; border-radius: 10px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);">
            <h1 style="font-size: 24px; font-weight: bold; color: #444;">Dear %s,</h1>
            <p style="margin: 15px 0;">
                I wanted to take a moment to express my sincere gratitude for your interest in my web portfolio. 
                I am truly excited about the possibility of collaborating with you on a great project.
            </p>
            <p style="margin: 15px 0;">
                Your message has been a wonderful encouragement, and I am eager to discuss how my skills and 
                experience could contribute to your team's success. I have attached my CV for your reference, 
                and I look forward to the opportunity to further discuss my qualifications and how I might be 
                able to assist you.
            </p>
            <p style="margin: 15px 0;">
                Thank you again for your time and consideration. I am available to chat at your convenience, 
                and I look forward to the possibility of working together.
            </p>
            <p style="margin-top: 30px; font-weight: bold;">
                Best regards,<br>Ramsés Ramírez Vallejo
            </p>
        </div>
    </body>
    </html>
    `, recipientName)
}
