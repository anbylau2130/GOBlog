//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(SysPrivilege))
}

type SysPrivilege struct {
	ID         int64     `orm:"column(ID);pk;unique;default();index;"`
	Parent     int64     `orm:"column(Parent);default();"`
	Menu       int64     `orm:"column(Menu);default();"`
	Name       string    `orm:"column(Name);size(0);default((''));"`
	Clazz      string    `orm:"column(Clazz);size(0);default((''));"`
	Area       string    `orm:"column(Area);size(0);default((''));"`
	Controller string    `orm:"column(Controller);size(0);default((''));"`
	Method     string    `orm:"column(Method);size(0);default((''));"`
	Parameter  string    `orm:"column(Parameter);size(0);default((''));"`
	Url        string    `orm:"column(Url);size(0);default((''));"`
	Reserve    string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark     string    `orm:"column(Remark);not null;size(0);default();"`
	Creator    int64     `orm:"column(Creator);default();"`
	CreateTime time.Time `orm:"column(CreateTime);auto_now_add;type(datetime);default((getdate()));"`
	Auditor    int64     `orm:"column(Auditor);not null;default();"`
	AuditTime  time.Time `orm:"column(AuditTime);not null;auto_now_add;type(datetime);default();"`
	Canceler   int64     `orm:"column(Canceler);not null;default();"`
	CancelTime time.Time `orm:"column(CancelTime);not null;auto_now_add;type(datetime);default();"`
}

func (this *SysPrivilege) TableName() string {
	return "SysPrivilege"
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
