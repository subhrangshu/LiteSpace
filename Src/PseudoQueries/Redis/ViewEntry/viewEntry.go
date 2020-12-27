package ViewEntry

import (
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func ViewEntry(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/view", func(ctx *fiber.Ctx) error {
		s, _ := redSess.Do("GET", "Favorite Movie").String()
		return ctx.SendString(s)
	})
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
