//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	orm.RegisterModel(new(SysMenu))
	//orm.RegisterModelWithPrefix("", new(SysMenu))
}

type MenuNode struct {
	ID           int64
	Iocn         string
	Name         string
	URL          string
	Direction    bool
	ChildrenNode []MenuNode
}
type SysMenu struct {
	ID         int64     `orm:"column(ID);pk;unique;default();index;"`
	Direction  bool      `orm:"column(Direction);not null;default();"`
	Parent     int64     `orm:"column(Parent);default();"`
	Name       string    `orm:"column(Name);size(0);default((''));"`
	Icon       string    `orm:"column(Icon);size(0);default((''));"`
	Clazz      string    `orm:"column(Clazz);size(0);default((''));"`
	Area       string    `orm:"column(Area);size(0);default((''));"`
	Controller string    `orm:"column(Controller);size(0);default((''));"`
	Method     string    `orm:"column(Method);size(0);default((''));"`
	Parameter  string    `orm:"column(Parameter);size(0);default((''));"`
	URL        string    `orm:"column(Url);size(0);default((''));"`
	Reserve    string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark     string    `orm:"column(Remark);not null;size(0);default();"`
	Creator    int64     `orm:"column(Creator);default();"`
	CreateTime time.Time `orm:"column(CreateTime);auto_now_add;type(datetime);default((getdate()));"`
	Auditor    int64     `orm:"column(Auditor);not null;default();"`
	AuditTime  time.Time `orm:"column(AuditTime);not null;auto_now_add;type(datetime);default();"`
	Canceler   int64     `orm:"column(Canceler);not null;default();"`
	CancelTime time.Time `orm:"column(CancelTime);not null;auto_now_add;type(datetime);default();"`
}

func (this *SysMenu) TableName() string {
	return "SysMenu"
}
func (this *SysMenu) GetMenuHorizontal(id int64) []MenuNode {
	pNode := new(MenuNode)
	return this.GetMenu(id, false, pNode)
}
func (this *SysMenu) GetMenusVertical(id int64) []MenuNode {
	pNode := new(MenuNode)
	return this.GetMenu(id, true, pNode)
}
func (this *SysMenu) GetMenu(parent int64, direction bool, pNode *MenuNode) []MenuNode {
	o := orm.NewOrm()
	sysmenu := new(SysMenu)
	var menus []*SysMenu
	qt := o.QueryTable(sysmenu)
	qt.Filter("parent", parent).Filter("Direction", direction).Exclude("ID", 0).All(&menus)
	menuNodes := make([]MenuNode, 0, 0)
	for _, value := range menus {
		pNode.ID = value.ID
		pNode.URL = value.URL
		pNode.Name = value.Name
		pNode.Direction = value.Direction
		pNode.Iocn = value.Icon
		count, error := qt.Filter("parent", value.ID).Filter("Direction", direction).Exclude("ID", 0).Count()
		if error == nil && count > 0 {
			pNode.ChildrenNode = this.GetMenu(value.ID, direction, new(MenuNode))
		}
		menuNodes = append(menuNodes, *pNode)
	}
	return menuNodes
}

//GetMenuByUser 通过用户获取菜单
func GetMenuByUser(user SysOperator) ([]SysMenu, []SysPrivilege) {
	o := orm.NewOrm()

	roleUserQt := o.QueryTable(new(SysRoleOperator))
	roleMenuQt := o.QueryTable(new(SysRoleMenu))
	rolePrivilegeQT := o.QueryTable(new(SysRolePrivilege))
	menuQt := o.QueryTable(user)
	privilegeQt := o.QueryTable(new(SysPrivilege))
	var roleUserStore []SysRoleOperator
	var menuStore []SysMenu
	var roleMenuStore []SysRoleMenu
	var rolePrivilegeStore []SysRolePrivilege
	var privilegeStore []SysPrivilege
	roleUserQt.Filter("Operator", user.ID).All(&roleUserStore)

	for _, item := range roleUserStore {
		roleMenuQt.Filter("Role", item.Role).Exclude("ID", 0).All(&roleMenuStore)
		for _, roleMenu := range roleMenuStore {
			menuQt.Filter("ID", roleMenu.Menu).Exclude("ID", 0).All(&menuStore)
		}
		rolePrivilegeQT.Filter("Role", item.Role).Exclude("ID", 0).All(&rolePrivilegeStore)
		for _, rolePrivilege := range rolePrivilegeStore {
			privilegeQt.Filter("ID", rolePrivilege.Privilege).Exclude("ID", 0).All(&privilegeStore)
		}
	}
	return menuStore, privilegeStore
}
