//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(OpenPlatformMT))
}

type OpenPlatformMT struct {
	ID           int64     `orm:"column(ID);pk;unique;default();index;"`
	PlatformType string    `orm:"column(PlatformType);size(0);default();"`
	Mo           int64     `orm:"column(Mo);default(((0)));"`
	Msg          int64     `orm:"column(Msg);default(((0)));"`
	ToUserName   string    `orm:"column(ToUserName);size(0);default();"`
	FromUserName string    `orm:"column(FromUserName);size(0);default();"`
	MsgType      string    `orm:"column(MsgType);size(0);default();"`
	Title        string    `orm:"column(Title);size(0);default((''));"`
	Description  string    `orm:"column(Description);size(0);default((''));"`
	PicUrl       string    `orm:"column(PicUrl);size(0);default((''));"`
	Content      string    `orm:"column(Content);default();"`
	Children     string    `orm:"column(Children);size(0);default((''));"`
	Priority     byte      `orm:"column(Priority);default(((0)));"`
	SendTime     time.Time `orm:"column(SendTime);auto_now_add;type(datetime);default();"`
	ErrCode      int32     `orm:"column(ErrCode);default(((-1)));"`
	ErrMsg       string    `orm:"column(ErrMsg);size(0);default((''));"`
	Reserve      string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark       string    `orm:"column(Remark);not null;size(0);default();"`
	PerformTime  time.Time `orm:"column(PerformTime);not null;auto_now_add;type(datetime);default();"`
	PresentTime  time.Time `orm:"column(PresentTime);auto_now_add;type(datetime);default();"`
	Presenter    int64     `orm:"column(Presenter);default(((0)));"`
	AuditTime    time.Time `orm:"column(AuditTime);not null;auto_now_add;type(datetime);default();"`
	Auditor      int64     `orm:"column(Auditor);not null;default();"`
}

func (this *OpenPlatformMT) TableName() string {
	return "OpenPlatformMT"
}
