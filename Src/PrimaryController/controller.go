package PrimaryController

import (
	"LiteSpace/Src"
	"fmt"
)

func Controller(sess Src.Sess) {
	sess.RedisSess.Ping()
	fmt.Println("Session Initiated")
}
