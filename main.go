package main

import (
	"log"

	"github.com/fehirde/backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Register the routes
	app.Get("/healthz", routes.Healthz)

	// Start the server on port 9020
	log.Fatal(app.Listen(":9020"))
}
