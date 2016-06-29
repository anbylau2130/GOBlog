//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(OpenPlatformMO))
}

type OpenPlatformMO struct {
	ID           int64     `orm:"column(ID);default();index;"`
	PlatformType string    `orm:"column(PlatformType);size(0);default();"`
	ToUserName   string    `orm:"column(ToUserName);size(0);default();"`
	FromUserName string    `orm:"column(FromUserName);size(0);default();"`
	CreateTime   int64     `orm:"column(CreateTime);default();"`
	MsgType      string    `orm:"column(MsgType);size(0);default();"`
	MsgID        int64     `orm:"column(MsgID);not null;default();index;"`
	Content      string    `orm:"column(Content);not null;default();"`
	PicUrl       string    `orm:"column(PicUrl);not null;size(0);default();"`
	MediaId      string    `orm:"column(MediaId);not null;size(0);default();index;"`
	Location_X   float64   `orm:"column(Location_X);not null;digits(9);decimals(18);default(((0)));"`
	Location_Y   float64   `orm:"column(Location_Y);not null;digits(9);decimals(18);default(((0)));"`
	Scale        float64   `orm:"column(Scale);not null;digits(9);decimals(18);default(((0)));"`
	Label        string    `orm:"column(Label);not null;size(0);default();"`
	Title        string    `orm:"column(Title);not null;size(0);default();"`
	Description  string    `orm:"column(Description);not null;size(0);default();"`
	Url          string    `orm:"column(Url);not null;size(0);default();"`
	Format       string    `orm:"column(Format);not null;size(0);default();"`
	Recognition  string    `orm:"column(Recognition);not null;size(0);default();"`
	Event        string    `orm:"column(Event);not null;size(0);default();"`
	EventKey     string    `orm:"column(EventKey);not null;size(0);default();"`
	Ticket       string    `orm:"column(Ticket);not null;size(0);default();"`
	Latitude     float64   `orm:"column(Latitude);not null;digits(9);decimals(18);default(((0)));"`
	Longitude    float64   `orm:"column(Longitude);not null;digits(9);decimals(18);default(((0)));"`
	Precision    float64   `orm:"column(Precision);not null;digits(9);decimals(18);default(((0)));"`
	ReceiveTime  time.Time `orm:"column(ReceiveTime);auto_now_add;type(datetime);default((getdate()));"`
	Reserve      string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark       string    `orm:"column(Remark);not null;size(0);default();"`
	PerformTime  time.Time `orm:"column(PerformTime);not null;auto_now_add;type(datetime);default();"`
	Performer    string    `orm:"column(Performer);not null;size(0);default();"`
}

func (this *OpenPlatformMO) TableName() string {
	return "OpenPlatformMO"
}
