package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// ValidateStruct validates the given struct based on tags.
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func GetUserIDFromContext(c *fiber.Ctx) (uint64, error) {
	userID := c.Locals("userID")

	// Use type assertion for different numeric types
	switch id := userID.(type) {
	case int:
		return uint64(id), nil
	case int64:
		return uint64(id), nil
	case float64:
		return uint64(id), nil
	case uint64:
		return id, nil
	case uint:
		return uint64(id), nil
	default:
		return 0, errors.New("user ID not found in context or invalid type")
	}
}
