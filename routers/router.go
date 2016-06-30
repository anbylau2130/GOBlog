package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Home")
	beego.Router("/Login", &controllers.MainController{}, "*:Login")
	beego.Router("/main", &controllers.MainController{}, "*:Home")
	beego.Router("/main/Home", &controllers.MainController{}, "*:Home")
	beego.Router("/main/Login", &controllers.MainController{}, "*:Login")
	beego.Router("/main/GetMenuHorizontal", &controllers.MainController{}, "*:GetMenuHorizontal")
	beego.Router("/main/GetMenusVertical", &controllers.MainController{}, "*:GetMenusVertical")
	beego.AutoRouter(&controllers.ObjectController{})
}
