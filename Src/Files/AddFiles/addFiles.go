package AddFiles

import (
	"github.com/gofiber/fiber/v2"
)

func AddFiles(ctx *fiber.Ctx) error {
	return ctx.SendString("Add Files")
}
