//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(OpenPlatformMsg))
}

type OpenPlatformMsg struct {
	ID           int64     `orm:"column(ID);pk;unique;default();index;"`
	Corp         int64     `orm:"column(Corp);default();"`
	PlatformType string    `orm:"column(PlatformType);size(0);default();"`
	MsgType      string    `orm:"column(MsgType);size(0);default((''));"`
	Title        string    `orm:"column(Title);size(0);default((''));"`
	Description  string    `orm:"column(Description);size(0);default((''));"`
	PicUrl       string    `orm:"column(PicUrl);size(0);default((''));"`
	Url          string    `orm:"column(Url);not null;size(0);default();"`
	Content      string    `orm:"column(Content);not null;default();"`
	Children     string    `orm:"column(Children);not null;size(0);default();"`
	Sort         byte      `orm:"column(Sort);default(((0)));"`
	Reserve      string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark       string    `orm:"column(Remark);not null;size(0);default();"`
	Creator      int64     `orm:"column(Creator);default();"`
	CreateTime   time.Time `orm:"column(CreateTime);auto_now_add;type(datetime);default();"`
	Auditor      int64     `orm:"column(Auditor);not null;default();"`
	AuditTime    time.Time `orm:"column(AuditTime);not null;auto_now_add;type(datetime);default();"`
	Canceler     int64     `orm:"column(Canceler);not null;default();"`
	CancelTime   time.Time `orm:"column(CancelTime);not null;auto_now_add;type(datetime);default();"`
}

func (this *OpenPlatformMsg) TableName() string {
	return "OpenPlatformMsg"
}
