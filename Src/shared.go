package Src

import (
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

type Sess struct {
	RedisSess  *redis.Client
	SqliteSess *sql.DB
	FiberSess  *fiber.App
}
