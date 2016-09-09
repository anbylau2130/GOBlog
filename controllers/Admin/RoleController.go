package admin

import (
	"blog/controllers"
	"blog/models"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// MenusController  controller
type RoleController struct {
	controllers.UspController
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"角色管理","parent":"系统设置"}
func (this *RoleController) List() {
	this.Layout = this.GetTemplatetype() + "/shared/layout.tpl"
	this.TplName = this.GetTemplatetype() + "/adminPages/role.tpl"
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"角色管理","parent":"系统设置"}
//@Privilege {"name":"获取角色列表","parent":"角色管理"}
func (this *RoleController) GetList() {
	pageindex, _ := this.GetInt64("start")
	pagesize, _ := this.GetInt64("length")
	draw, _ := this.GetInt64("draw")
	model := new(models.SysRole)
	models, _ := model.Getlist(nil, pageindex, pagesize, "ID")
	count, _ := model.Count(nil)
	this.Data["json"] = &map[string]interface{}{"draw": draw, "recordsTotal": count, "recordsFiltered": count, "data": models}
	this.ServeJSON()
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"角色管理","parent":"系统设置"}
//@Privilege {"name":"获取角色","parent":"角色管理"}
func (this *RoleController) GetModel() {
	model := models.SysRole{}
	roleMenuModel := models.SysRoleMenu{}
	rolePrivilegeModel := models.SysRolePrivilege{}
	var menus []int64
	var privileges []int64
	id, error := this.GetInt64("id")
	if id != 0 && error == nil {
		model.ID = id
		model.Read("ID")
		roleMenus := roleMenuModel.GetModelByRole(id)
		rolePriviliges := rolePrivilegeModel.GetModelByRole(id)
		for _, v := range roleMenus {
			menus = append(menus, v.ID)
		}
		for _, v := range rolePriviliges {
			privileges = append(privileges, v.ID)
		}
		var viewModel struct {
			Name       string
			Remark     string
			ID         int64
			Menus      []int64
			Privileges []int64
		}
		viewModel.Name = model.Name
		viewModel.Remark = model.Remark
		viewModel.ID = model.ID
		viewModel.Menus = menus
		viewModel.Privileges = privileges
		this.Data["json"] = &map[string]interface{}{"success": true, "data": viewModel}
		this.ServeJSON()
		return
	}
	this.Rsp(false, error.Error())
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"角色管理","parent":"系统设置"}
//@Privilege {"name":"新增角色","parent":"角色管理"}
func (this *RoleController) Add() {
	if this.IsAjax() {
		userinfo := this.GetSession(controllers.CurrentUserSession)
		opSession, _ := userinfo.(*controllers.UserInfo)
		model := new(models.SysRole)
		var name, remark, menus, privileges string
		if name = this.GetString("Name"); name == "" {
			this.Rsp(false, "角色名不能为空！")
			return
		}
		if remark = this.GetString("Remark"); remark == "" {
			this.Rsp(false, "描述信息不能为空！")
			return
		}
		menus = this.GetString("Menus")
		privileges = this.GetString("Privileges")

		res, error := model.Add(opSession.Corp.ID, opSession.Operator.ID, 0, name, remark, strings.TrimSuffix(menus, ","), strings.TrimSuffix(privileges, ","))
		if error == nil {
			if res.IsSuccess == "true" {
				this.Rsp(true, "Success")
			} else {
				this.Rsp(false, res.ProcMsg)
			}
		} else {
			this.Rsp(false, error.Error())
		}
		return
	}
	this.Layout = this.GetTemplatetype() + "/shared/layout.tpl"
	this.TplName = this.GetTemplatetype() + "/adminPages/addrole.tpl"
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"角色管理","parent":"系统设置"}
//@Privilege {"name":"删除角色","parent":"角色管理"}
func (this *RoleController) Del() {
	model := new(models.SysRole)
	if this.IsAjax() {
		id, error := this.GetInt64("id")
		if id != 0 && error == nil {
			model.ID = id
			count, error := model.Delete()
			if error != nil || count == 0 {
				this.Rsp(false, error.Error())
			}
		}
	}
	this.Rsp(true, "数据删除成功!")
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"角色管理","parent":"系统设置"}
//@Privilege {"name":"编辑角色","parent":"角色管理"}
func (this *RoleController) Edit() {
	model := new(models.SysRole)
	//ajax提交
	if this.IsAjax() {
		userinfo := this.GetSession(controllers.CurrentUserSession)
		opSession, _ := userinfo.(*controllers.UserInfo)
		model := new(models.SysRole)
		var name, remark, menus, privileges string
		if name = this.GetString("Name"); name == "" {
			this.Rsp(false, "角色名不能为空！")
			return
		}
		if remark = this.GetString("Remark"); remark == "" {
			this.Rsp(false, "描述信息不能为空！")
			return
		}
		role, error := this.GetInt64("ID")
		if error == nil && role == 0 {
			this.Rsp(false, "未获取RoleID")
			return
		}
		menus = this.GetString("Menus")
		privileges = this.GetString("Privileges")

		res, error := model.Update(
			role,
			opSession.Operator.ID, name, remark,
			strings.TrimSuffix(menus, ","),
			strings.TrimSuffix(privileges, ","))
		if error == nil {
			if res.IsSuccess == "true" {
				this.Rsp(true, "Success")
			} else {
				this.Rsp(false, res.ProcMsg)
			}
		} else {
			this.Rsp(false, error.Error())
		}
		return
	}
	//页面访问
	id, error := this.GetInt64("id")
	if id != 0 && error == nil {
		model.ID = id
		model.Read("ID")
	}
	this.Data["Model"] = model
	this.Layout = this.GetTemplatetype() + "/shared/layout.tpl"
	this.TplName = this.GetTemplatetype() + "/adminPages/editrole.tpl"

}
