package utils

import (
	"errors"

	"github.com/arizdn234/EvoPay/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// ValidateStruct validates the given struct based on tags.
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func GetUserIDFromContext(c *fiber.Ctx) (uint64, error) {
	user, exists := c.Locals("user").(*models.User) // Adjust based on how you store the user info in context
	if !exists {
		return 0, errors.New("user not found in context")
	}
	return uint64(user.ID), nil // Assuming `ID` is of type uint64
}
