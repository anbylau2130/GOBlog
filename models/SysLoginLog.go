//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(SysLoginLog))
}

type SysLoginLog struct {
	ID         int64     `orm:"column(ID);pk;unique;default();index;"`
	Operator   int64     `orm:"column(Operator);default();"`
	Ip         string    `orm:"column(Ip);size(0);default((''));"`
	LoginAgent string    `orm:"column(LoginAgent);not null;size(0);default();"`
	Success    bool      `orm:"column(Success);default(((0)));"`
	Time       time.Time `orm:"column(Time);auto_now_add;type(datetime);default();"`
	Reserve    string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark     string    `orm:"column(Remark);not null;size(0);default();"`
}

func (this *SysLoginLog) TableName() string {
	return "SysLoginLog"
}
