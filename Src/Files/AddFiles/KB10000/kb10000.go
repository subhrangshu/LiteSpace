package KB10000

import "github.com/gofiber/fiber/v2"

func KB10000(c *fiber.Ctx) error {
	return c.SendString("100000KB")
}
