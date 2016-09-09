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
	orm.RegisterModel(new(SysCorp))
}

type ProcResult struct {
	IsSuccess string `orm:"column(IsSuccess)"`
	ProcMsg   string `orm:"column(ProcMsg)"`
}
type SysCorp struct {
	ID                int64     `orm:"column(ID);pk;unique;index;auto;"`
	Parent            int64     `orm:"column(Parent);"`
	Priority          byte      `orm:"column(Priority);"`
	Name              string    `orm:"column(Name);size(50);"`
	BriefName         string    `orm:"column(BriefName);size(50);"`
	Certificate       string    `orm:"column(Certificate);size(50);"`
	CertificateNumber string    `orm:"column(CertificateNumber);size(50);"`
	Ceo               string    `orm:"column(Ceo);size(100);"`
	Postcode          string    `orm:"column(Postcode);size(6);"`
	Faxcode           string    `orm:"column(Faxcode);size(32);"`
	Linkman           string    `orm:"column(Linkman);size(50);"`
	Mobile            string    `orm:"column(Mobile);size(32);"`
	Email             string    `orm:"column(Email);size(10);"`
	Qq                string    `orm:"column(Qq);size(50);"`
	Wechat            string    `orm:"column(Wechat);size(100);"`
	Weibo             string    `orm:"column(Weibo);size(100);"`
	VirtualIntegral   int64     `orm:"column(VirtualIntegral);"`
	RealIntegral      int64     `orm:"column(RealIntegral);"`
	FeeType           int64     `orm:"column(FeeType);"`
	PrepayValve       float64   `orm:"column(PrepayValve);digits(28);decimals(8);"`
	Balance           float64   `orm:"column(Balance);digits(28);decimals(8);"`
	FrozenBalance     float64   `orm:"column(FrozenBalance);digits(28);decimals(8);"`
	IncomingBalance   float64   `orm:"column(IncomingBalance);digits(28);decimals(8);"`
	Commission        float64   `orm:"column(Commission);digits(28);decimals(8);"`
	Discount          float64   `orm:"column(Discount);digits(28);decimals(8);"`
	Province          string    `orm:"column(Province);size(100);"`
	Area              string    `orm:"column(Area);size(100);"`
	County            string    `orm:"column(County);size(100);"`
	Community         string    `orm:"column(Community);size(100);"`
	Address           string    `orm:"column(Address);size(100);"`
	Status            int64     `orm:"column(Status);"`
	Type              int64     `orm:"column(Type);"`
	Grade             int64     `orm:"column(Grade);"`
	Vocation          int64     `orm:"column(Vocation);"`
	Reserve           string    `orm:"column(Reserve);null;size(250);"`
	Remark            string    `orm:"column(Remark);null;size(250);"`
	LogoIcon          string    `orm:"column(LogoIcon);size(50);"`
	LogoUrl           string    `orm:"column(LogoUrl);size(50);"`
	Creator           int64     `orm:"column(Creator);null;"`
	CreateTime        time.Time `orm:"column(CreateTime);null;type(datetime);"`
	Auditor           int64     `orm:"column(Auditor);null;"`
	AuditTime         time.Time `orm:"column(AuditTime);null;type(datetime);"`
	Canceler          int64     `orm:"column(Canceler);null;"`
	CancelTime        time.Time `orm:"column(CancelTime);null;type(datetime);"`
}

func (this *SysCorp) TableName() string {
	return "SysCorp"
}

func (this *SysCorp) Add(operator, corpType, parentCorp int64, loginName, password, corpName string) (ProcResult, error) {
	o := orm.NewOrm()
	var res ProcResult
	err := o.Raw("call usp_addcorp(?,?,?,?,?,?);", corpName, corpType, operator, parentCorp, loginName, password).QueryRow(&res)
	return res, err
}

func (this *SysCorp) Count(condation *orm.Condition) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		return qs.SetCond(condation).Count()
	}
	return qs.Count()
}

func (this *SysCorp) Update(cols ...string) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Update(this, cols...)
	return num, err
}

func (this *SysCorp) Delete() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(this)
	return num, err
}

func (this *SysCorp) Read(cols ...string) (*SysCorp, error) {
	o := orm.NewOrm()
	err := o.Read(this, cols...)
	if err != nil {
		return this, err
	}
	return this, nil
}

func (this *SysCorp) GetAll(condation *orm.Condition, sort string) (models []SysCorp) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation).All(&models)
	} else {
		qs.All(&models)
	}
	return models
}

func (this *SysCorp) Getlist(condation *orm.Condition, page int64, page_size int64, sort string) (models []orm.Params, count int64) {
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

func (this *SysCorp) Validation() (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&this)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

//GetCorpByUser
func GetCorpByUser(User SysOperator) (corp SysCorp) {
	corp = SysCorp{ID: User.Corp}
	o := orm.NewOrm()
	o.Read(&corp, "ID")
	return corp
}
