package main

import (
	"log"
	"os"

	"github.com/arizdn234/EvoPay/internal/config"
	"github.com/arizdn234/EvoPay/internal/db"
	"github.com/arizdn234/EvoPay/internal/redis"
	"github.com/arizdn234/EvoPay/internal/repositories"
	"github.com/arizdn234/EvoPay/internal/seeders"
	"github.com/arizdn234/EvoPay/internal/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize Redis
	redis.Initialize()

	// Initialize the database
	db.InitDB()

	// Migrate
	if err := repositories.Migrate(db.DB); err != nil {
		log.Fatalf("Failed to migrate the table: %v", err)
	}

	// Seeder
	if err := seeders.SeedAll(db.DB); err != nil {
		log.Fatalln("Seeding failed:", err)
	}

	// Initialize Fiber app
	app := fiber.New()

	// Enable CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
	}))

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
