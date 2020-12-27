package ViewFiles

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func ViewFiles(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/view", func(ctx *fiber.Ctx) error {
		if ctx.Query("id") != "" {
			stmt, _ := sqlSess.Query("select location from hashlocation where id like " + ctx.Query("id"))
			var location string
			for stmt.Next() {
				stmt.Scan(&location)
			}
			return ctx.SendString(location)
		}
		fmt.Print("ID Not Provided")
		return ctx.SendString("ID Not Provided")
	})
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
