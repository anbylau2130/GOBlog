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
//@MenuH {"name":"菜单管理","parent":"系统管理"}
//@MenuV {"name":"菜单列表","parent":"菜单管理"}
func (menus *MenusController) List() {
	menus.Layout = menus.GetTemplatetype() + "/shared/layout.tpl"
	menus.TplName = menus.GetTemplatetype() + "/adminPages/menus.tpl"
}

func (menus *MenusController) GetList() {
	menuModel := new(models.SysMenu)
	models, count := menuModel.Getlist(nil, 1, 10, "ID")
	menus.Data["json"] = &map[string]interface{}{"draw": 1, "recordsTotal": count, "recordsFiltered": count, "data": models}
	menus.ServeJSON()
}

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

func (menus *MenusController) create(model models.SysMenu) *models.SysMenu {

	if v, error := model.Add(); error == nil {
		model.ID = v
	}
	return &model
}

func (menus *MenusController) update(model models.SysMenu) *models.SysMenu {

	if v, error := model.Update(); error == nil && v == 0 {
		return nil
	}
	return &model
}

func (menus *MenusController) delete(model models.SysMenu) *models.SysMenu {

	if v, error := model.Delete(); error == nil && v == 0 {
		return nil
	}
	return &model
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"菜单管理","parent":"系统管理"}
//@MenuV {"name":"菜单明细","parent":"菜单管理"}
func (menus *MenusController) Detail() {
	menus.Layout = menus.GetTemplatetype() + "/shared/layout.tpl"
	menus.TplName = menus.GetTemplatetype() + "/adminPages/menuDetail.tpl"
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"菜单管理","parent":"系统管理"}
//@MenuV {"name":"菜单编辑","parent":"菜单管理"}
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
			menus.Rsp(true, "Success")
		} else {
			menus.Rsp(false, "")
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

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"菜单管理","parent":"系统管理"}
//@MenuV {"name":"菜单编辑","parent":"菜单管理"}
//@Privilege {"name":"操作","parent":"菜单编辑"}
func (menus *MenusController) CreateMenu() {
	lib.RegisterMenus(&MenusController{})
}
