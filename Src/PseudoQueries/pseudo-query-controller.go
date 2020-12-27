package PseudoQueries

import (
	"LiteSpace/Src/PseudoQueries/Redis"
	"LiteSpace/Src/PseudoQueries/SQLite"
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func PseudoQueryController(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client) {
	SQLite.SqlPseudoController(fibSess, sqlSess, redSess, fibSess.Group("fake-sqlite"))
	Redis.RedisPseudoController(fibSess, sqlSess, redSess, fibSess.Group("fake-redis"))
}
