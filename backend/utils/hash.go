package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	fmt.Printf("hashedPassword-hash: %v\n", string(hashedPassword))
	return string(hashedPassword), nil
}

// CheckPassword checks the given password against the hashed password.
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Printf("Password check failed: %v\n", err)
		return false
	}
	fmt.Println("Password check succeeded")
	return true
}
