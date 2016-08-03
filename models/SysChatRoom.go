//该代码使用CodeSmith
//author：tsui

package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(SysChatRoom))
}

type SysChatRoom struct {
	ID         int64     `orm:"column(ID);pk;unique;index;auto;"`
	UUID       string    `orm:"column(UUID);unique;"`
	PicURL     string    `orm:"column(PicURL);size(250);"`
	Name       string    `orm:"column(Name);size(100);"`
	Title      string    `orm:"column(Title);size(250);"`
	Reserve    string    `orm:"column(Reserve);null;size(250);"`
	Remark     string    `orm:"column(Remark);null;size(250);"`
	Creator    int64     `orm:"column(Creator);null;"`
	CreateTime time.Time `orm:"column(CreateTime);null;type(datetime);"`
	Auditor    int64     `orm:"column(Auditor);null;"`
	AuditTime  time.Time `orm:"column(AuditTime);null;type(datetime);"`
	Canceler   int64     `orm:"column(Canceler);null;"`
	CancelTime time.Time `orm:"column(CancelTime);null;type(datetime);"`
}

func (this *SysChatRoom) TableName() string {
	return "SysChatRoom"
}

func (this *SysChatRoom) Add() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(this)
	return id, err
}

func (this *SysChatRoom) Count(condation *orm.Condition) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	return qs.Count()
}

func (this *SysChatRoom) Update(cols ...string) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Update(this, cols...)
	return num, err
}

func (this *SysChatRoom) Delete() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(this)
	return num, err
}

func (this *SysChatRoom) Read(cols ...string) (*SysChatRoom, error) {
	o := orm.NewOrm()
	err := o.Read(this, cols...)
	if err != nil {
		return this, err
	}
	return this, nil
}

func (this *SysChatRoom) GetAll(condation *orm.Condition, sort string) (models []SysChatRoom) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	qs.All(&models)
	return models
}

func (this *SysChatRoom) Getlist(condation *orm.Condition, page int64, page_size int64, sort string) (models []orm.Params, count int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	qs.Limit(page_size, page).OrderBy(sort).Values(&models)
	count, _ = qs.Count()
	return models, count
}

func (this *SysChatRoom) Validation() (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&this)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

//GetMenuByUser 通过用户获取菜单
func (this *SysChatRoom) GetModelsBySql(sql string, params ...interface{}) []SysChatRoom {
	o := orm.NewOrm()

	var result []SysChatRoom

	o.Raw(sql, params...).QueryRows(&result)

	return result
}
