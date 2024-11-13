package handlers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/jumayevgadam/go-mail-service/internal/mail"
	"github.com/jumayevgadam/go-mail-service/internal/model"
)

// Handler function for configuring informations detailed about gmail
func ConfigureMailOps(c fiber.Ctx) error {
	var request model.MailSenderOps
	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"HasError": true,
			"Error":    fmt.Errorf("error parsing mail sender ops: %w", err),
		})
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"HasError": true,
			"Error":    fmt.Errorf("error validating mail sender request body: %w", err),
		})
	}

	dirPath := "./mailDetails"
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"HasError": true,
			"Error":    fmt.Errorf("error creating directory named mailDetails :%w", err),
		})
	}

	config := model.MailConfig{
		SMTPOps: model.SMTPOps{
			SMTPServer: "smtp.gmail.com",
			SMTPPort:   587,
		},
		MailSenderOps: model.MailSenderOps{
			MailSender:  request.MailSender,
			AppPassword: request.AppPassword,
		},
	}

	configPath := "./mailDetails/config.json"
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// file does not exist, create it with the configuration
		if err := os.MkdirAll("./mailDetails", 0755); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"HasError": true,
				"Error":    fmt.Errorf("error creating directory: %w", err),
			})
		}
	} else {
		// File Exists, load existing configuration to avoid overwriting settings
		configData, err := os.ReadFile(configPath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"HasError": true,
				"Error":    fmt.Errorf("error reading config file: %w", err),
			})
		}
		json.Unmarshal(configData, &config) // load existing config without overwriting SMTP settings
	}

	// update sender email and app passwords only
	config.MailSenderOps.MailSender = request.MailSender
	config.MailSenderOps.AppPassword = request.AppPassword

	// Write the updated configuration to the file
	updatedConfigData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"HasError": true,
			"Error":    fmt.Sprintf("error marshaling updated config data: %v", err),
		})
	}
	if err := os.WriteFile(configPath, updatedConfigData, 0644); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"HasError": true,
			"Error":    fmt.Sprintf("error saving config file: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"HasError": false,
		"Message":  "Sender email configuration updated successfully",
	})
}

// Handler function for sending mail
func SendMail(c fiber.Ctx) error {
	var request model.UserData
	var data string
	var statusCode int

	// Parse the request body from POST request
	c.Bind().JSON(&request)

	// Send the mail
	sendMailError := mail.SendMail(request)
	if sendMailError != nil {
		data = sendMailError.Error()
		statusCode = fiber.StatusInternalServerError
	} else {
		data = "mail sent successfully"
		statusCode = fiber.StatusOK
	}

	// Return the response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"HasError":   sendMailError != nil,
		"Error":      sendMailError,
		"Data":       data,
		"StatusCode": statusCode,
	})
}

// Handler to get the mail configuration (config.json)
func GetMailConfig(c fiber.Ctx) error {
	configPath := "./mailDetails/config.json"

	// Check if the config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"HasError": true,
			"Error":    "Config file not found",
		})
	}

	// Read the config file
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"HasError": true,
			"Error":    fmt.Sprintf("Error reading config file: %v", err),
		})
	}

	// Parse the JSON data
	var config model.MailConfig
	if err := json.Unmarshal(configData, &config); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"HasError": true,
			"Error":    fmt.Sprintf("Error parsing config file: %v", err),
		})
	}

	// Return the configuration as a JSON response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"HasError": false,
		"Config":   config,
	})
}
