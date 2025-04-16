package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// AuthorizeIP is a middleware that checks if the request comes from an authorized IP address.
func AuthorizeIP() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.IP() != "127.0.0.1" {
			return c.Status(fiber.StatusForbidden).SendString(fmt.Sprintf("IP %s is not authorized", c.IP()))
		}
		return c.Next()
	}
}
