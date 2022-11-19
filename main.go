package main

import (
	"androidmarket/global"
	"androidmarket/inlitialize"
)

func main() {
	r := inlitialize.Route()
	global.GlobalDB = inlitialize.GormDb()
	r.Run(":1234")
}
