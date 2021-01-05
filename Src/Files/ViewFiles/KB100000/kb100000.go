package KB100000

import (
	"LiteSpace/Src/Shared"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
)

func KB100000(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/100000kb", func(ctx *fiber.Ctx) error {
		if ctx.Query("id") != "" {
			// _,_ = http.Get("localhost:3000/100000KB-" + ctx.Query("id") + ".txt")
			_, _ = ioutil.ReadFile(Shared.FileTestAddress + "100000KB-" + ctx.Query("id") + ".txt")
			return ctx.SendString("Viewed: " + "100000KB-" + ctx.Query("id"))
		}
		fmt.Print("ID Not Provided")
		return ctx.SendString("100000KB File: ID Not Provided")
	})
	return nil
}
