package routers

import (
	"blog/controllers"
	"blog/controllers/admin"
	"blog/models"
	"encoding/json"
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

func init() {

	RegisterMenus(&admin.MenusController{})
	beego.Router("/", &controllers.MainController{}, "*:Home")
	beego.Router("/Login", &controllers.MainController{}, "*:Login")
	beego.Router("/main", &controllers.MainController{}, "*:Home")
	beego.Router("/main/Home", &controllers.MainController{}, "*:Home")
	beego.Router("/main/Login", &controllers.MainController{}, "*:Login")
	beego.Router("/main/GetMenuHorizontal", &controllers.MainController{}, "*:GetMenuHorizontal")
	beego.Router("/main/GetMenusVertical", &controllers.MainController{}, "*:GetMenusVertical")
	beego.AutoPrefix("/Admin", &admin.MenusController{})
	//beego.Include(&controllers.MainController{})
	//beego.AutoPrefix("/Admin",&controllers.AdminController{})

}

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

func parserPkg(pkgRealpath, pkgpath string) error {
	rep := strings.NewReplacer("/", "_", ".", "_")
	commentFilename := rep.Replace(pkgpath) + ".go"
	genInfoList := make(map[string][]beego.ControllerComments)
	fmt.Println(commentFilename, genInfoList)
	fileSet := token.NewFileSet()
	astPkgs, err := parser.ParseDir(fileSet, pkgRealpath, func(info os.FileInfo) bool {
		name := info.Name()
		return !info.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go")
	}, parser.ParseComments)

	if err != nil {
		return err
	}
	for _, pkg := range astPkgs {
		for _, fl := range pkg.Files {
			for _, d := range fl.Decls {
				switch specDecl := d.(type) {
				case *ast.FuncDecl:
					if specDecl.Recv != nil {
						exp, ok := specDecl.Recv.List[0].Type.(*ast.StarExpr) // Check that the type is correct first beforing throwing to parser
						if ok {
							parserComments(specDecl.Doc, specDecl.Name.String(), fmt.Sprint(exp.X), pkgpath)
						}
					}
				}
			}
		}
	}
	return nil
}

func parserComments(comments *ast.CommentGroup, funcName, controllerName, pkgpath string) error {
	if comments != nil && comments.List != nil {
		for _, c := range comments.List {

			t := strings.TrimSpace(strings.TrimLeft(c.Text, "//")) //@Menu name:菜单管理
			if menuH, error := parserPower(t, "@MenuH "); error == nil {
				menu := new(models.SysMenu)

				menu.Name = menuH["name"]
				menu.Icon = menuH["icon"]
				o := orm.NewOrm()
				pmenu := models.SysMenu{Name: menuH["parent"]}
				if menuH["parent"] != "0" {
					if error := o.Read(&pmenu, "Name"); error == nil {
						menu.Parent = pmenu.ID
					}
				}
				o.Insert(&menu)
			}
			if menuV, error := parserPower(t, "@MenuV "); error == nil {
				menu := new(models.SysMenu)
				menu.Name = menuV["name"]
				menu.Icon = menuV["icon"]
				pmenu := new(models.SysMenu)
				o := orm.NewOrm()
				qt := o.QueryTable(&pmenu)

				if menuV["parent"] != "" {
					if error := qt.Filter("Name", menuV["parent"]).One(&pmenu, "ID"); error == nil {
						menu.Parent = pmenu.ID
					}
				}
				o.Insert(&menu)
			}
			if privilege, error := parserPower(t, "@Privilege "); error == nil {
				plg := new(models.SysPrivilege)
				plg.Name = privilege["name"]
				menu := new(models.SysMenu)
				o := orm.NewOrm()
				qt := o.QueryTable(&menu)
				if privilege["parent"] != "" {
					if error := qt.Filter("Name", privilege["parent"]).One(&menu, "ID"); error == nil {
						plg.Menu = menu.ID
					}
				}
				o.Insert(&plg)
			}
			// 	//路由地址 blog/controllers/admin/MenusController
			// 	//key := pkgpath + "/" + controllerName
		}
	}
	return nil
}

//comment:@MenuH {name:"菜单管理";parent:"0"}
//commentType:"@MenuH"
func parserPower(comment string, commentType string) (commentJSON map[string]string, err error) {
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
