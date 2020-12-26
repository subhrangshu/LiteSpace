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
		stmt, err := sqlSess.Prepare("INSERT INTO hashlocation (location) values(?)")
		checkErr(err)
		res, err := stmt.Exec(rand.Float64())
		checkErr(err)
		id, err := res.LastInsertId()
		checkErr(err)
		return ctx.SendString(strconv.FormatInt(id, 10))
	})
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
