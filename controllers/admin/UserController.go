package admin

import (
	"blog/controllers"
	"blog/models"
)

// MenusController  controller
type UserController struct {
	controllers.UspController
}

func (user *UserController) Profile() {
	sessionModel := user.GetSession(controllers.CurrentUserSession)
	userinfo, ok := sessionModel.(*controllers.UserInfo)
	if ok {
		user.Data["LoginName"] = userinfo.Operator.LoginName
		user.Data["RealName"] = userinfo.Operator.RealName
		user.Data["UImage"] = userinfo.Operator.UImage
		user.Data["Operator"] = userinfo.Operator
	}
	user.Layout = user.GetTemplatetype() + "/shared/layout.tpl"
	user.TplName = user.GetTemplatetype() + "/adminPages/profile.tpl"
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"员工管理","parent":"系统设置"}
func (this *UserController) List() {
	this.Layout = this.GetTemplatetype() + "/shared/layout.tpl"
	this.TplName = this.GetTemplatetype() + "/adminPages/users.tpl"
}

func (this *UserController) GetList() {
	pageindex, _ := this.GetInt64("start")
	pagesize, _ := this.GetInt64("length")
	draw, _ := this.GetInt64("draw")
	model := new(models.SysOperator)
	models, _ := model.Getlist(nil, pageindex, pagesize, "ID")
	count, _ := model.Count(nil)
	this.Data["json"] = &map[string]interface{}{"draw": draw, "recordsTotal": count, "recordsFiltered": count, "data": models}
	this.ServeJSON()
}

func (this *UserController) GetModel() {
	model := models.SysOperator{}
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

func (this *UserController) Add() {
	if this.IsAjax() {
		userinfo := this.GetSession(controllers.CurrentUserSession)
		opSession, _ := userinfo.(*controllers.UserInfo)
		model := new(models.SysOperator)

		var LoginName, RealName, Password, Role, PasswordConfirm string
		if RealName = this.GetString("RealName"); RealName == "" {
			this.Rsp(false, "姓名不能为空！")
			return
		}
		if Password = this.GetString("Password"); Password == "" {
			this.Rsp(false, "密码不能为空！")
			return
		}
		if PasswordConfirm = this.GetString("PasswordConfirm"); PasswordConfirm == "" {
			this.Rsp(false, "密码确认不能为空！")
			return
		}
		if Role = this.GetString("Role"); Role == "" {
			this.Rsp(false, "角色不能为空！")
			return
		}
		if LoginName := this.GetString("LoginName"); LoginName == "" {
			this.Rsp(false, "登录名不能为空！")
			return
		}
		if Password != PasswordConfirm {
			this.Rsp(false, "两次输入的密码必须相同!")
			return
		}
		_, error := model.Add(opSession.Corp.ID, opSession.Operator.ID, LoginName, RealName, Password, "", "", "", Role)
		if error == nil {
			this.Rsp(true, "Success")
		} else {
			this.Rsp(false, error.Error())
		}
		return
	}

	this.Layout = this.GetTemplatetype() + "/shared/layout.tpl"
	this.TplName = this.GetTemplatetype() + "/adminPages/adduser.tpl"
}

func (this *UserController) Del() {
	model := new(models.SysOperator)
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

func (this *UserController) Edit() {
	model := new(models.SysOperator)
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
	this.TplName = this.GetTemplatetype() + "/adminPages/edituser.tpl"

}
