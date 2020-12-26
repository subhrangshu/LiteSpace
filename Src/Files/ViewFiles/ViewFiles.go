package ViewFiles

import "github.com/gofiber/fiber/v2"

func ViewFiles(ctx *fiber.Ctx) error {
	return ctx.SendString("View Files")
}
