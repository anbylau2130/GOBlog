package admin

import (
	"blog/controllers"
	"blog/models"

	_ "github.com/go-sql-driver/mysql"
)

// MenusController  controller
type CorpController struct {
	controllers.UspController
}

//@MenuH {"name":"系统管理","parent":"0"}
//@MenuH {"name":"系统设置","parent":"系统管理"}
//@MenuV {"name":"公司列表","parent":"系统设置"}
func (this *CorpController) List() {
	this.Layout = this.GetTemplatetype() + "/shared/layout.tpl"
	this.TplName = this.GetTemplatetype() + "/adminPages/corp.tpl"
}

func (this *CorpController) GetList() {
	pageindex, _ := this.GetInt64("start")
	pagesize, _ := this.GetInt64("length")
	draw, _ := this.GetInt64("draw")
	model := new(models.SysCorp)
	models, _ := model.Getlist(nil, pageindex, pagesize, "ID")
	count, _ := model.Count(nil)
	this.Data["json"] = &map[string]interface{}{"draw": draw, "recordsTotal": count, "recordsFiltered": count, "data": models}
	this.ServeJSON()
}

func (this *CorpController) GetModel() {
	model := models.SysCorp{}
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

func (this *CorpController) Add() {
	if this.IsAjax() {
		model := new(models.SysCorp)
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
	this.TplName = this.GetTemplatetype() + "/adminPages/addcorp.tpl"
}

func (this *CorpController) Del() {
	model := new(models.SysCorp)
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

func (this *CorpController) Edit() {
	model := new(models.SysCorp)
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
	this.TplName = this.GetTemplatetype() + "/adminPages/editcorp.tpl"

}
