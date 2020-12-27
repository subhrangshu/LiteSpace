package DeleteEntry

import (
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func DeleteEntry(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/delete", func(ctx *fiber.Ctx) error {
		stmt, _ := redSess.Do("DEL", "Favorite Movie").String()
		return ctx.SendString(stmt)
	})
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
