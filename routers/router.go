package routers

import (
	"blog/controllers"
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
	RegisterMenus(&controllers.MainController{})
	beego.Router("/", &controllers.MainController{}, "*:Home")
	beego.Router("/Login", &controllers.MainController{}, "*:Login")
	beego.Router("/main", &controllers.MainController{}, "*:Home")
	beego.Router("/main/Home", &controllers.MainController{}, "*:Home")
	beego.Router("/main/Login", &controllers.MainController{}, "*:Login")
	beego.Router("/main/GetMenuHorizontal", &controllers.MainController{}, "*:GetMenuHorizontal")
	beego.Router("/main/GetMenusVertical", &controllers.MainController{}, "*:GetMenusVertical")
	beego.Include(&controllers.MainController{})
	//beego.AutoPrefix("/Admin",&controllers.AdminController{})
	//beego.AutoRouter(&controllers.AdminController{})
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
			t := strings.TrimSpace(strings.TrimLeft(c.Text, "//"))
			if strings.HasPrefix(t, "@router") {
				elements := strings.TrimLeft(t, "@router ")
				e1 := strings.SplitN(elements, " ", 2)
				if len(e1) < 1 {
					return errors.New("you should has router infomation")
				}
				key := pkgpath + ":" + controllerName
				cc := beego.ControllerComments{}
				cc.Method = funcName
				cc.Router = e1[0]
				if len(e1) == 2 && e1[1] != "" {
					e1 = strings.SplitN(e1[1], " ", 2)
					if len(e1) >= 1 {
						cc.AllowHTTPMethods = strings.Split(strings.Trim(e1[0], "[]"), ",")
					} else {
						cc.AllowHTTPMethods = append(cc.AllowHTTPMethods, "get")
					}
				} else {
					cc.AllowHTTPMethods = append(cc.AllowHTTPMethods, "get")
				}
				if len(e1) == 2 && e1[1] != "" {
					keyval := strings.Split(strings.Trim(e1[1], "[]"), " ")
					for _, kv := range keyval {
						kk := strings.Split(kv, ":")
						cc.Params = append(cc.Params, map[string]string{strings.Join(kk[:len(kk)-1], ":"): kk[len(kk)-1]})
					}
				}
				genInfoList[key] = append(genInfoList[key], cc)
			}
		}
	}
	return nil
}
