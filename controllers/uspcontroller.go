package controllers

import "github.com/astaxie/beego"

// UspController 定义
type UspController struct {
	beego.Controller
	Templatetype string //ui template type
}

//Rsp Json输出
func (usp *UspController) Rsp(status bool, str string) {
	usp.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	usp.ServeJSON()
}

// GetTemplatetype 获取模板类型
func (usp *UspController) GetTemplatetype() string {
	templatetype := beego.AppConfig.String("template_type")
	if templatetype == "" {
		templatetype = "metro"
	}
	return templatetype
}
