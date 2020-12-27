package KB10

import "github.com/gofiber/fiber/v2"

func KB10(c *fiber.Ctx) error {
	return c.SendString("10KB")
}
