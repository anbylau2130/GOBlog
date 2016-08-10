package controllers

import (
	"blog/models"

	_ "github.com/go-sql-driver/mysql"
)

// MainController main controller
type ServiceController struct {
	UspController
}

func (this *ServiceController) RoleSelect() {
	role := new(models.SysRole)
	this.Data["json"] = role.GetSelect()
	this.ServeJSON()
	return
}
