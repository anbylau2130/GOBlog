package admin

import "blog/controllers"

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
