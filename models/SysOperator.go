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
	orm.RegisterModel(new(SysOperator))
}

type SysOperator struct {
	ID              int64     `orm:"column(ID);pk;unique;index;auto;"`
	Corp            int64     `orm:"column(Corp);"`
	LoginName       string    `orm:"column(LoginName);unique;size(100);"`
	RealName        string    `orm:"column(RealName);size(100);"`
	Password        string    `orm:"column(Password);size(32);"`
	UImage          string    `orm:"column(UImage);null;size(255);"`
	Mobile          string    `orm:"column(Mobile);unique;size(32);"`
	IdCard          string    `orm:"column(IdCard);unique;size(32);"`
	Email           string    `orm:"column(Email);size(100);"`
	WechatOpenid    string    `orm:"column(WechatOpenid);size(100);"`
	AlipayOpenid    string    `orm:"column(AlipayOpenid);size(100);"`
	Weibo           string    `orm:"column(Weibo);size(100);"`
	AvailableIP     string    `orm:"column(AvailableIP);size(40);"`
	WeatherCode     string    `orm:"column(WeatherCode);not null;size(6);"`
	VirtualIntegral int64     `orm:"column(VirtualIntegral);"`
	RealIntegral    int64     `orm:"column(RealIntegral);"`
	Balance         float64   `orm:"column(Balance);digits(28);decimals(8);"`
	FrozenBalance   float64   `orm:"column(FrozenBalance);digits(28);decimals(8);"`
	IncomingBalance float64   `orm:"column(IncomingBalance);digits(28);decimals(8);"`
	Commission      float64   `orm:"column(Commission);digits(28);decimals(8);"`
	Discount        float64   `orm:"column(Discount);digits(28);decimals(8);"`
	Province        string    `orm:"column(Province);size(100);"`
	Area            string    `orm:"column(Area);size(100);"`
	County          string    `orm:"column(County);size(100);"`
	Community       string    `orm:"column(Community);size(100);"`
	Address         string    `orm:"column(Address);size(100);"`
	Status          int64     `orm:"column(Status);"`
	Skin            int64     `orm:"column(Skin);"`
	Grade           int64     `orm:"column(Grade);not null;"`
	Star            int64     `orm:"column(Star);not null;"`
	Session         string    `orm:"column(Session);not null;size(100);"`
	LoginTime       time.Time `orm:"column(LoginTime);null;auto_now_add;type(datetime);"`
	LoginIP         string    `orm:"column(LoginIP);null;size(40);"`
	LoginAgent      string    `orm:"column(LoginAgent);null;size(250);"`
	LoginCount      int64     `orm:"column(LoginCount);null;"`
	LoginErrorCount int64     `orm:"column(LoginErrorCount);null;"`
	FrozenStartTime time.Time `orm:"column(FrozenStartTime);null;auto_now_add;type(datetime);"`
	FrozenEndTime   time.Time `orm:"column(FrozenEndTime);null;auto_now_add;type(datetime);"`
	Reserve         string    `orm:"column(Reserve);null;size(50);"`
	Remark          string    `orm:"column(Remark);null;size(250);"`
	Creator         int64     `orm:"column(Creator);null;"`
	CreateTime      time.Time `orm:"column(CreateTime);null;type(datetime);"`
	Auditor         int64     `orm:"column(Auditor);null;"`
	AuditTime       time.Time `orm:"column(AuditTime);null;type(datetime);"`
	Canceler        int64     `orm:"column(Canceler);null;"`
	CancelTime      time.Time `orm:"column(CancelTime);null;type(datetime);"`
}

func (this *SysOperator) TableName() string {
	return "SysOperator"
}

func (this *SysOperator) Add(Corp, Creator int64, LoginName, RealName, Password, IdCard, Email, Mobile, Role string) (*ProcResult, error) {
	o := orm.NewOrm()
	res := new(ProcResult)
	_, err := o.Raw("call usp_addoperator(?,?,?,?,?,?,?,?,?)", Corp, LoginName, RealName, Password, IdCard, Email, Mobile, Creator, Role).RowsToStruct(res, "IsSuccess", "ProcMsg")
	return res, err
}

func (this *SysOperator) Count(condation *orm.Condition) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	return qs.Count()
}

func (this *SysOperator) Update(cols ...string) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Update(this, cols...)
	return num, err
	return 0, errors.New("找不到ID=‘" + string(this.ID) + "’的数据!")
}

func (this *SysOperator) Delete() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(this)
	return num, err
}

func (this *SysOperator) Read(cols ...string) (*SysOperator, error) {
	o := orm.NewOrm()
	err := o.Read(this, cols...)
	if err != nil {
		return this, err
	}
	return this, nil
}

func (this *SysOperator) GetAll(condation *orm.Condition, sort string) (models []SysOperator) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	qs.All(&models)
	return models
}

func (this *SysOperator) Getlist(condation *orm.Condition, page int64, page_size int64, sort string) (models []orm.Params, count int64) {
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

func (this *SysOperator) Validation() (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&this)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

func GetOperatorByLoginName(loginName string) (operator SysOperator) {
	operator = SysOperator{LoginName: loginName}
	o := orm.NewOrm()
	o.Read(&operator, "LoginName")
	return operator
}
