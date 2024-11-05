package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/jumayevgadam/go-mail-service/api/routes"
	"github.com/jumayevgadam/go-mail-service/internal/validator"
)

func main() {
	log.Info("Starting mail service")
	// Load Environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Errorf("error loading .env file")
	}

	// Creating a new fiber app
	log.Info("Setting up Fiber instance with CORS middleware and routes")
	app := fiber.New()
	// Setting up CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))
	log.Info("Setting up routes")

	// Setting up routes
	routes.SetRoutes(app)
	log.Info("Setting up validator")
	// Setting up validator
	validator.SetUpValidation()

	// Starting the server
	log.Infof("Server is running on http http://localhost:4000")
	app.Listen(":4000")
}
