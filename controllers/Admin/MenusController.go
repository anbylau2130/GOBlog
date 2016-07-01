package admin

import (
	"blog/controllers"

	_ "github.com/go-sql-driver/mysql"
)

func init() {

}

// MenusController  controller
type MenusController struct {
	controllers.UspController
}

//@MenuH name:菜单管理;parent:0
func (menus *MenusController) List() {
	menus.Layout = menus.GetTemplatetype() + "/shared/layout.tpl"
	menus.TplName = menus.GetTemplatetype() + "/adminPages/Menus.tpl"
}

//@MenuH name:菜单明细;parent:菜单管理
func (menus *MenusController) Detail() {
	menus.Layout = menus.GetTemplatetype() + "/shared/layout.tpl"
	menus.TplName = menus.GetTemplatetype() + "/adminPages/MenuDetail.tpl"
}

//@MenuV name:菜单编辑;parent:菜单管理
func (menus *MenusController) Edit() {
	menus.Layout = menus.GetTemplatetype() + "/shared/layout.tpl"
	menus.TplName = menus.GetTemplatetype() + "/adminPages/MenuDetail.tpl"
}

//@Privilege name:操作;parent:菜单编辑
func (menus *MenusController) DoSth() {

}
