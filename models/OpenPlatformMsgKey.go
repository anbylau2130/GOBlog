//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(OpenPlatformMsgKey))
}

type OpenPlatformMsgKey struct {
	ID         int64     `orm:"column(ID);pk;unique;default();index;"`
	Corp       int64     `orm:"column(Corp);default();"`
	Msg        int64     `orm:"column(Msg);default();"`
	Command    string    `orm:"column(Command);size(0);default((''));"`
	Matching   byte      `orm:"column(Matching);default(((0)));"`
	Reserve    string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark     string    `orm:"column(Remark);not null;size(0);default();"`
	Creator    int64     `orm:"column(Creator);default();"`
	CreateTime time.Time `orm:"column(CreateTime);auto_now_add;type(datetime);default();"`
	Auditor    int64     `orm:"column(Auditor);not null;default();"`
	AuditTime  time.Time `orm:"column(AuditTime);not null;auto_now_add;type(datetime);default();"`
}

func (this *OpenPlatformMsgKey) TableName() string {
	return "OpenPlatformMsgKey"
}
