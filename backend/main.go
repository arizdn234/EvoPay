package main

import (
	"log"
	"os"

	"github.com/arizdn234/EvoPay/internal/config"
	"github.com/arizdn234/EvoPay/internal/db"
	"github.com/arizdn234/EvoPay/internal/redis"
	"github.com/arizdn234/EvoPay/internal/repository"
	"github.com/arizdn234/EvoPay/internal/server"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize Redis
	redis.Initialize()

	// Initialize the database
	db.InitDB()

	// Migrate
	if err := repository.Migrate(db.DB); err != nil {
		log.Fatalf("Failed to migrate the table: %v", err)
	}

	// Seed data
	if err := repository.NewUserRepository(db.DB).Seed(); err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}

	// Initialize Fiber app
	app := fiber.New()

	// Enable CORS middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		return c.Next()
	})

	// Get the port from the environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.RunServer(app, db.DB, port)
	log.Printf("Server is running on port %v\n\nhttp://localhost:%v", port, port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
