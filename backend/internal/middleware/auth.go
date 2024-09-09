package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authToken := c.Cookies("auth_token")
		if authToken == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		tokenClaims, err := VerifyToken(authToken)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		c.Locals("userID", tokenClaims.ID)
		c.Locals("userRole", tokenClaims.RoleID)

		return c.Next()
	}
}
