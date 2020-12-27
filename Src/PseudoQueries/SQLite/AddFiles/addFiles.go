package AddFiles

import (
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"strconv"
)

func AddFiles(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/add", func(ctx *fiber.Ctx) error {
		stmt, _ := sqlSess.Prepare("INSERT INTO hashlocation (location) values(?)")
		res, _ := stmt.Exec(rand.Float64())
		id, _ := res.LastInsertId()
		return ctx.SendString(strconv.FormatInt(id, 10))
	})
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
