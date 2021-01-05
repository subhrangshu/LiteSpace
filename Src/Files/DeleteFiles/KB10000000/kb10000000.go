package KB10000000

import "github.com/gofiber/fiber/v2"

func KB10000000(c *fiber.Ctx) error {
	return c.SendString("10000000KB")
}
