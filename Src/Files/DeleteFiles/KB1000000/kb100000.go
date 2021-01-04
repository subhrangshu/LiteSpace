package KB1000000

import "github.com/gofiber/fiber/v2"

func KB1000000(c *fiber.Ctx) error {
	return c.SendString("1000000KB")
}
