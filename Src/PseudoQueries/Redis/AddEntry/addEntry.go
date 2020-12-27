package AddEntry

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func AddEntry(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/add", func(ctx *fiber.Ctx) error {
		if ctx.Query("id") != "" {
			res, _ := redSess.Do("SET", ctx.Query("id"), ctx.Query("id")).String()
			return ctx.SendString(res)
			// => Hello john
		}
		fmt.Print("ID Not Provided")
		//out, err := exec.Command("uuidgen").Output()
		res, err := redSess.Do("SET", "out", "Repo Man").String()
		checkErr(err)
		return ctx.SendString(res)
	})
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
