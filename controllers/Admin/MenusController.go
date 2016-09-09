package admin

import (
	"blog/controllers"
	"blog/lib"
	"blog/models"

	_ "github.com/go-sql-driver/mysql"
)

// MenusController  controller
type MenusController struct {
	controllers.UspController
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"菜单管理","parent":"系统设置"}
func (menus *MenusController) List() {
	menus.Layout = menus.GetTemplatetype() + "/shared/layout.tpl"
	menus.TplName = menus.GetTemplatetype() + "/adminPages/menus.tpl"
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"菜单管理","parent":"系统设置"}
//@Privilege {"name":"获取菜单列表","parent":"菜单管理"}
func (menus *MenusController) GetList() {
	pageindex, _ := menus.GetInt64("start")
	pagesize, _ := menus.GetInt64("length")
	draw, _ := menus.GetInt64("draw")
	menuModel := new(models.SysMenu)
	models, _ := menuModel.Getlist(nil, pageindex, pagesize, "ID")
	count, _ := menuModel.Count(nil)
	menus.Data["json"] = &map[string]interface{}{"draw": draw, "recordsTotal": count, "recordsFiltered": count, "data": models}
	menus.ServeJSON()
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"菜单管理","parent":"系统设置"}
//@Privilege {"name":"获取菜单","parent":"菜单管理"}
func (menus *MenusController) GetModel() {
	menuModel := models.SysMenu{}
	id, error := menus.GetInt64("id")
	if id != 0 && error == nil {
		menuModel.ID = id
		menuModel.Read("ID")
		menus.Data["json"] = &map[string]interface{}{"success": true, "data": menuModel}
		menus.ServeJSON()
		return
	}
	menus.Rsp(false, error.Error())
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"菜单管理","parent":"系统设置"}
//@Privilege {"name":"新增菜单","parent":"菜单管理"}
func (menus *MenusController) Add() {
	if menus.IsAjax() {
		menuModel := new(models.SysMenu)
		if error := menus.ParseForm(menuModel); error != nil {
			menus.Rsp(false, error.Error())
			return
		}
		count, error := menuModel.Add()
		if error == nil && count > 0 {
			menus.Rsp(true, "Success")
		} else {
			menus.Rsp(false, "")
		}
		return
	}
	menus.Layout = menus.GetTemplatetype() + "/shared/layout.tpl"
	menus.TplName = menus.GetTemplatetype() + "/adminPages/addmenu.tpl"
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"菜单管理","parent":"系统设置"}
//@Privilege {"name":"删除菜单","parent":"菜单管理"}
func (menus *MenusController) Del() {
	menuModel := new(models.SysMenu)
	if menus.IsAjax() {
		id, error := menus.GetInt64("id")
		if id != 0 && error == nil {
			menuModel.ID = id
			count, error := menuModel.Delete()
			if error != nil || count == 0 {
				menus.Rsp(false, error.Error())
			}
		}
	}
	menus.Rsp(true, "数据删除成功!")
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"菜单管理","parent":"系统设置"}
//@Privilege {"name":"编辑菜单","parent":"菜单管理"}
func (menus *MenusController) Edit() {
	menuModel := new(models.SysMenu)
	//ajax提交
	if menus.IsAjax() {
		if error := menus.ParseForm(menuModel); error != nil {
			menus.Rsp(false, error.Error())
			return
		}
		count, error := menuModel.Update()
		if error == nil && count > 0 {
			menus.Rsp(true, "数据编辑成功!")
		} else {
			menus.Rsp(false, error.Error())
		}
		return
	}
	//页面访问
	id, error := menus.GetInt64("id")
	if id != 0 && error == nil {
		menuModel.ID = id
		menuModel.Read("ID")
	}
	menus.Data["Model"] = menuModel
	menus.Layout = menus.GetTemplatetype() + "/shared/layout.tpl"
	menus.TplName = menus.GetTemplatetype() + "/adminPages/editmenu.tpl"

}

func (menus *MenusController) CreateMenu() {
	lib.RegisterMenus(&MenusController{})
}
