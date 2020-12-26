package DeleteFiles

import (
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func DeleteFiles(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/delete", func(ctx *fiber.Ctx) error {
		stmt, err := sqlSess.Prepare(`DELETE FROM hashlocation WHERE id == 6`)
		checkErr(err)
		res, err := stmt.Exec()
		checkErr(err)
		id, err := res.RowsAffected()
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
