package KB100000

import "github.com/gofiber/fiber/v2"

func KB100000(c *fiber.Ctx) error {
	return c.SendFile("/mnt/1F8376FA6D76E846/Documents/Research Documents/More/Bigdata/Tests/100000KB.txt")
}
