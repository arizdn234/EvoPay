package config

import (
	"log"

	"github.com/joho/godotenv"
)

// env
func LoadEnv() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
