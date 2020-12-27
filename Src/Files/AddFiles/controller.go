package AddFiles

import (
	"LiteSpace/Src/Files/AddFiles/KB1"
	"LiteSpace/Src/Files/AddFiles/KB10"
	"LiteSpace/Src/Files/AddFiles/KB100"
	"LiteSpace/Src/Files/AddFiles/KB1000"
	"LiteSpace/Src/Files/AddFiles/KB10000"
	"LiteSpace/Src/Files/AddFiles/KB100000"
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func AddFilesController(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) {
	router.Get("/", AddFiles)
	router.Get("/1kb", KB1.KB1)
	router.Get("/10kb", KB10.KB10)
	router.Get("/100kb", KB100.KB100)
	router.Get("/1000kb", KB1000.KB1000)
	router.Get("/10000kb", KB10000.KB10000)
	router.Get("/100000kb", KB100000.KB100000)
}
