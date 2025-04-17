package middlewares

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// AuthorizeIP is a middleware that checks if the request comes from an authorized IP address.
func AuthorizeIP() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		if ip != "127.0.0.1" && !strings.HasPrefix(ip, "192.168.") && !strings.HasPrefix(ip, "172.") {
			return c.Status(fiber.StatusForbidden).SendString(fmt.Sprintf("IP %s is not authorized", ip))
		}
		return c.Next()
	}
}
