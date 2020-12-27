package Files

import (
	"LiteSpace/Src/Files/AddFiles"
	"LiteSpace/Src/Files/DeleteFiles"
	"LiteSpace/Src/Files/ViewFiles"
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func FilesController(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) {
	router.Get("/", FilesPage)
	ViewFiles.ViewFilesController(fibSess, sqlSess, redSess, router.Group("/view"))
	DeleteFiles.DeleteFilesController(fibSess, sqlSess, redSess, router.Group("/delete"))
	AddFiles.AddFilesController(fibSess, sqlSess, redSess, router.Group("/add"))
}

func FilesPage(ctx *fiber.Ctx) error {
	return ctx.SendString("Files Page")
}
