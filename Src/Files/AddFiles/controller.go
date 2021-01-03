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
	KB1.KB1(fibSess, sqlSess, redSess, router)
	KB10.KB10(fibSess, sqlSess, redSess, router)
	KB100.KB100(fibSess, sqlSess, redSess, router)
	KB1000.KB1000(fibSess, sqlSess, redSess, router)
	KB10000.KB10000(fibSess, sqlSess, redSess, router)
	KB100000.KB100000(fibSess, sqlSess, redSess, router)
}
