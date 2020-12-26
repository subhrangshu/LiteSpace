package Src

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func RedisSession() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}

func SQLiteSession() *sql.DB {
	db, err := sql.Open("sqlite3", "hash.db")
	checkErr(err)

	stmt, err := db.Prepare("create table if not exists hashlocation (id integer primary key , location text)")
	checkErr(err)

	queryReturn, err := stmt.Exec()
	checkErr(err)
	fmt.Print(queryReturn.LastInsertId())

	// insert
	/*stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("as", "abc", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)*/
	return db
}

func FiberSession() *fiber.App {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	return app
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
