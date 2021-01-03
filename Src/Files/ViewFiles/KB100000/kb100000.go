package KB100000

import (
	"LiteSpace/Src/Files"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"os"
)

func KB100000(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/100000kb", func(ctx *fiber.Ctx) error {
		if ctx.Query("id") != "" {
			data, _ := os.Open(Files.FileTestAddress + "100000KB-" + ctx.Query("id") + ".txt")
			data.Close()
			return ctx.SendString("Viewed: " + "100000KB-" + ctx.Query("id"))
		}
		fmt.Print("ID Not Provided")
		return ctx.SendString("100000KB File: ID Not Provided")
	})
	return nil
}
