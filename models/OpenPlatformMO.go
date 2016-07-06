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
	orm.RegisterModel(new(OpenPlatformMO))
}

type OpenPlatformMO struct {
	ID           int64     `orm:"column(ID);pk;unique;index;auto;"`
	PlatformType string    `orm:"column(PlatformType);size(100);"`
	ToUserName   string    `orm:"column(ToUserName);size(32);"`
	FromUserName string    `orm:"column(FromUserName);size(32);"`
	CreateTime   int64     `orm:"column(CreateTime);"`
	MsgType      string    `orm:"column(MsgType);size(50);"`
	MsgID        int64     `orm:"column(MsgID);not null;"`
	Content      string    `orm:"column(Content);not null;"`
	PicUrl       string    `orm:"column(PicUrl);not null;size(200);"`
	MediaId      string    `orm:"column(MediaId);not null;size(100);"`
	Location_X   float64   `orm:"column(Location_X);not null;digits(28);decimals(8);"`
	Location_Y   float64   `orm:"column(Location_Y);not null;digits(28);decimals(8);"`
	Scale        float64   `orm:"column(Scale);not null;digits(28);decimals(8);"`
	Label        string    `orm:"column(Label);not null;size(100);"`
	Title        string    `orm:"column(Title);not null;size(100);"`
	Description  string    `orm:"column(Description);not null;size(200);"`
	Url          string    `orm:"column(Url);not null;size(100);"`
	Format       string    `orm:"column(Format);not null;size(100);"`
	Recognition  string    `orm:"column(Recognition);not null;size(100);"`
	Event        string    `orm:"column(Event);not null;size(100);"`
	EventKey     string    `orm:"column(EventKey);not null;size(100);"`
	Ticket       string    `orm:"column(Ticket);not null;size(100);"`
	Latitude     float64   `orm:"column(Latitude);not null;digits(28);decimals(8);"`
	Longitude    float64   `orm:"column(Longitude);not null;digits(28);decimals(8);"`
	Precision    float64   `orm:"column(Precision);not null;digits(28);decimals(8);"`
	ReceiveTime  time.Time `orm:"column(ReceiveTime);auto_now_add;type(datetime);"`
	Reserve      string    `orm:"column(Reserve);not null;size(250);"`
	Remark       string    `orm:"column(Remark);not null;size(250);"`
	PerformTime  time.Time `orm:"column(PerformTime);not null;auto_now_add;type(datetime);"`
	Performer    string    `orm:"column(Performer);not null;size(20);"`
}

func (this *OpenPlatformMO) TableName() string {
	return "OpenPlatformMO"
}

func (this *OpenPlatformMO) Add() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(this)
	return id, err
}

func (this *OpenPlatformMO) Count(condation *orm.Condition) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	return qs.Count()
}

func (this *OpenPlatformMO) Update(cols ...string) (num int64, err error) {
	o := orm.NewOrm()
	if o.Read(this) == nil {
		num, err := o.Update(this, cols...)
		return num, err
	}
	return 0, errors.New("找不到ID=‘" + string(this.ID) + "’的数据!")
}

func (this *OpenPlatformMO) Delete() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(this)
	return num, err
}

func (this *OpenPlatformMO) Read(cols ...string) (*OpenPlatformMO, error) {
	o := orm.NewOrm()
	err := o.Read(this, cols...)
	if err != nil {
		return this, err
	}
	return this, nil
}

func (this *OpenPlatformMO) GetAll(condation *orm.Condition, sort string) (models []OpenPlatformMO) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	qs.All(&models)
	return models
}

func (this *OpenPlatformMO) Getlist(condation *orm.Condition, page int64, page_size int64, sort string) (models []orm.Params, count int64) {
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

func (this *OpenPlatformMO) Validation() (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&this)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}
