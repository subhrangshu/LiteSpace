package Redis

import (
	"LiteSpace/Src/PseudoQueries/Redis/AddEntry"
	"LiteSpace/Src/PseudoQueries/Redis/DeleteEntry"
	"LiteSpace/Src/PseudoQueries/Redis/ViewEntry"
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func RedisPseudoController(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) {
	router.Get("/", FakeRedis)
	AddEntry.AddEntry(fibSess, sqlSess, redSess, router)
	DeleteEntry.DeleteEntry(fibSess, sqlSess, redSess, router)
	ViewEntry.ViewEntry(fibSess, sqlSess, redSess, router)
}

func FakeRedis(ctx *fiber.Ctx) error {
	return ctx.SendString("Fake Redis Page")
}
