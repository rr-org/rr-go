package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/naflafadia/go-trivia-app/db"
	"github.com/naflafadia/go-trivia-app/routers"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Setup routes
	routers.SetupRoutes(app)

	// Start Server
	err := app.Listen(":8000")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	// Handler
	client := db.MgoConn()
	defer client.Disconnect(context.TODO())
}
