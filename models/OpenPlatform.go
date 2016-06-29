//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(OpenPlatform))
}

type OpenPlatform struct {
	ID           int64     `orm:"column(ID);default();index;"`
	Corp         int64     `orm:"column(Corp);default();"`
	PlatformType string    `orm:"column(PlatformType);size(0);default();"`
	Token        string    `orm:"column(Token);size(0);default((''));"`
	Appid        string    `orm:"column(Appid);size(0);default((''));index;"`
	AppSecret    string    `orm:"column(AppSecret);size(0);default((''));"`
	Openid       string    `orm:"column(Openid);size(0);default((''));index;"`
	AesKey       string    `orm:"column(AesKey);size(0);default((''));"`
	PrivateKey   string    `orm:"column(PrivateKey);size(0);default((''));"`
	PublicKey    string    `orm:"column(PublicKey);size(0);default((''));"`
	Reserve      string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark       string    `orm:"column(Remark);not null;size(0);default();"`
	Creator      int64     `orm:"column(Creator);default();"`
	CreateTime   time.Time `orm:"column(CreateTime);auto_now_add;type(datetime);default((getdate()));"`
	Auditor      int64     `orm:"column(Auditor);not null;default();"`
	AuditTime    time.Time `orm:"column(AuditTime);not null;auto_now_add;type(datetime);default();"`
	Canceler     int64     `orm:"column(Canceler);not null;default();"`
	CancelTime   time.Time `orm:"column(CancelTime);not null;auto_now_add;type(datetime);default();"`
}

func (this *OpenPlatform) TableName() string {
	return "OpenPlatform"
}
