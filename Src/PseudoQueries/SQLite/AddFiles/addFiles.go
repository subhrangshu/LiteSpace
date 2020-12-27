package AddFiles

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func AddFiles(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/add", func(ctx *fiber.Ctx) error {
		if ctx.Query("id") != "" {
			stmt, _ := sqlSess.Prepare("INSERT INTO hashlocation (location) values(?)")
			stmt.Exec(ctx.Query("id"))
			return ctx.SendString(ctx.Query("id"))
		}
		fmt.Print("ID Not Provided")
		stmt, _ := sqlSess.Prepare("INSERT INTO hashlocation (location) values(?)")
		stmt.Exec("Random")
		return ctx.SendString("ID Not Provided")
	})
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
