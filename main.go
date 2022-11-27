package main

import (
	"androidmarket/global"
	"androidmarket/initialize"
)

func main() {
	r := initialize.Route()
	global.GlobalDB = initialize.GormDb()
	r.Run(":1234")
}
