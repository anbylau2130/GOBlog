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
	orm.RegisterModel(new(OpenPlatformMsgType))
}

type OpenPlatformMsgType struct {
	ID           int64     `orm:"column(ID);pk;unique;index;auto;"`
	PlatformType string    `orm:"column(PlatformType);size(100);"`
	Type         string    `orm:"column(Type);unique;size(50);"`
	Name         string    `orm:"column(Name);size(50);"`
	Reserve      string    `orm:"column(Reserve);not null;size(250);"`
	Remark       string    `orm:"column(Remark);not null;size(250);"`
	Creator      int64     `orm:"column(Creator);null;"`
	CreateTime   time.Time `orm:"column(CreateTime);null;type(datetime);"`
	Auditor      int64     `orm:"column(Auditor);null;"`
	AuditTime    time.Time `orm:"column(AuditTime);null;type(datetime);"`
	Canceler     int64     `orm:"column(Canceler);null;"`
	CancelTime   time.Time `orm:"column(CancelTime);null;type(datetime);"`
}

func (this *OpenPlatformMsgType) TableName() string {
	return "OpenPlatformMsgType"
}

func (this *OpenPlatformMsgType) Add() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(this)
	return id, err
}

func (this *OpenPlatformMsgType) Count(condation *orm.Condition) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	return qs.Count()
}

func (this *OpenPlatformMsgType) Update(cols ...string) (num int64, err error) {
	o := orm.NewOrm()
	if o.Read(this) == nil {
		num, err := o.Update(this, cols...)
		return num, err
	}
	return 0, errors.New("找不到ID=‘" + string(this.ID) + "’的数据!")
}

func (this *OpenPlatformMsgType) Delete() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(this)
	return num, err
}

func (this *OpenPlatformMsgType) Read(cols ...string) (*OpenPlatformMsgType, error) {
	o := orm.NewOrm()
	err := o.Read(this, cols...)
	if err != nil {
		return this, err
	}
	return this, nil
}

func (this *OpenPlatformMsgType) GetAll(condation *orm.Condition, sort string) (models []OpenPlatformMsgType) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	qs.All(&models)
	return models
}

func (this *OpenPlatformMsgType) Getlist(condation *orm.Condition, page int64, page_size int64, sort string) (models []orm.Params, count int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	if condation != nil {
		qs.SetCond(condation)
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&models)
	count, _ = qs.Count()
	return models, count
}

func (this *OpenPlatformMsgType) Validation() (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&this)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}
