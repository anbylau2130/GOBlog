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
	orm.RegisterModel(new(OpenPlatformMsg))
}

type OpenPlatformMsg struct {
	ID           int64     `orm:"column(ID);pk;unique;index;auto;"`
	Corp         int64     `orm:"column(Corp);"`
	PlatformType string    `orm:"column(PlatformType);size(100);"`
	MsgType      string    `orm:"column(MsgType);size(50);"`
	Title        string    `orm:"column(Title);size(1000);"`
	Description  string    `orm:"column(Description);size(1000);"`
	PicUrl       string    `orm:"column(PicUrl);size(1000);"`
	Url          string    `orm:"column(Url);not null;size(100);"`
	Content      string    `orm:"column(Content);not null;"`
	Children     string    `orm:"column(Children);not null;size(1000);"`
	Sort         byte      `orm:"column(Sort);"`
	Reserve      string    `orm:"column(Reserve);not null;size(250);"`
	Remark       string    `orm:"column(Remark);not null;size(250);"`
	Creator      int64     `orm:"column(Creator);null;"`
	CreateTime   time.Time `orm:"column(CreateTime);null;type(datetime);"`
	Auditor      int64     `orm:"column(Auditor);null;"`
	AuditTime    time.Time `orm:"column(AuditTime);null;type(datetime);"`
	Canceler     int64     `orm:"column(Canceler);null;"`
	CancelTime   time.Time `orm:"column(CancelTime);null;type(datetime);"`
}

func (this *OpenPlatformMsg) TableName() string {
	return "OpenPlatformMsg"
}

func (this *OpenPlatformMsg) Add() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(this)
	return id, err
}

func (this *OpenPlatformMsg) Count(condation *orm.Condition) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	return qs.Count()
}

func (this *OpenPlatformMsg) Update(cols ...string) (num int64, err error) {
	o := orm.NewOrm()
	if o.Read(this) == nil {
		num, err := o.Update(this, cols...)
		return num, err
	}
	return 0, errors.New("找不到ID=‘" + string(this.ID) + "’的数据!")
}

func (this *OpenPlatformMsg) Delete() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(this)
	return num, err
}

func (this *OpenPlatformMsg) Read(cols ...string) (*OpenPlatformMsg, error) {
	o := orm.NewOrm()
	err := o.Read(this, cols...)
	if err != nil {
		return this, err
	}
	return this, nil
}

func (this *OpenPlatformMsg) GetAll(condation *orm.Condition, sort string) (models []OpenPlatformMsg) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	qs.All(&models)
	return models
}

func (this *OpenPlatformMsg) Getlist(condation *orm.Condition, page int64, page_size int64, sort string) (models []orm.Params, count int64) {
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

func (this *OpenPlatformMsg) Validation() (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&this)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}
