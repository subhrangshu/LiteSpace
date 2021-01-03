package KB10000

import (
	"LiteSpace/Src/Files"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"os"
)

func KB10000(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/10000kb", func(ctx *fiber.Ctx) error {
		if ctx.Query("id") != "" {
			data, _ := os.Open(Files.FileTestAddress + "10000KB-" + ctx.Query("id") + ".txt")
			data.Close()
			return ctx.SendString("Viewed: " + "10000KB-" + ctx.Query("id"))
		}
		fmt.Print("ID Not Provided")
		return ctx.SendString("10000KB File: ID Not Provided")
	})
	return nil
}
