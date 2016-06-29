//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(SysCorp))
}

type SysCorp struct {
	ID                int64     `orm:"column(ID);pk;unique;default();index;"`
	Parent            int64     `orm:"column(Parent);default();"`
	Priority          byte      `orm:"column(Priority);default();"`
	Name              string    `orm:"column(Name);size(0);default((''));"`
	BriefName         string    `orm:"column(BriefName);size(0);default((''));"`
	Certificate       string    `orm:"column(Certificate);size(0);default((''));"`
	CertificateNumber string    `orm:"column(CertificateNumber);size(0);default((''));"`
	Ceo               string    `orm:"column(Ceo);size(0);default((''));"`
	Postcode          string    `orm:"column(Postcode);size(0);default((''));"`
	Faxcode           string    `orm:"column(Faxcode);size(0);default((''));"`
	Linkman           string    `orm:"column(Linkman);size(0);default((''));"`
	Mobile            string    `orm:"column(Mobile);size(0);default((''));"`
	Email             string    `orm:"column(Email);size(0);default((''));"`
	Qq                string    `orm:"column(Qq);size(0);default((''));"`
	Wechat            string    `orm:"column(Wechat);size(0);default((''));"`
	Weibo             string    `orm:"column(Weibo);size(0);default((''));"`
	VirtualIntegral   int64     `orm:"column(VirtualIntegral);default(((0)));"`
	RealIntegral      int64     `orm:"column(RealIntegral);default(((0)));"`
	FeeType           int64     `orm:"column(FeeType);default();"`
	PrepayValve       float64   `orm:"column(PrepayValve);digits(13);decimals(20);default(((0)));"`
	Balance           float64   `orm:"column(Balance);digits(13);decimals(20);default(((0)));"`
	FrozenBalance     float64   `orm:"column(FrozenBalance);digits(13);decimals(20);default(((0)));"`
	IncomingBalance   float64   `orm:"column(IncomingBalance);digits(13);decimals(20);default(((0)));"`
	Commission        float64   `orm:"column(Commission);digits(13);decimals(20);default(((0)));"`
	Discount          float64   `orm:"column(Discount);digits(13);decimals(20);default(((1)));"`
	Province          string    `orm:"column(Province);size(0);default();"`
	Area              string    `orm:"column(Area);size(0);default();"`
	County            string    `orm:"column(County);size(0);default();"`
	Community         string    `orm:"column(Community);size(0);default();"`
	Address           string    `orm:"column(Address);size(0);default((''));"`
	Status            int64     `orm:"column(Status);default();"`
	Type              int64     `orm:"column(Type);default();"`
	Grade             int64     `orm:"column(Grade);default();"`
	Vocation          int64     `orm:"column(Vocation);default();"`
	Reserve           string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark            string    `orm:"column(Remark);not null;size(0);default();"`
	Creator           int64     `orm:"column(Creator);default();"`
	CreateTime        time.Time `orm:"column(CreateTime);auto_now_add;type(datetime);default((getdate()));"`
	Auditor           int64     `orm:"column(Auditor);not null;default();"`
	AuditTime         time.Time `orm:"column(AuditTime);not null;auto_now_add;type(datetime);default();"`
	Canceler          int64     `orm:"column(Canceler);not null;default();"`
	CancelTime        time.Time `orm:"column(CancelTime);not null;auto_now_add;type(datetime);default();"`
	LogoIcon          string    `orm:"column(LogoIcon);size(0);default((''));"`
	LogoUrl           string    `orm:"column(LogoUrl);size(0);default((''));"`
}

//TableName
func (this *SysCorp) TableName() string {
	return "SysCorp"
}

//GetCorpByUser
func GetCorpByUser(User SysOperator) (corp SysCorp) {
	corp = SysCorp{ID: User.Corp}
	o := orm.NewOrm()
	o.Read(&corp, "ID")
	return corp
}
