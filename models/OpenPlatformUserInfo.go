//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(OpenPlatformUserInfo))
}

type OpenPlatformUserInfo struct {
	Subscribe     byte      `orm:"column(Subscribe);default(((0)));"`
	PlatformType  string    `orm:"column(PlatformType);size(0);default();"`
	Openid        string    `orm:"column(Openid);pk;unique;size(0);default();index;"`
	Nickname      string    `orm:"column(Nickname);size(0);default();"`
	Sex           byte      `orm:"column(Sex);default();"`
	Province      string    `orm:"column(Province);size(0);default();"`
	City          string    `orm:"column(City);size(0);default();"`
	Country       string    `orm:"column(Country);size(0);default();"`
	Headimgurl    string    `orm:"column(Headimgurl);size(0);default();"`
	Privilege     string    `orm:"column(Privilege);not null;size(0);default();"`
	SubscribeTime time.Time `orm:"column(SubscribeTime);not null;auto_now_add;type(datetime);default();"`
	Reserve       string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark        string    `orm:"column(Remark);not null;size(0);default();"`
	CreateTime    time.Time `orm:"column(CreateTime);auto_now_add;type(datetime);default((getdate()));"`
}

func (this *OpenPlatformUserInfo) TableName() string {
	return "OpenPlatformUserInfo"
}
