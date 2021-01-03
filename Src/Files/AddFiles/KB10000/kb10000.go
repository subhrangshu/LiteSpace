package KB10000

import (
	"LiteSpace/Src/Shared"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"io"
	"os"
)

func KB10000(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/10000kb", func(ctx *fiber.Ctx) error {
		if ctx.Query("id") != "" {
			oldFile, _ := os.Open(Shared.FileBaseAddress + "10000KB.txt")
			newFile, _ := os.Create(Shared.FileTestAddress + "10000KB-" + ctx.Query("id") + ".txt")
			io.Copy(newFile, oldFile)
			oldFile.Close()
			newFile.Close()
			return ctx.SendString("Saved: " + "10000KB-" + ctx.Query("id") + ".txt")
		}
		fmt.Print("ID Not Provided")
		return ctx.SendString("10000KB: ID Not Provided")
	})
	return nil
}
