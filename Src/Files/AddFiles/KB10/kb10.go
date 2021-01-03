package KB10

import (
	"LiteSpace/Src/Files"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"io"
	"os"
)

func KB10(fibSess *fiber.App, sqlSess *sql.DB, redSess *redis.Client, router fiber.Router) error {
	router.Get("/10kb", func(ctx *fiber.Ctx) error {
		if ctx.Query("id") != "" {
			oldFile, _ := os.Open(Files.FileBaseAddress + "10KB.txt")
			newFile, _ := os.Create(Files.FileTestAddress + "10KB-" + ctx.Query("id") + ".txt")
			io.Copy(newFile, oldFile)
			oldFile.Close()
			newFile.Close()
			return ctx.SendString("Saved: " + "10KB-" + ctx.Query("id") + ".txt")
		}
		fmt.Print("ID Not Provided")
		return ctx.SendString("10KB: D Not Provided")
	})
	return nil
}
