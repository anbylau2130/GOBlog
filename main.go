package main

import (
	"blog/controllers/admin"
	"blog/lib"
	_ "blog/routers"
	"mime"
	"os"

	"github.com/astaxie/beego"
)

func main() {
	Initialize()
	beego.SetStaticPath("/data", "data")
	beego.Run()
}

func Initialize() {
	mime.AddExtensionType(".css", "text/css")
	//判断初始化参数
	initArgs()

	lib.Connect()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			lib.Syncdb()
			os.Exit(0)
		}
		if v == "-syncmenu" {
			lib.RegisterMenus(&admin.MenusController{})
			os.Exit(0)
		}
	}
}
