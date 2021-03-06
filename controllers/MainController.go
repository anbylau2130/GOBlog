package controllers

import (
	"blog/lib"
	"blog/models"
	"strconv"
	"time"

	"strings"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

// MainController main controller
type MainController struct {
	UspController
}

// Home page
func (main *MainController) Home() {
	userinfo := main.GetSession(CurrentUserSession)
	if userinfo == nil {
		main.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
	}
	menuHorizontal := main.GetMenuHorizontal(userinfo.(*UserInfo).Role)
	main.Data["menuHorizontal"] = menuHorizontal
	main.Data["CorpName"] = userinfo.(*UserInfo).Corp.Name
	main.Data["UserName"] = userinfo.(*UserInfo).Operator.RealName
	now := struct {
		Year   int
		Month  int
		Day    int
		Hour   int
		Minute int
		Second int
	}{
		time.Now().Year(),
		int(time.Now().Month()),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(),
	}
	main.Data["Now"] = now
	main.Layout = main.GetTemplatetype() + "/shared/layout.tpl"
	main.TplName = main.GetTemplatetype() + "/mainPages/home.tpl"
}

// Login page
func (main *MainController) Login() {
	isajax := main.GetString("isajax")
	if isajax == "1" {
		username := main.GetString("username")
		password := main.GetString("password")
		user, err := lib.CheckLogin(username, password) //读取operater

		if err == nil {
			userinfo := new(UserInfo)
			userinfo.Operator = user
			userinfo.Menus = models.GetMenusByUser(user)
			userinfo.Privileges = models.GetPrivilegesByUser(user)
			userinfo.Corp = models.GetCorpByUser(user)
			userinfo.Role = models.GetRolesByOperator(user.ID)
			main.SetSession(CurrentUserSession, userinfo)
			logmodel := new(models.SysLoginLog)
			logmodel.Ip = main.GetClientIp()
			logmodel.LoginAgent = main.Ctx.Input.UserAgent()
			logmodel.Success = true
			logmodel.Time = time.Now()
			logmodel.Operator = userinfo.Operator.ID
			logmodel.Add()
			main.SessionRegenerateID()
			user.Session = main.CruSession.SessionID()
			user.Update("Session")
			main.Rsp(true, "登录成功")
			return
		}

		main.Rsp(false, err.Error())
	}
	userinfo := main.GetSession(CurrentUserSession)
	if userinfo != nil {
		main.Ctx.Redirect(302, main.GetHostAddress()+"main/Home")
	}
	main.Layout = main.GetTemplatetype() + "/shared/layout.tpl"
	main.TplName = main.GetTemplatetype() + "/mainPages/login.tpl"
}

func (main *MainController) CheckSSO() {
	userinfo := main.GetSession(CurrentUserSession)
	opSession, ok := userinfo.(*UserInfo)
	if userinfo == nil || !ok {
		main.Data["json"] = &map[string]interface{}{"flag": false, "dateTime": "session is null"}
		main.ServeJSON()
		return
	}
	user := new(models.SysOperator)
	user.ID = opSession.Operator.ID
	_, error := user.Read("ID")
	if error == nil {
		if user.Session == main.CruSession.SessionID() {
			main.Data["json"] = &map[string]interface{}{"flag": true, "dateTime": time.Now()}
			main.ServeJSON()
			return
		}
	}
	main.DestroySession()
	main.Data["json"] = &map[string]interface{}{"flag": false, "dateTime": "有相同用户登陆或同一机器两用户登陆,您已被系统强制退出!"}
	main.ServeJSON()
}

func (main *MainController) Logout() {
	main.DestroySession()
	main.Data["json"] = &map[string]interface{}{"success": true, "data": "/Login"}
	main.ServeJSON()
}

//GetMenuHorizontal 水平菜单
func (main *MainController) GetMenuHorizontal(roles []models.SysRole) []models.MenuNode {
	rolestr := ""
	for _, item := range roles {
		rolestr += strconv.FormatInt(item.ID, 10) + ","
	}
	return new(models.SysMenu).GetMenuHorizontal(0, strings.TrimSuffix(rolestr, ","))
}

// GetMenusVertical 垂直菜单
func (main *MainController) GetMenusVertical() {
	menuOrm := new(models.SysMenu)
	id, errid := main.GetInt64("ID")
	userinfo := main.GetSession(CurrentUserSession)
	roles := ""
	for _, item := range userinfo.(*UserInfo).Role {
		roles += strconv.FormatInt(item.ID, 10) + ","
	}

	if errid == nil {
		rows := menuOrm.GetMenusVertical(id, strings.TrimSuffix(roles, ","))
		main.Data["json"] = &map[string]interface{}{"success": true, "data": &rows}
	} else {
		main.Data["json"] = &map[string]interface{}{"success": false, "data": errid.Error()}
	}
	main.ServeJSON()
}

func (main *MainController) GetPWD() {
	userName := main.GetString("UserName")
	password := main.GetString("Password")
	md5Pwd := lib.Pwdhash(lib.GetPassWord(userName, password))
	main.Ctx.WriteString(md5Pwd)
}
