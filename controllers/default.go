package controllers

import (
	"blog/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// MainController main controller
type MainController struct {
	UspController
}

// func (this *MainController) Rsp(status bool, str string) {
// 	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
// 	this.ServeJSON()
// }

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:toor@tcp(127.0.0.1:3306)/usp?charset=utf8")
}

// Home page
func (main *MainController) Home() {
	main.Layout = "metro/shared/layout.tpl"
	main.TplName = "metro/pages/home.tpl"
}

//GetMenuHorizontal 水平菜单
func (main *MainController) GetMenuHorizontal() {
	menuOrm := new(models.SysMenu)
	rows := menuOrm.GetMenuHorizontal(0)
	main.Data["json"] = &map[string]interface{}{"success": true, "data": &rows}
	main.ServeJSON()
}

// GetMenusVertical 垂直菜单
func (main *MainController) GetMenusVertical() {
	menuOrm := new(models.SysMenu)
	rows := menuOrm.GetMenusVertical(0)
	main.Data["json"] = &map[string]interface{}{"success": true, "rows": &rows}
	main.ServeJSON()
}
