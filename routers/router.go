package routers

import (
	"blog/controllers"
	"blog/controllers/admin"
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
	var genInfoList map[string][]beego.ControllerComments
	if comments != nil && comments.List != nil {
		for _, c := range comments.List {

			t := strings.TrimSpace(strings.TrimLeft(c.Text, "//")) //@Menu name:菜单管理
			if strings.HasPrefix(t, "@Menu") {
				elements := strings.TrimLeft(t, "@MenuH ") //name:菜单管理;parent:0
				e1 := strings.SplitN(elements, ";", 2)
				if len(e1) < 1 {
					return errors.New("你需要设置菜单信息")
				}
				var menuComment map[string]string
				for _, item := range e1 {
					if strings.HasPrefix(item, "name") {
						menuComment["name"] = strings.TrimLeft(item, "name:")
					}
					if strings.HasPrefix(item, "parent") {
						menuComment["parent"] = strings.TrimLeft(item, "parent:")
					}
					if strings.HasPrefix(item, "icon") {
						menuComment["icon"] = strings.TrimLeft(item, "icon:")
					}
				}
				key := pkgpath + "/" + controllerName
				// o := orm.NewOrm()
				// menunew(models.SysMenu)
				// o.Raw("select *　from SysMenu where Name ='?' ", menuComment["parent"]).QueryRow(&)

				// menu := models.SysMenu{
				// 	Name: menuComment["name"],
				// 	Icon: menuComment["icon"],
				// 	URL:  strings.TrimRight(strings.TrimLeft(key, controllers.ControllerPath), "controllers"),
				// }

				//路由地址 blog/controllers/admin/MenusController

				cc := beego.ControllerComments{}
				genInfoList[key] = append(genInfoList[key], cc)
			}
		}
	}
	return nil
}
