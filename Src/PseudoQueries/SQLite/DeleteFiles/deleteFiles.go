package DeleteFiles

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func DeleteFiles(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/delete", func(ctx *fiber.Ctx) error {
		if ctx.Query("id") != "" {
			stmt, _ := sqlSess.Prepare(`DELETE FROM hashlocation WHERE id == ?`)
			stmt.Exec(ctx.Query("id"))
			return ctx.SendString(ctx.Query("id"))
		}
		fmt.Print("ID Not Provided")
		return ctx.SendString("ID Not Provided")
	})
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
