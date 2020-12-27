package KB100000

import "github.com/gofiber/fiber/v2"

func KB100000(c *fiber.Ctx) error {
	return c.SendString("100000KB")
}
