package KB1

import "github.com/gofiber/fiber/v2"

func KB1(c *fiber.Ctx) error {
	return c.SendString("1KB")
}
