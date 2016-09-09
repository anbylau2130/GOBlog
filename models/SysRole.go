//该代码使用CodeSmith
//author：tsui

package models

import (
	"bytes"
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(SysRole))
}

type Select struct {
	Value int64  `json:"id"`
	Text  string `json:"text"`
}
type SysRole struct {
	ID         int64     `orm:"column(ID);pk;unique;index;auto;"`
	Corp       int64     `orm:"column(Corp);"`
	Name       string    `orm:"column(Name);size(100);"`
	Type       bool      `orm:"column(Type);"`
	Reserve    string    `orm:"column(Reserve);null;size(250);"`
	Remark     string    `orm:"column(Remark);null;size(250);"`
	Creator    int64     `orm:"column(Creator);null;"`
	CreateTime time.Time `orm:"column(CreateTime);null;type(datetime);"`
	Auditor    int64     `orm:"column(Auditor);null;"`
	AuditTime  time.Time `orm:"column(AuditTime);null;type(datetime);"`
	Canceler   int64     `orm:"column(Canceler);null;"`
	CancelTime time.Time `orm:"column(CancelTime);null;type(datetime);"`
}

func (this *SysRole) TableName() string {
	return "SysRole"
}

func (this *SysRole) Add(Corp, Creator, Type int64, Name, Remark, Menus, Privileges string) (ProcResult, error) {
	o := orm.NewOrm()
	var res ProcResult
	var buffer bytes.Buffer
	buffer.WriteString(`call usp_addrole( `)
	buffer.WriteString(`'` + strconv.FormatInt(Corp, 10) + `',`)
	buffer.WriteString(`'` + Name + `',`)
	buffer.WriteString(`'` + strconv.FormatInt(Type, 10) + `',`)
	buffer.WriteString(`'` + Remark + `',`)
	buffer.WriteString(`'` + strconv.FormatInt(Creator, 10) + `',`)
	buffer.WriteString(`'` + Menus + `',`)
	buffer.WriteString(`'` + Privileges + `');`)
	err := o.Raw(buffer.String()).QueryRow(&res)
	return res, err
}

func (this *SysRole) Count(condation *orm.Condition) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	return qs.Count()
}

func (this *SysRole) Update(role, creator int64, name, remark, menus, privileges string) (res ProcResult, err error) {
	o := orm.NewOrm()
	var buffer bytes.Buffer
	buffer.WriteString(`call usp_editrole( `)
	buffer.WriteString(`'` + strconv.FormatInt(role, 10) + `',`)
	buffer.WriteString(`'` + name + `',`)
	buffer.WriteString(`'` + remark + `',`)
	buffer.WriteString(`'` + strconv.FormatInt(creator, 10) + `',`)
	buffer.WriteString(`'` + menus + `',`)
	buffer.WriteString(`'` + privileges + `');`)
	err = o.Raw(buffer.String()).QueryRow(&res)
	return res, err
}

func (this *SysRole) Delete() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(this)
	return num, err
}

func (this *SysRole) Read(cols ...string) (*SysRole, error) {
	o := orm.NewOrm()
	err := o.Read(this, cols...)
	if err != nil {
		return this, err
	}
	return this, nil
}

func (this *SysRole) GetAll(condation *orm.Condition, sort string) (models []SysRole) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation).All(&models)
	} else {
		qs.All(&models)
	}

	return models
}

func (this *SysRole) Getlist(condation *orm.Condition, page int64, page_size int64, sort string) (models []orm.Params, count int64) {
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

func (this *SysRole) Validation() (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&this)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

func (this *SysRole) GetSelect() []Select {
	var result []Select
	model := SysRole{}
	models := model.GetAll(nil, "ID")
	for _, item := range models {
		result = append(result, *&Select{Text: item.Name, Value: item.ID})
	}
	return result
}

func GetRolesByOperator(operator int64) []SysRole {
	o := orm.NewOrm()
	var models []SysRole
	o.Raw(`
   		select b.* from SysRoleOperator a 
		left join SysRole b on a.Role=b.ID
		   where a.operator=? 
   `, operator).QueryRows(&models)
	return models
}
