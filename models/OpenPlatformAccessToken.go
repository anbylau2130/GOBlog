//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(OpenPlatformAccessToken))
}

type OpenPlatformAccessToken struct {
	ID           int64     `orm:"column(ID);pk;unique;default();index;"`
	Corp         int64     `orm:"column(Corp);default();"`
	PlatformType string    `orm:"column(PlatformType);size(0);default();"`
	AccessToken  string    `orm:"column(AccessToken);size(0);default();"`
	AccessTime   time.Time `orm:"column(AccessTime);auto_now_add;type(datetime);default();"`
	ExpiresIn    int64     `orm:"column(ExpiresIn);default();"`
	Expires      int64     `orm:"column(Expires);default();"`
	Reserve      string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark       string    `orm:"column(Remark);not null;size(0);default();"`
	CreateTime   time.Time `orm:"column(CreateTime);auto_now_add;type(datetime);default((getdate()));"`
}

func (this *OpenPlatformAccessToken) TableName() string {
	return "OpenPlatformAccessToken"
}
