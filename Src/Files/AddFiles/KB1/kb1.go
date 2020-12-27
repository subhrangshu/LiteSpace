package KB1

import (
	"github.com/gofiber/fiber/v2"
	"io"
	"os"
)

func KB1(c *fiber.Ctx) error {
	oldFile, _ := os.Open("/mnt/1F8376FA6D76E846/Documents/Research Documents/More/Bigdata/1KB.txt")
	newFile, _ := os.Create("/mnt/1F8376FA6D76E846/Documents/Research Documents/More/Bigdata/Tests/1KB.txt")
	io.Copy(newFile, oldFile)
	oldFile.Close()
	newFile.Close()
	return c.SendString("1KB")
}
