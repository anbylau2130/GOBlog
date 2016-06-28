package controllers

import "github.com/astaxie/beego"
import "reflect"

// StructInfo struct to save structInfo
type StructInfo struct {
	StructType reflect.Type
	PkgName    string
}

// UspController 定义
type UspController struct {
	beego.Controller
}

// GetAllStruct 注册struct
func GetAllStruct() []StructInfo {
	return []StructInfo{
		{
			StructType: reflect.TypeOf(new(UspController)),
			PkgName:    "Controller",
		},
		{
			StructType: reflect.TypeOf(new(MainController)),
			PkgName:    "Controller",
		},
	}
}
