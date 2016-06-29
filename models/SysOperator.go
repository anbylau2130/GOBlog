//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(SysOperator))
}

type SysOperator struct {
	ID              int64     `orm:"column(ID);pk;unique;default();index;"`
	Corp            int64     `orm:"column(Corp);default();"`
	LoginName       string    `orm:"column(LoginName);unique;size(0);default();"`
	RealName        string    `orm:"column(RealName);size(0);default();"`
	Password        string    `orm:"column(Password);size(0);default();"`
	Mobile          string    `orm:"column(Mobile);unique;size(0);default((''));"`
	IdCard          string    `orm:"column(IdCard);unique;size(0);default();index;"`
	Email           string    `orm:"column(Email);size(0);default((''));"`
	WechatOpenid    string    `orm:"column(WechatOpenid);size(0);default((''));index;"`
	AlipayOpenid    string    `orm:"column(AlipayOpenid);size(0);default((''));index;"`
	Weibo           string    `orm:"column(Weibo);size(0);default((''));"`
	AvailableIP     string    `orm:"column(AvailableIP);size(0);default(('0.0.0.0'));"`
	WeatherCode     string    `orm:"column(WeatherCode);not null;size(0);default();"`
	VirtualIntegral int64     `orm:"column(VirtualIntegral);default(((0)));"`
	RealIntegral    int64     `orm:"column(RealIntegral);default(((0)));"`
	Balance         float64   `orm:"column(Balance);digits(13);decimals(20);default(((1)));"`
	FrozenBalance   float64   `orm:"column(FrozenBalance);digits(13);decimals(20);default(((1)));"`
	IncomingBalance float64   `orm:"column(IncomingBalance);digits(13);decimals(20);default(((1)));"`
	Commission      float64   `orm:"column(Commission);digits(13);decimals(20);default(((0)));"`
	Discount        float64   `orm:"column(Discount);digits(13);decimals(20);default(((1)));"`
	Province        string    `orm:"column(Province);size(0);default();"`
	Area            string    `orm:"column(Area);size(0);default();"`
	County          string    `orm:"column(County);size(0);default();"`
	Community       string    `orm:"column(Community);size(0);default();"`
	Address         string    `orm:"column(Address);size(0);default((''));"`
	Status          int64     `orm:"column(Status);default(((0)));"`
	Skin            int64     `orm:"column(Skin);default(((0)));"`
	Grade           int64     `orm:"column(Grade);not null;default(((0)));"`
	Star            int64     `orm:"column(Star);not null;default(((0)));"`
	Session         string    `orm:"column(Session);not null;size(0);default();"`
	LoginTime       time.Time `orm:"column(LoginTime);not null;auto_now_add;type(datetime);default();"`
	LoginIP         string    `orm:"column(LoginIP);not null;size(0);default();"`
	LoginAgent      string    `orm:"column(LoginAgent);not null;size(0);default();"`
	LoginCount      int64     `orm:"column(LoginCount);not null;default(((0)));"`
	LoginErrorCount int64     `orm:"column(LoginErrorCount);not null;default(((0)));"`
	FrozenStartTime time.Time `orm:"column(FrozenStartTime);not null;auto_now_add;type(datetime);default();"`
	FrozenEndTime   time.Time `orm:"column(FrozenEndTime);not null;auto_now_add;type(datetime);default();"`
	Reserve         string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark          string    `orm:"column(Remark);not null;size(0);default();"`
	Creator         int64     `orm:"column(Creator);default();"`
	CreateTime      time.Time `orm:"column(CreateTime);auto_now_add;type(datetime);default((getdate()));"`
	Auditor         int64     `orm:"column(Auditor);not null;default();"`
	AuditTime       time.Time `orm:"column(AuditTime);not null;auto_now_add;type(datetime);default();"`
	Canceler        int64     `orm:"column(Canceler);not null;default();"`
	CancelTime      time.Time `orm:"column(CancelTime);not null;auto_now_add;type(datetime);default();"`
}

func (this *SysOperator) TableName() string {
	return "SysOperator"
}
func GetOperatorByLoginName(loginName string) (operator SysOperator) {
	operator = SysOperator{LoginName: loginName}
	o := orm.NewOrm()
	o.Read(&operator, "LoginName")
	return operator
}
