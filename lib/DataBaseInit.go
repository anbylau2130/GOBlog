package lib

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/wendal/go-oci8"
)

var o orm.Ormer

func init() {

	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", db_type, dns)
	//orm.RegisterDataBase("default", "mysql", "root:toor@tcp(127.0.0.1:3306)/usp?charset=utf8")
	// orm.RegisterDriver("oci8", orm.DROracle)
	// dns := fmt.Sprintf("%s/%s@%s:%s/%s", db_user, db_pass, db_host, db_port, db_name)
	// orm.RegisterDataBase("default", db_type, dns)
	// o := orm.NewOrm()
	// var t test
	// err := o.Raw("SELECT * FROM emp").QueryRow(&t)
	// fmt.Print(err)
}

func Syncdb() {
	Createdb()
	Connect()
	o = orm.NewOrm()
	// 数据库别名
	name := "default"
	// drop table 后再建表
	force := true
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("database init is complete.\nPlease restart the application")
	initDbData()
}

//创建数据库
func Createdb() {

	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	db_path := beego.AppConfig.String("db_path")
	db_sslmode := beego.AppConfig.String("db_sslmode")

	var dns string
	var sqlstring string
	switch db_type {
	case "mysql":
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", db_user, db_pass, db_host, db_port)
		sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", db_name)
		break
	case "postgres":
		dns = fmt.Sprintf("host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_host, db_user, db_pass, db_port, db_sslmode)
		sqlstring = fmt.Sprintf("CREATE DATABASE %s", db_name)
		break
	case "sqlite3":
		if db_path == "" {
			db_path = "./"
		}
		dns = fmt.Sprintf("%s%s.db", db_path, db_name)
		os.Remove(dns)
		sqlstring = "create table init (n varchar(32));drop table init;"
		break
	// case "oci8":
	// 	dns = fmt.Sprintf("%s/%s@%s:%s/%s", db_user, db_pass, db_host, db_port, db_name)
	// 	sqlstring = fmt.Sprintf(" create tablespace test datafile '%s' size 500M;", db_name)
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}
	db, err := sql.Open(db_type, dns)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
	} else {
		log.Println("Database ", db_name, " created")
	}
	defer db.Close()

}

//数据库连接
func Connect() {
	var dns string
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	db_path := beego.AppConfig.String("db_path")
	db_sslmode := beego.AppConfig.String("db_sslmode")
	switch db_type {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
		break
	case "postgres":
		orm.RegisterDriver("postgres", orm.DRPostgres)
		dns = fmt.Sprintf("dbname=%s host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_name, db_host, db_user, db_pass, db_port, db_sslmode)
	case "sqlite3":
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		if db_path == "" {
			db_path = "./"
		}
		dns = fmt.Sprintf("%s%s.db", db_path, db_name)
		break
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}
	orm.RegisterDataBase("default", db_type, dns)
}

func initDbData() {
	o := orm.NewOrm()
	result, error := o.Raw(`

	`).Exec()
	fmt.Println(result, error)
}
