package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jumayevgadam/go-mail-service/api/handlers"
	"github.com/jumayevgadam/go-mail-service/api/middleware"
)

// SetRoutes sets up routes for the application
func SetRoutes(app *fiber.App) {
	mailRoutes := app.Group("/api/mail")
	mailRoutes.Post("/send", handlers.SendMail, middleware.ValidatorMiddleware())
}
