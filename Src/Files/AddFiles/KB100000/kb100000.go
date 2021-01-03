package KB100000

import (
	"LiteSpace/Src/Files"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"io"
	"os"
)

func KB100000(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/100000kb", func(ctx *fiber.Ctx) error {
		if ctx.Query("id") != "" {
			oldFile, _ := os.Open(Files.FileBaseAddress + "100000KB.txt")
			newFile, _ := os.Create(Files.FileTestAddress + "100000KB-" + ctx.Query("id") + ".txt")
			io.Copy(newFile, oldFile)
			oldFile.Close()
			newFile.Close()
			return ctx.SendString("Saved: " + "100000KB-" + ctx.Query("id") + ".txt")
		}
		fmt.Print("ID Not Provided")
		return ctx.SendString("100000KB: ID Not Provided")
	})
	return nil
}
