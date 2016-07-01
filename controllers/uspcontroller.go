package controllers

import (
	"blog/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

const (
	CurrentUserSession = "CURRENT_USER"
	ControllerPath     = "blog/controllers/"
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
}

func init() {

	//验证权限
	AccessRegister()
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

func AccessRegister() {
	var Check = func(ctx *context.Context) {
		user_auth_type, _ := strconv.Atoi(beego.AppConfig.String("user_auth_type"))
		rbac_auth_gateway := beego.AppConfig.String("rbac_auth_gateway")

		if user_auth_type > 0 {
			params := strings.Split(strings.ToLower(ctx.Request.RequestURI), "/")
			//area,controller,action
			if CheckAccess(params) {
				uinfo := ctx.Input.Session(CurrentUserSession)

				if uinfo == nil {
					ctx.Redirect(302, rbac_auth_gateway)
				}
				//admin用户不用认证权限
				adminuser := beego.AppConfig.String("rbac_admin_user")
				userinfo := uinfo.(*UserInfo)
				if userinfo.Operator.LoginName == adminuser {
					return
				}

				if user_auth_type == 1 {
					if userinfo.Menus == nil || userinfo.Privileges == nil {
						ctx.Redirect(302, rbac_auth_gateway)
					}
				} else if user_auth_type == 2 {
					userinfo.Menus, userinfo.Privileges = GetAccessList(userinfo.Operator)
				}

				ret := AccessDecision(params, *userinfo)
				if !ret {
					ctx.Output.JSON(&map[string]interface{}{"status": false, "info": "权限不足"}, true, false)
				}
			}
		}
	}
	beego.InsertFilter("/*", beego.BeforeRouter, Check)
}

func GetAccessList(user models.SysOperator) ([]models.SysMenu, []models.SysPrivilege) {
	return models.GetMenusByUser(user), models.GetPrivilegesByUser(user)
}
func AccessDecision(params []string, userInfo UserInfo) bool {
	if CheckAccess(params) {
		s := fmt.Sprintf("%s/%s/%s", params[1], params[2], params[3])
		if len(userInfo.Menus) < 1 || len(userInfo.Privileges) < 1 {
			return false
		}

		for _, item := range userInfo.Menus {
			if item.URL == s {
				return true
			}
		}
		for _, item := range userInfo.Privileges {

			if item.Url == s {
				return true
			}
		}

	} else {
		return true
	}
	return false
}
func CheckAccess(params []string) bool {
	if len(params) < 3 {
		return false
	}
	for _, nap := range strings.Split(beego.AppConfig.String("not_auth_package"), ",") {
		if params[1] == nap {
			return false
		}
	}
	return true
}
