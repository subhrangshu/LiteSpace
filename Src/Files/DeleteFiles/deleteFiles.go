package DeleteFiles

import "github.com/gofiber/fiber/v2"

func DeleteFiles(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete Files")
}
