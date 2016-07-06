package routers

import (
	"blog/controllers"
	"blog/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{}, "*:Home")
	beego.Router("/Login", &controllers.MainController{}, "*:Login")
	beego.Router("/Logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/main", &controllers.MainController{}, "*:Home")
	beego.Router("/main/Home", &controllers.MainController{}, "*:Home")
	beego.Router("/main/Login", &controllers.MainController{}, "*:Login")
	beego.Router("/main/GetMenuHorizontal", &controllers.MainController{}, "*:GetMenuHorizontal")
	beego.Router("/main/GetMenusVertical", &controllers.MainController{}, "*:GetMenusVertical")
	beego.AutoPrefix("/admin", &admin.MenusController{})
	beego.AutoPrefix("/admin", &admin.UserController{})
	//beego.Include(&controllers.MainController{})
	//beego.AutoPrefix("/Admin",&controllers.AdminController{})

}
