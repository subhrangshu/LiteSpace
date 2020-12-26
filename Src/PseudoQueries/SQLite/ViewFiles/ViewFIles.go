package ViewFiles

import (
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func ViewFiles(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/view", func(ctx *fiber.Ctx) error {
		stmt, err := sqlSess.Query("select location from hashlocation where id like 5")
		checkErr(err)
		var location string
		for stmt.Next() {
			stmt.Scan(&location)
		}
		return ctx.SendString(location)
	})
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
