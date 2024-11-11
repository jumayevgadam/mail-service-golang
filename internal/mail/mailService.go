package mail

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/jumayevgadam/go-mail-service/internal/model"
	"gopkg.in/gomail.v2"
)

// SendMail sends an email to the user
func SendMail(userData model.UserData) error {
	// Create a new message
	for _, mail := range userData.Email {
		m := gomail.NewMessage()
		m.SetHeader("From", "hypergadam@gmail.com")

		m.SetHeader("To", mail)

		m.SetHeader("Subject", userData.Subject)
		m.SetBody("text/html", userData.Message)

		d := gomail.NewDialer("smtp.gmail.com", 587, "hypergadam@gmail.com", "zmkzqrsnxhkvqppt")
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		// Send the email
		if err := d.DialAndSend(m); err != nil {
			return fmt.Errorf("error sending mail: %w", err)
		}
	}

	log.Println("Email sent successfully!!")
	return nil
}
