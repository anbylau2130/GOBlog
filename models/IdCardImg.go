//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(IdCardImg))
}

type IdCardImg struct {
	ID         int64     `orm:"column(ID);default();index;"`
	Operator   int64     `orm:"column(Operator);default();"`
	ImageUrl   string    `orm:"column(ImageUrl);size(0);default();"`
	Front      bool      `orm:"column(Front);default(((0)));"`
	Reserve    string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark     string    `orm:"column(Remark);not null;size(0);default();"`
	Creator    int64     `orm:"column(Creator);default();"`
	CreateTime time.Time `orm:"column(CreateTime);auto_now_add;type(datetime);default((getdate()));"`
	Auditor    int64     `orm:"column(Auditor);not null;default();"`
	AuditTime  time.Time `orm:"column(AuditTime);not null;auto_now_add;type(datetime);default();"`
}

func (this *IdCardImg) TableName() string {
	return "IdCardImg"
}
