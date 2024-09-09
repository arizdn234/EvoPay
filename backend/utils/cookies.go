package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// SetTokenCookie sets the JWT token in a cookie.
func SetTokenCookie(c *fiber.Ctx, token string) {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
}
