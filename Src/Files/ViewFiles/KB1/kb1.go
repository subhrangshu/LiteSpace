package KB1

import (
	"LiteSpace/Src/Shared"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
)

func KB1(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/1kb", func(ctx *fiber.Ctx) error {
		if ctx.Query("id") != "" {
			_, _ = ioutil.ReadFile(Shared.FileTestAddress + "1KB-" + ctx.Query("id") + ".txt")
			return ctx.SendString("Viewed: " + "1KB-" + ctx.Query("id"))
		}
		fmt.Print("ID Not Provided")
		return ctx.SendString("1KB File: ID Not Provided")
	})
	return nil
}
