package main

import (
	"blog/lib"
	_ "blog/routers"

	"github.com/astaxie/beego"
)

func main() {
	lib.Initialize()
	beego.Run()
}
