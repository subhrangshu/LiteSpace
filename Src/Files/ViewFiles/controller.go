package ViewFiles

import (
	"LiteSpace/Src/Files/ViewFiles/KB1"
	"LiteSpace/Src/Files/ViewFiles/KB10"
	"LiteSpace/Src/Files/ViewFiles/KB100"
	"LiteSpace/Src/Files/ViewFiles/KB1000"
	"LiteSpace/Src/Files/ViewFiles/KB10000"
	"LiteSpace/Src/Files/ViewFiles/KB100000"
	"LiteSpace/Src/Files/ViewFiles/KB1000000"
	"LiteSpace/Src/Files/ViewFiles/KB10000000"
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func ViewFilesController(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) {
	router.Get("/", ViewFiles)

	KB1.KB1(fibSess, sqlSess, redSess, router)
	KB10.KB10(fibSess, sqlSess, redSess, router)
	KB100.KB100(fibSess, sqlSess, redSess, router)
	KB1000.KB1000(fibSess, sqlSess, redSess, router)
	KB10000.KB10000(fibSess, sqlSess, redSess, router)
	KB100000.KB100000(fibSess, sqlSess, redSess, router)
	KB1000000.KB1000000(fibSess, sqlSess, redSess, router)
	KB10000000.KB10000000(fibSess, sqlSess, redSess, router)
}
