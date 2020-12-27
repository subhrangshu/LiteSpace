package KB1000

import "github.com/gofiber/fiber/v2"

func KB1000(c *fiber.Ctx) error {
	return c.SendString("1000KB")
}
