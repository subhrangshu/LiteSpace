package KB100

import "github.com/gofiber/fiber/v2"

func KB100(c *fiber.Ctx) error {
	return c.SendString("100KB")
}
