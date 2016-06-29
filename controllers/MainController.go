package controllers

import (
	"blog/lib"
	"blog/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	CurrentUserSession = "CURRENT_USER"
	CurrentCorpSession = "CurrentCorpSession"

	CurrentAccessMenusSession      = "CURRENT_ACCESS_MENUS"
	CurrentAccessPrivilegesSession = "Current_ACCESS_PRIVILEGES_SESSION"
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
	menuHorizontal := main.GetMenuHorizontal()
	main.Data["menuHorizontal"] = menuHorizontal
	main.Layout = "metro/shared/layout.tpl"
	main.TplName = "metro/pages/home.tpl"
}

// Login page
func (main *MainController) Login() {
	isajax := main.GetString("isajax")
	if isajax == "1" {
		username := main.GetString("username")
		password := main.GetString("password")

		user, err := lib.CheckLogin(username, password)
		if err == nil {
			main.SetSession(CurrentUserSession, user)
			//accesslist, _ := GetAccessList(user.Id)
			//main.SetSession("CURRENT_MENULIST", accesslist)
			main.Rsp(true, "登录成功")
			return
		}
		main.Rsp(false, err.Error())
	}
	userinfo := main.GetSession("userinfo")
	if userinfo != nil {
		main.Ctx.Redirect(302, "/public/index")
	}
	main.Layout = main.GetTemplatetype() + "/shared/layout.tpl"
	main.TplName = main.GetTemplatetype() + "/pages/login.tpl"
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
