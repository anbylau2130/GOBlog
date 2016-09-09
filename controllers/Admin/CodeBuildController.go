package admin

import (
	"blog/controllers"

	_ "github.com/go-sql-driver/mysql"
)

// ChartRoomController  controller
type CodeBuildController struct {
	controllers.UspController
}

func (this *CodeBuildController) CodeBuild() {

	this.Layout = this.GetTemplatetype() + "/shared/layout.tpl"
	this.TplName = this.GetTemplatetype() + "/adminPages/codebuild.tpl"
}

func (this *CodeBuildController) DBBuild() {

	this.Layout = this.GetTemplatetype() + "/shared/layout.tpl"
	this.TplName = this.GetTemplatetype() + "/adminPages/tablecreate.tpl"
}
