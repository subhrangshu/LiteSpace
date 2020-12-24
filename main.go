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

	sess := Src.Sess{
		RedisSess:  redisSess,
		SqliteSess: sqliteSess,
		FiberSess:  fiberSess,
	}

	PrimaryController.Controller(sess)
	//if err := session.Query('Insert int')
}
