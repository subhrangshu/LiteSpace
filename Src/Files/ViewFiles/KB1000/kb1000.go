package KB1000

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func KB1000(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/1000kb", func(ctx *fiber.Ctx) error {
		if ctx.Query("id") != "" {
			_, _ = http.Get("localhost:3000/1000KB-" + ctx.Query("id") + ".txt")
			// _, _ = ioutil.ReadFile(Shared.FileTestAddress + "1000KB-" + ctx.Query("id") + ".txt")
			return ctx.SendString("Viewed: " + "1000KB-" + ctx.Query("id"))
		}
		fmt.Print("ID Not Provided")
		return ctx.SendString("1000KB File: ID Not Provided")
	})
	return nil
}
