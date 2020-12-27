package ViewEntry

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func ViewEntry(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/view", func(ctx *fiber.Ctx) error {
		if ctx.Query("id") != "" {
			s, _ := redSess.Do("GET", ctx.Query("id")).String()
			return ctx.SendString(s)
		}
		fmt.Print("ID Not Provided")
		return ctx.SendString("ID not provided")
	})
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
