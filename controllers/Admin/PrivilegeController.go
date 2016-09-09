package admin

import (
	"blog/controllers"
	"blog/models"

	_ "github.com/go-sql-driver/mysql"
)

// PrivilegeController  controller
type PrivilegeController struct {
	controllers.UspController
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"权限管理","parent":"系统设置"}
func (this *PrivilegeController) List() {
	this.Layout = this.GetTemplatetype() + "/shared/layout.tpl"
	this.TplName = this.GetTemplatetype() + "/adminPages/privilege.tpl"
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"权限管理","parent":"系统设置"}
//@Privilege {"name":"获取权限列表","parent":"权限管理"}
func (this *PrivilegeController) GetList() {
	pageindex, _ := this.GetInt64("start")
	pagesize, _ := this.GetInt64("length")
	draw, _ := this.GetInt64("draw")
	model := new(models.SysPrivilege)
	models, _ := model.Getlist(nil, pageindex, pagesize, "ID")
	count, _ := model.Count(nil)
	this.Data["json"] = &map[string]interface{}{"draw": draw, "recordsTotal": count, "recordsFiltered": count, "data": models}
	this.ServeJSON()
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"权限管理","parent":"系统设置"}
//@Privilege {"name":"获取权限","parent":"权限管理"}
func (this *PrivilegeController) GetModel() {
	model := models.SysPrivilege{}
	id, error := this.GetInt64("id")
	if id != 0 && error == nil {
		model.ID = id
		model.Read("ID")
		this.Data["json"] = &map[string]interface{}{"success": true, "data": model}
		this.ServeJSON()
		return
	}
	this.Rsp(false, error.Error())
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"权限管理","parent":"系统设置"}
//@Privilege {"name":"新增权限","parent":"权限管理"}
func (this *PrivilegeController) Add() {
	if this.IsAjax() {
		model := new(models.SysPrivilege)
		if error := this.ParseForm(model); error != nil {
			this.Rsp(false, error.Error())
			return
		}
		count, error := model.Add()
		if error == nil && count > 0 {
			this.Rsp(true, "Success")
		} else {
			this.Rsp(false, "")
		}
		return
	}
	this.Layout = this.GetTemplatetype() + "/shared/layout.tpl"
	this.TplName = this.GetTemplatetype() + "/adminPages/addprivilege.tpl"
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"权限管理","parent":"系统设置"}
//@Privilege {"name":"删除权限","parent":"权限管理"}
func (this *PrivilegeController) Del() {
	model := new(models.SysPrivilege)
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
//@MenuV {"name":"权限管理","parent":"系统设置"}
//@Privilege {"name":"修改权限","parent":"权限管理"}
func (this *PrivilegeController) Edit() {
	model := new(models.SysPrivilege)
	//ajax提交
	if this.IsAjax() {
		if error := this.ParseForm(model); error != nil {
			this.Rsp(false, error.Error())
			return
		}
		count, error := model.Update()
		if error == nil && count > 0 {
			this.Rsp(true, "数据编辑成功!")
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
	this.TplName = this.GetTemplatetype() + "/adminPages/editprivilege.tpl"

}
