package middlewares

import (
	"fmt"
	"net"

	"github.com/gofiber/fiber/v2"
)

var privateIPBlocks []*net.IPNet

func init() {
	for _, cidr := range []string{
		"127.0.0.0/8",    // loopback
		"10.0.0.0/8",     // class A
		"172.16.0.0/12",  // class B
		"192.168.0.0/16", // class C
	} {
		_, block, _ := net.ParseCIDR(cidr)
		privateIPBlocks = append(privateIPBlocks, block)
	}
}

// isPrivateIP checks if the given IP is in a private range
func isPrivateIP(ip net.IP) bool {
	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

// AuthorizeIP checks if the request comes from an authorized IP address.
func AuthorizeIP() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := net.ParseIP(c.IP())
		if ip == nil || !isPrivateIP(ip) {
			return c.Status(fiber.StatusForbidden).SendString(fmt.Sprintf("IP %s is not authorized", c.IP()))
		}
		return c.Next()
	}
}
