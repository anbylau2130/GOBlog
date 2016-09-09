package controllers

import (
	"blog/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TreeNodeState struct {
	Opened   bool `json:"opened"`
	Disabled bool `json:"disabled"`
	Selected bool `json:"selected"`
}
type TreeNode struct {
	Id       string     `json:"id"`
	Text     string     `json:"text"`
	Parent   string     `json:"parent"`
	Icon     string     `json:"icon"`
	Children []TreeNode `json:"children"`
	IsLeaf   string
	State    TreeNodeState `json:"state"`
	Tpe      string        `json:"type"`
}

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

func (this *ServiceController) CorpTypeSelect() {
	corpType := new(models.SysCorpType)
	this.Data["json"] = corpType.GetSelect()
	this.ServeJSON()
	return
}

func (this *ServiceController) RoleTree() {
	treenodes := this.PowerTree(0)
	this.Data["json"] = treenodes
	this.ServeJSON()
	return
}

func (this *ServiceController) RoleTreeInit() {
	role, _ := this.GetInt64("Role")
	treenodes := this.PowerTree(0)
	roleMenuModel := new(models.SysRoleMenu)
	rolePrivilegeModel := new(models.SysRolePrivilege)
	for i := 0; i < len(treenodes); i++ {
		if treenodes[i].Tpe == "menu" {
			if count, error := roleMenuModel.Count(orm.NewCondition().And("Menu__exact", strings.TrimPrefix(treenodes[i].Id, "m_")).And("Role__exact", role)); count > 0 && error == nil {
				treenodes[i].State.Selected = true
			}
		}
		if treenodes[i].Tpe == "privilege" {
			if count, error := rolePrivilegeModel.Count(orm.NewCondition().And("Privilege__exact", strings.TrimPrefix(treenodes[i].Id, "p_")).And("Role__exact", role)); count > 0 && error == nil {
				treenodes[i].State.Selected = true
			}
		}
	}
	this.Data["json"] = treenodes
	fmt.Print(treenodes)
	this.ServeJSON()
	return
}

func (this *ServiceController) PowerTree(id int64) []TreeNode {
	menu := new(models.SysMenu)
	menus := menu.GetAll(nil, "ID")
	privilege := new(models.SysPrivilege)

	var treenodes []TreeNode
	for _, item := range menus {
		treenode := TreeNode{}
		treenode.Id = "m_" + strconv.FormatInt(item.ID, 10)
		treenode.Text = item.Name
		treenode.Tpe = "menu"
		if item.Parent == 0 {
			treenode.Parent = "#"
		} else {
			treenode.Parent = "m_" + strconv.FormatInt(item.Parent, 10)
		}
		if count, error := menu.Count(orm.NewCondition().And("Parent__exact", item.ID)); count > 0 && error == nil {
			treenode.IsLeaf = "false"
		} else {
			treenode.IsLeaf = "true"
		}

		if count, error := privilege.Count(orm.NewCondition().And("Menu__exact", item.ID)); count > 0 && error == nil {
			treenode.IsLeaf = "false"
			privileges := privilege.GetAll(orm.NewCondition().And("Menu__exact", item.ID), "ID")
			for _, item := range privileges {
				treenodep := TreeNode{}
				treenodep.Id = "p_" + strconv.FormatInt(item.ID, 10)
				treenodep.Text = item.Name
				treenodep.Tpe = "privilege"
				treenodep.Parent = "m_" + strconv.FormatInt(item.Menu, 10)
				treenodep.IsLeaf = "true"
				treenodes = append(treenodes, treenodep)
			}
		}
		treenodes = append(treenodes, treenode)
	}
	return treenodes
}

// func (this *ServiceController) menusTree(id int64) []TreeNode {
// 	menu := new(models.SysMenu)
// 	privilege := new(models.SysPrivilege)
// 	menus := menu.GetAll(orm.NewCondition().And("parent__exact", id), "ID")
// 	var treenodes []TreeNode
// 	for _, item := range menus {
// 		treenode := TreeNode{}
// 		treenode.Id = strconv.FormatInt(item.ID, 10)
// 		treenode.Text = item.Name
// 		treenode.Tpe = "menu"
// 		if id == 0 {
// 			treenode.Parent = "#"
// 		} else {
// 			treenode.Parent = strconv.FormatInt(id, 10)
// 		}

// 		if count, error := menu.Count(orm.NewCondition().And("parent__exact", item.ID)); count > 0 && error == nil {
// 			treenode.Children = append(treenode.Children, this.menusTree(item.ID)...)
// 		}
// 		if count, error := privilege.Count(orm.NewCondition().And("Menu__exact", item.ID)); count > 0 && error == nil {
// 			treenode.Children = append(treenode.Children, this.privilegeTree(item.ID)...)
// 		}
// 		treenodes = append(treenodes, treenode)
// 	}
// 	return treenodes
// }

// func (this *ServiceController) privilegeTree(menuid int64) []TreeNode {
// 	privilege := new(models.SysPrivilege)
// 	privileges := privilege.GetAll(orm.NewCondition().And("Menu__exact", menuid), "ID")
// 	var treenodes []TreeNode
// 	for _, item := range privileges {
// 		treenode := TreeNode{}
// 		treenode.Id = strconv.FormatInt(item.ID, 10)
// 		treenode.Text = item.Name
// 		treenode.Tpe = "Privilege"
// 		treenodes = append(treenodes, treenode)
// 	}
// 	return treenodes
// }
