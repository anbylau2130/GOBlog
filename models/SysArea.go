//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(SysArea))
}

type SysArea struct {
	Code        string    `orm:"column(Code);pk;unique;default();"`
	Parent      string    `orm:"column(Parent);default();"`
	Name        string    `orm:"column(Name);size(0);default();"`
	WeatherCode string    `orm:"column(WeatherCode);not null;size(0);default();"`
	LocationX   float64   `orm:"column(LocationX);not null;digits(9);decimals(18);default(((0)));"`
	LocationY   float64   `orm:"column(LocationY);not null;digits(9);decimals(18);default(((0)));"`
	Reserve     string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark      string    `orm:"column(Remark);not null;size(0);default();"`
	Creator     int64     `orm:"column(Creator);default();"`
	CreateTime  time.Time `orm:"column(CreateTime);auto_now_add;type(datetime);default((getdate()));"`
	Auditor     int64     `orm:"column(Auditor);not null;default();"`
	AuditTime   time.Time `orm:"column(AuditTime);not null;auto_now_add;type(datetime);default();"`
	Canceler    int64     `orm:"column(Canceler);not null;default();"`
	CancelTime  time.Time `orm:"column(CancelTime);not null;auto_now_add;type(datetime);default();"`
}

func (this *SysArea) TableName() string {
	return "SysArea"
}
