package Files

import (
	"LiteSpace/Src/Files/AddFiles"
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func FilesController(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) {
	router.Get("/", FilesPage)
	router.Get("/add", AddFiles.AddFiles)
}

func FilesPage(ctx *fiber.Ctx) error {
	return ctx.SendString("Files Page")
}
