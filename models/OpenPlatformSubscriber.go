//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(OpenPlatformSubscriber))
}

type OpenPlatformSubscriber struct {
	Corp         int64     `orm:"column(Corp);default();"`
	Openid       string    `orm:"column(Openid);pk;unique;size(0);default();index;"`
	PlatformType string    `orm:"column(PlatformType);size(0);default();"`
	Nickname     string    `orm:"column(Nickname);size(0);default((''));"`
	Headimgurl   string    `orm:"column(Headimgurl);size(0);default((''));"`
	Name         string    `orm:"column(Name);size(0);default((''));"`
	Alias        string    `orm:"column(Alias);size(0);default((''));"`
	Gender       string    `orm:"column(Gender);size(0);default(('男'));"`
	TelePhone    string    `orm:"column(TelePhone);size(0);default((''));"`
	MobilePhone  string    `orm:"column(MobilePhone);size(0);default((''));"`
	Integrate    int32     `orm:"column(Integrate);default(((0)));"`
	Reserve      string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark       string    `orm:"column(Remark);not null;size(0);default();"`
	EnrollTime   time.Time `orm:"column(EnrollTime);auto_now_add;type(datetime);default();"`
	CancelTime   time.Time `orm:"column(CancelTime);not null;auto_now_add;type(datetime);default();"`
}

func (this *OpenPlatformSubscriber) TableName() string {
	return "OpenPlatformSubscriber"
}
