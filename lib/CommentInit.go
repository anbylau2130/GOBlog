package lib

import (
	"blog/models"
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
)

const ControllerPath = "blog/controllers/"

func RegisterMenus(cList ...beego.ControllerInterface) {
	for _, c := range cList {
		reflectVal := reflect.ValueOf(c)
		t := reflect.Indirect(reflectVal).Type()
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			panic("未能找到GOPATH路径,请设置GOPATH！")
		}
		pkgpath := ""

		wgopath := filepath.SplitList(gopath)
		for _, wg := range wgopath {
			wg, _ = filepath.EvalSymlinks(filepath.Join(wg, "src", t.PkgPath()))
			if utils.FileExists(wg) {
				pkgpath = wg
				break
			}
		}
		if pkgpath != "" {
			parserPkg(pkgpath, t.PkgPath())
		}
	}
}

func parserPkg(pkgRealpath, pkgpath string) (menuList []models.SysMenu, privilegeList []models.SysPrivilege, err error) {
	fileSet := token.NewFileSet()
	astPkgs, err := parser.ParseDir(fileSet, pkgRealpath, func(info os.FileInfo) bool {
		name := info.Name()
		return !info.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go")
	}, parser.ParseComments)

	if err != nil {
		return nil, nil, err
	}
	for _, pkg := range astPkgs {
		for _, fl := range pkg.Files {
			for _, d := range fl.Decls {
				switch specDecl := d.(type) {
				case *ast.FuncDecl:
					if specDecl.Recv != nil {
						exp, ok := specDecl.Recv.List[0].Type.(*ast.StarExpr) // Check that the type is correct first beforing throwing to parser
						if ok {
							menus, privileges, error := parserComments(specDecl.Doc, specDecl.Name.String(), fmt.Sprint(exp.X), pkgpath)
							if error == nil {
								for _, menu := range menus {
									menuList = append(menuList, menu)
								}
								for _, privilige := range privileges {
									privilegeList = append(privilegeList, privilige)
								}
							}
						}
					}
				}
			}
		}
	}
	return menuList, privilegeList, nil
}

func parserComments(comments *ast.CommentGroup, funcName, controllerName, pkgpath string) ([]models.SysMenu, []models.SysPrivilege, error) {
	var menuList []models.SysMenu
	var privilegeList []models.SysPrivilege
	if comments != nil && comments.List != nil {
		for _, c := range comments.List {
			t := strings.TrimSpace(strings.TrimLeft(c.Text, "//")) //@Menu name:菜单管理
			switch {
			case strings.HasPrefix(t, "@MenuH"):
				menu, error := parserMenuH(t, pkgpath, funcName, controllerName)
				if error != nil {
					return nil, nil, error
				}
				if menu != nil {
					menuList = append(menuList, *menu)
				}
			case strings.HasPrefix(t, "@MenuV"):
				menu, error := parserMenuV(t, pkgpath, funcName, controllerName)
				if error != nil {
					return nil, nil, error
				}
				if menu != nil {
					menuList = append(menuList, *menu)
				}
			case strings.HasPrefix(t, "@Privilege"):
				privilege, error := parserPrivilege(t, pkgpath, funcName, controllerName)
				if error != nil {
					return nil, nil, error
				}
				if privilege != nil {
					privilegeList = append(privilegeList, *privilege)
				}
			}
		}
	}
	return menuList, privilegeList, nil
}

//comment:@MenuH {name:"菜单管理";parent:"0"}
//commentType:"@MenuH"
func parserJson(comment string, commentType string) (commentJSON map[string]string, err error) {
	commentJSON = make(map[string]string)
	if strings.HasPrefix(comment, commentType) {
		elements := strings.TrimLeft(comment, commentType) //{name:"菜单管理";parent:"0"}
		temp := []byte(elements)
		if error := json.Unmarshal(temp, &commentJSON); error == nil {
			commentJSON["name"] = getMapValueByKey(commentJSON, "name")
			commentJSON["parent"] = getMapValueByKey(commentJSON, "parent")
			commentJSON["icon"] = getMapValueByKey(commentJSON, "icon")
		} else {
			return nil, error
		}

	}
	return commentJSON, nil
}

