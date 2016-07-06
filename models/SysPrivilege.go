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
	orm.RegisterModel(new(SysPrivilege))
}

type SysPrivilege struct {
	ID         int64     `orm:"column(ID);pk;unique;index;auto;"`
	Parent     int64     `orm:"column(Parent);"`
	Menu       int64     `orm:"column(Menu);"`
	Name       string    `orm:"column(Name);size(100);"`
	Clazz      string    `orm:"column(Clazz);size(100);"`
	Area       string    `orm:"column(Area);size(100);"`
	Controller string    `orm:"column(Controller);size(100);"`
	Method     string    `orm:"column(Method);size(100);"`
	Parameter  string    `orm:"column(Parameter);size(100);"`
	Url        string    `orm:"column(Url);size(100);"`
	Reserve    string    `orm:"column(Reserve);null;size(50);"`
	Remark     string    `orm:"column(Remark);null;size(250);"`
	Creator    int64     `orm:"column(Creator);null;"`
	CreateTime time.Time `orm:"column(CreateTime);null;type(datetime);"`
	Auditor    int64     `orm:"column(Auditor);null;"`
	AuditTime  time.Time `orm:"column(AuditTime);null;type(datetime);"`
	Canceler   int64     `orm:"column(Canceler);null;"`
	CancelTime time.Time `orm:"column(CancelTime);null;type(datetime);"`
}

func (this *SysPrivilege) TableName() string {
	return "SysPrivilege"
}

func (this *SysPrivilege) Add() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(this)
	return id, err
}

func (this *SysPrivilege) Count(condation *orm.Condition) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	return qs.Count()
}

func (this *SysPrivilege) Update(cols ...string) (num int64, err error) {
	o := orm.NewOrm()
	if o.Read(this) == nil {
		num, err := o.Update(this, cols...)
		return num, err
	}
	return 0, errors.New("找不到ID=‘" + string(this.ID) + "’的数据!")
}

func (this *SysPrivilege) Delete() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(this)
	return num, err
}

func (this *SysPrivilege) Read(cols ...string) (*SysPrivilege, error) {
	o := orm.NewOrm()
	err := o.Read(this, cols...)
	if err != nil {
		return this, err
	}
	return this, nil
}

func (this *SysPrivilege) GetAll(condation *orm.Condition, sort string) (models []SysPrivilege) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	qs.All(&models)
	return models
}

func (this *SysPrivilege) Getlist(condation *orm.Condition, page int64, page_size int64, sort string) (models []orm.Params, count int64) {
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

func (this *SysPrivilege) Validation() (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&this)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

// GetPrivilegesByUser
func GetPrivilegesByUser(user SysOperator) []SysPrivilege {
	o := orm.NewOrm()
	var privilegesResult []SysPrivilege
	o.Raw(` 
			select * from sysroleoperator a
			LEFT JOIN sysroleprivilege b on a.Role=b.Role
			left JOIN sysprivilege c on b.Privilege=c.ID
			where a.Operator=? 
	   `, user.ID).QueryRows(&privilegesResult)

	return privilegesResult
}
