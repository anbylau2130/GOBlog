package controllers

import (
	"blog/lib"
	"blog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:toor@tcp(127.0.0.1:3306)/usp?charset=utf8")
}

// MainController main controller
type MainController struct {
	UspController
}

// Home page
func (main *MainController) Home() {
	userinfo := main.GetSession(CurrentUserSession)
	if userinfo == nil {
		main.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
	}

	menuHorizontal := main.GetMenuHorizontal()
	main.Data["menuHorizontal"] = menuHorizontal
	main.Layout = main.GetTemplatetype() + "/shared/layout.tpl"
	main.TplName = main.GetTemplatetype() + "/mainPages/home.tpl"
}

// Login page
func (main *MainController) Login() {
	isajax := main.GetString("isajax")
	if isajax == "1" {
		username := main.GetString("username")
		password := main.GetString("password")
		user, err := lib.CheckLogin(username, password) //读取operater

		if err == nil {
			userinfo := new(UserInfo)
			userinfo.Operator = user
			userinfo.Menus = models.GetMenusByUser(user)
			userinfo.Privileges = models.GetPrivilegesByUser(user)
			userinfo.Corp = models.GetCorpByUser(user)
			main.SetSession(CurrentUserSession, userinfo)
			main.Rsp(true, "登录成功")
			return
		}
		main.Rsp(false, err.Error())
	}
	userinfo := main.GetSession(CurrentUserSession)
	if userinfo != nil {
		main.Ctx.Redirect(302, main.GetHostAddress()+"main/Home")
	}
	main.Layout = main.GetTemplatetype() + "/shared/layout.tpl"
	main.TplName = main.GetTemplatetype() + "/mainPages/login.tpl"
}

//GetMenuHorizontal 水平菜单
func (main *MainController) GetMenuHorizontal() []models.MenuNode {
	return new(models.SysMenu).GetMenuHorizontal(0)
}

// GetMenusVertical 垂直菜单
func (main *MainController) GetMenusVertical() {
	menuOrm := new(models.SysMenu)
	id, error := main.GetInt64("ID")
	if error == nil {
		rows := menuOrm.GetMenusVertical(id)
		main.Data["json"] = &map[string]interface{}{"success": true, "data": &rows}
	} else {
		main.Data["json"] = &map[string]interface{}{"success": false, "data": error}
	}
	main.ServeJSON()
}
