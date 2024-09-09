package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// RequireAdminRole is a middleware to check if the user has an Admin role
func RequireAdminRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authToken := c.Cookies("auth_token")
		if authToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized access, token missing"})
		}

		tokenClaims, err := VerifyToken(authToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized access, invalid token"})
		}

		// Check if user role is "admin"
		if tokenClaims.RoleID != 1 {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden access, Admin role required"})
		}

		c.Locals("userID", tokenClaims.ID)

		return c.Next()
	}
}
