package PrimaryController

import (
	"LiteSpace/Src/Files"
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func Controller(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client) {
	Files.FilesController(fibSess, sqlSess, redSess, fibSess.Group("files"))
}
