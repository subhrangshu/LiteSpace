package main

import (
	"LiteSpace/Src"
	"LiteSpace/Src/PrimaryController"
	"fmt"
)

func main() {
	redisSess := Src.RedisSession()
	if redisSess.Ping().String() == "temp bypass" {
		fmt.Print("bypass parameter")
	}
	sqliteSess := Src.SQLiteSession()
	sqliteSess.Ping()

	fiberSess := Src.FiberSession()

	PrimaryController.Controller(fiberSess, sqliteSess, redisSess)
	fiberSess.Listen(":3000")
	//if err := session.Query('Insert int')
}
