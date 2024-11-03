package main

import (
	"log"
	"os"

	"github.com/fehirde/backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize a new Fiber app,
	// with the logger middleware enabled
	app := fiber.New()
	app.Use(logger.New())

	// Register the routes
	app.Get("/healthz", routes.Healthz)

	// Start the server on port set by the SERVER_PORT environment variable
	// If the SERVER_PORT environment variable is not set, default to 9020
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ":9020"
	}
	log.Fatal(app.Listen(serverPort))
}
