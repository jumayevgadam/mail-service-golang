package middleware

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

// OriginMiddleware is a middleware to check the origin of the request
func OriginMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		log.Info("Origin middleware started")
		// Get the origin header
		requestOrigin := c.Get("Origin")
		// Check if the origin is allowed
		if requestOrigin == os.Getenv("ALLOWED_ORIGIN") {
			// Allow the request to continue
			return c.Next()
		}

		// Return an error if the origin is not allowed
		log.Warnf("Origin not allowed: %s", requestOrigin)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"HasError":   true,
			"Error":      "Origin not allowed",
			"Data":       requestOrigin,
			"StatusCode": fiber.StatusForbidden,
		})
	}
}
