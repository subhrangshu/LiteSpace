package AddEntry

import (
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"os/exec"
)

func AddEntry(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/add", func(ctx *fiber.Ctx) error {
		out, err := exec.Command("uuidgen").Output()
		res, err := redSess.Do("SET", out, "Repo Man").String()
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
