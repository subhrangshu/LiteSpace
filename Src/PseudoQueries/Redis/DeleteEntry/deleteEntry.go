package DeleteEntry

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func DeleteEntry(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/delete", func(ctx *fiber.Ctx) error {
		if ctx.Query("id") != "" {
			stmt, _ := redSess.Do("DEL", ctx.Query("id")).String()
			return ctx.SendString("Deleted: " + ctx.Query("id") + " " + stmt)
		}
		fmt.Print("ID Not Provided")
		return ctx.SendString("ID Not Found")
	})
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
