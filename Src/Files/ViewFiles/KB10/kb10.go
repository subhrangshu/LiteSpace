package KB10

import "github.com/gofiber/fiber/v2"

func KB10(c *fiber.Ctx) error {
	return c.SendFile("/mnt/1F8376FA6D76E846/Documents/Research Documents/More/Bigdata/Tests/10KB.txt")
}
