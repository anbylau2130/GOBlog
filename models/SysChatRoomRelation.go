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
	orm.RegisterModel(new(SysChatRoomRelation))
}

type SysChatRoomRelation struct {
	ID         int64     `orm:"column(ID);pk;unique;index;auto;"`
	ChatRoom   int64     `orm:"column(ChatRoom);unique;"`
	Operator   int64     `orm:"column(Operator);"`
	Reserve    string    `orm:"column(Reserve);null;size(250);"`
	Remark     string    `orm:"column(Remark);null;size(250);"`
	Creator    int64     `orm:"column(Creator);null;"`
	CreateTime time.Time `orm:"column(CreateTime);null;type(datetime);"`
	Auditor    int64     `orm:"column(Auditor);null;"`
	AuditTime  time.Time `orm:"column(AuditTime);null;type(datetime);"`
	Canceler   int64     `orm:"column(Canceler);null;"`
	CancelTime time.Time `orm:"column(CancelTime);null;type(datetime);"`
}

func (this *SysChatRoomRelation) TableName() string {
	return "SysChatRoomRelation"
}

func (this *SysChatRoomRelation) Add() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(this)
	return id, err
}

func (this *SysChatRoomRelation) Count(condation *orm.Condition) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	return qs.Count()
}

func (this *SysChatRoomRelation) Update(cols ...string) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Update(this, cols...)
	return num, err
}

func (this *SysChatRoomRelation) Delete() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(this)
	return num, err
}

func (this *SysChatRoomRelation) Read(cols ...string) (*SysChatRoomRelation, error) {
	o := orm.NewOrm()
	err := o.Read(this, cols...)
	if err != nil {
		return this, err
	}
	return this, nil
}

func (this *SysChatRoomRelation) GetAll(condation *orm.Condition, sort string) (models []SysChatRoomRelation) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	qs.All(&models)
	return models
}

func (this *SysChatRoomRelation) Getlist(condation *orm.Condition, page int64, page_size int64, sort string) (models []orm.Params, count int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	qs.Limit(page_size, page).OrderBy(sort).Values(&models)
	count, _ = qs.Count()
	return models, count
}

func (this *SysChatRoomRelation) Validation() (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&this)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}
