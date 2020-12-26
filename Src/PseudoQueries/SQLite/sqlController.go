package SQLite

import (
	"LiteSpace/Src/PseudoQueries/SQLite/AddFiles"
	"LiteSpace/Src/PseudoQueries/SQLite/DeleteFiles"
	"LiteSpace/Src/PseudoQueries/SQLite/ViewFiles"
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func SqlPseudoController(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) {
	router.Get("/", FakeSqlite)
	AddFiles.AddFiles(fibSess, sqlSess, redSess, router)
	DeleteFiles.DeleteFiles(fibSess, sqlSess, redSess, router)
	ViewFiles.ViewFiles(fibSess, sqlSess, redSess, router)
}

func FakeSqlite(ctx *fiber.Ctx) error {
	return ctx.SendString("Fake SQL Page")
}
