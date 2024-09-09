package utils

import (
	"log"
	"os"
)

// InitializeLogger initializes a logger with a specific format.
func InitializeLogger() *log.Logger {
	return log.New(os.Stdout, "EvoPay: ", log.LstdFlags|log.Lshortfile)
}