func getMapValueByKey(m map[string]string, key string) string {
	if value, ok := m[key]; ok {
		return value
	}
	return ""
}

func parserMenuH(t string, pkgpath string, funcName string, controllerName string) (*models.SysMenu, error) {
	menu := models.SysMenu{}
	o := orm.NewOrm()
	menuH, error := parserJson(t, "@MenuH ")
	if error != nil {
		return nil, error
	}
	menu.Name = menuH["name"]
	if o.Read(&menu, "Name"); menu.ID != 0 {
		return nil, nil
	}

	menu.Icon = menuH["icon"]
	menu.Direction = false
	menu.Area = strings.TrimLeft(pkgpath, ControllerPath)
	menu.Controller = controllerName
	menu.Method = funcName
	menu.Clazz = pkgpath + "/" + controllerName
	menu.URL = "/" + menu.Area + "/" + strings.TrimRight(menu.Controller, "Controller") + "/" + menu.Method

	if menuH["parent"] != "0" {
		menuModel := models.SysMenu{Name: menuH["parent"]}
		if error := o.Read(&menuModel, "Name"); error == nil {
			menu.Parent = menuModel.ID
		}
	} else {
		menu.Parent = 0
	}
	menu.ID, _ = o.Insert(&menu)
	return &menu, nil
}

func parserMenuV(t string, pkgpath string, funcName string, controllerName string) (*models.SysMenu, error) {
	o := orm.NewOrm()
	menu := models.SysMenu{}
	menuV, error := parserJson(t, "@MenuV ")
	if error != nil {
		return nil, error
	}

	menu.Name = menuV["name"]
	if o.Read(&menu, "Name"); menu.ID != 0 {
		return nil, nil
	}
	menu.Icon = menuV["icon"]
	menu.Direction = true
	menu.Area = strings.TrimLeft(pkgpath, ControllerPath)
	menu.Controller = controllerName
	menu.Method = funcName
	menu.Clazz = pkgpath + "/" + controllerName
	menu.URL = "/" + menu.Area + "/" + strings.TrimRight(menu.Controller, "Controller") + "/" + menu.Method
	menuModel := models.SysMenu{Name: menuV["parent"]}

	if menuModel.Name != "" {
		if error := o.Read(&menuModel, "Name"); error == nil {
			menu.Parent = menuModel.ID
		}
	} else {
		return nil, errors.New(menu.URL + "注释,请设置 @MenuV parent,")
	}
	menu.ID, _ = o.Insert(&menu)
	return &menu, nil
}

func parserPrivilege(t string, pkgpath string, funcName string, controllerName string) (*models.SysPrivilege, error) {
	o := orm.NewOrm()
	plg := models.SysPrivilege{}
	privilege, error := parserJson(t, "@Privilege ")
	if error != nil {
		return nil, error
	}
	plg.Name = privilege["name"]
	plg.Area = strings.TrimLeft(pkgpath, ControllerPath)
	plg.Controller = controllerName
	plg.Method = funcName
	plg.Clazz = pkgpath + "/" + controllerName
	plg.Url = "/" + plg.Area + "/" + strings.TrimRight(plg.Controller, "Controller") + "/" + plg.Method
	menu := models.SysMenu{Name: privilege["parent"]}
	if menu.Name != "" {
		if error := o.Read(&menu, "Name"); error == nil {
			plg.Menu = menu.ID
		}
	} else {
		return nil, errors.New(plg.Url + "该权限未设置所属页面!")
	}
	if o.Read(&plg, "Area", "Controller", "Method", "Name"); plg.ID == 0 {
		plg.ID, _ = o.Insert(&plg)
	}
	return &plg, nil
}
