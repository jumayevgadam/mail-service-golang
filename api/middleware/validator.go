package middleware

import (
	"github.com/gofiber/fiber/v3"
	log "github.com/gofiber/fiber/v3/log"
	"github.com/jumayevgadam/go-mail-service/internal/model"
	"github.com/jumayevgadam/go-mail-service/internal/validator"
)

// ValidatorMiddleware is a middleware to validate the request body
func ValidatorMiddleware() fiber.Handler {
	log.Info("Validation middleware started")
	return func(c fiber.Ctx) error {
		var request model.UserData
		if err := c.Bind().JSON(&request); err != nil {
			log.Error("error parsing request body")
			return badRequestError(c, "error parsing request body")
		}

		// Validate the request body
		if err := validator.Validate.Struct(request); err != nil {
			log.Error("error validating request body")
			return badRequestError(c, "error validating request body")
		}

		return c.Next()
	}
}

// badRequestError is a helper function to return a bad request error
func badRequestError(c fiber.Ctx, errorMsg string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"HasError":   true,
		"Error":      errorMsg,
		"Data":       nil,
		"StatusCode": fiber.StatusBadRequest,
	})
}
