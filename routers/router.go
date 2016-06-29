package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Home")
	beego.Router("/Home", &controllers.MainController{}, "*:Home")
	beego.Router("/Login", &controllers.MainController{}, "*:Login")
	beego.Router("/GetMenuHorizontal", &controllers.MainController{}, "*:GetMenuHorizontal")
	beego.Router("/GetMenusVertical", &controllers.MainController{}, "*:GetMenusVertical")
}
