package lib

import (
	"mime"
	"os"

	"github.com/beego/admin/src/models"
)

func Initialize() {
	mime.AddExtensionType(".css", "text/css")
	//判断初始化参数
	initArgs()

	models.Connect()

}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			Syncdb()
			os.Exit(0)
		}
	}
}
