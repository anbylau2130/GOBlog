package controllers

import (
	"blog/models"

	"github.com/astaxie/beego"
)

type UserInfo struct {
	Operator   models.SysOperator
	Corp       models.SysCorp
	Menus      []models.SysMenu
	Privileges []models.SysPrivilege
}

// UspController 定义
type UspController struct {
	beego.Controller
	Templatetype string //ui template type
	HostAddress  string //当前主机地址
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

// GetTemplatetype 获取模板类型
func (usp *UspController) GetHostAddress() string {
	hostaddress := beego.AppConfig.String("HttpHostAddress")
	port := beego.AppConfig.String("HttpHostAddressPort")
	if hostaddress == "" {
		hostaddress = "127.0.0.1"
	}
	if port == "" {
		port = "8080"
	}
	return hostaddress + ":" + port
}
