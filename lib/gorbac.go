package lib

import (
	"blog/models"

	"strings"

	"github.com/mikespook/gorbac"
)

type UspRole struct {
	*gorbac.StdRole
	Role *models.SysRole
}
type UspMenu struct {
	gorbac.Permission
	Menu *models.SysMenu
}
type UspPrivilege struct {
	gorbac.Permission
	Privilege *models.SysPrivilege
}

func NewUspRole(usprole *models.SysRole) gorbac.Role {
	role := &UspRole{
		Role: usprole,
	}
	role.StdRole = gorbac.NewStdRole(usprole.Name)
	return role
}

func NewUspPermissionPrivilege(p *models.SysPrivilege) gorbac.Permission {
	// loading extra properties by `name`.
	Privilege := &UspPrivilege{
		Privilege: p,
	}

	Privilege.Permission = gorbac.NewStdPermission(strings.ToLower(p.Url))
	return Privilege
}

func NewUspPermissionMenu(m *models.SysMenu) gorbac.Permission {
	Menu := &UspMenu{
		Menu: m,
	}
	Menu.Permission = gorbac.NewStdPermission(strings.ToLower(m.URL))
	return Menu
}

func InitRole(role *models.SysRole, menus []models.SysMenu, privileges []models.SysPrivilege) *gorbac.RBAC {
	rbac := gorbac.New()
	r := NewUspRole(role)
	result := r.(*UspRole)
	for _, item := range menus {
		t := NewUspPermissionMenu(&item)
		result.Assign(t)
	}
	for _, item := range privileges {
		t := NewUspPermissionPrivilege(&item)
		result.Assign(t)
	}
	rbac.Add(r)
	return rbac
}

// func CheckPrivilege(rbac *gorbac.RBAC, url string) {
// 	rbac.IsGranted()
// }
