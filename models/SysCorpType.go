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
	orm.RegisterModel(new(SysCorpType))
}

type SysCorpType struct {
	ID         int64     `orm:"column(ID);pk;unique;index;auto;"`
	Name       string    `orm:"column(Name);unique;size(100);"`
	Reserve    string    `orm:"column(Reserve);not null;size(250);"`
	Remark     string    `orm:"column(Remark);not null;size(250);"`
	Creator    int64     `orm:"column(Creator);null;"`
	CreateTime time.Time `orm:"column(CreateTime);null;type(datetime);"`
	Auditor    int64     `orm:"column(Auditor);null;"`
	AuditTime  time.Time `orm:"column(AuditTime);null;type(datetime);"`
	Canceler   int64     `orm:"column(Canceler);null;"`
	CancelTime time.Time `orm:"column(CancelTime);null;type(datetime);"`
}

func (this *SysCorpType) TableName() string {
	return "SysCorpType"
}

func (this *SysCorpType) Add() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(this)
	return id, err
}

func (this *SysCorpType) Count(condation *orm.Condition) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		return qs.SetCond(condation).Count()
	}
	return qs.Count()
}

func (this *SysCorpType) Update(cols ...string) (num int64, err error) {
	o := orm.NewOrm()
	if o.Read(this) == nil {
		num, err := o.Update(this, cols...)
		return num, err
	}
	return 0, errors.New("找不到ID=‘" + string(this.ID) + "’的数据!")
}

func (this *SysCorpType) Delete() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(this)
	return num, err
}

func (this *SysCorpType) Read(cols ...string) (*SysCorpType, error) {
	o := orm.NewOrm()
	err := o.Read(this, cols...)
	if err != nil {
		return this, err
	}
	return this, nil
}

func (this *SysCorpType) GetAll(condation *orm.Condition, sort string) (models []SysCorpType) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation).All(&models)
	} else {
		qs.All(&models)
	}
	return models
}

func (this *SysCorpType) Getlist(condation *orm.Condition, page int64, page_size int64, sort string) (models []orm.Params, count int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation).Limit(page_size, page).OrderBy(sort).Values(&models)
	} else {
		qs.Limit(page_size, page).OrderBy(sort).Values(&models)
	}
	count, _ = qs.Count()
	return models, count
}

func (this *SysCorpType) Validation() (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&this)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

func (this *SysCorpType) GetSelect() []Select {
	var result []Select
	model := SysCorpType{}
	models := model.GetAll(nil, "ID")
	for _, item := range models {
		result = append(result, *&Select{Text: item.Name, Value: item.ID})
	}
	return result
}
