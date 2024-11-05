package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jumayevgadam/go-mail-service/internal/mail"
	"github.com/jumayevgadam/go-mail-service/internal/model"
)

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
