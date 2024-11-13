package mail

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jumayevgadam/go-mail-service/internal/model"
	"gopkg.in/gomail.v2"
)

// Helper function to load email configuration from config.json
func loadMailConfig() (*model.MailConfig, error) {
	configPath := "./mailDetails/config.json"
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config model.MailConfig
	if err := json.Unmarshal(configData, &config); err != nil {
		return nil, fmt.Errorf("error unmarshalling config data: %w", err)
	}

	return &config, nil
}

// SendMail sends an email to the user
func SendMail(userData model.UserData) error {
	// load the email configuration
	config, err := loadMailConfig()
	if err != nil {
		return fmt.Errorf("failed to load mail configuration: %w", err)
	}

	// Create a new message
	for _, mail := range userData.Email {
		m := gomail.NewMessage()
		m.SetHeader("From", config.MailSenderOps.MailSender)

		m.SetHeader("To", mail)

		m.SetHeader("Subject", userData.Subject)
		m.SetBody("text/html", userData.Message)

		d := gomail.NewDialer(
			config.SMTPOps.SMTPServer,
			config.SMTPOps.SMTPPort,
			config.MailSenderOps.MailSender,
			config.MailSenderOps.AppPassword,
		) // jclwgcmcxriiusnk
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		// Send the email
		if err := d.DialAndSend(m); err != nil {
			return fmt.Errorf("error sending mail: %w", err)
		}
	}

	log.Println("Email sent successfully!!")
	return nil
}
