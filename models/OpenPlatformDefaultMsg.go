﻿//该代码使用CodeSmith
//author：tsui

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(OpenPlatformDefaultMsg))
}

type OpenPlatformDefaultMsg struct {
	ID           int64     `orm:"column(ID);default();index;"`
	Corp         int64     `orm:"column(Corp);default();"`
	PlatformType string    `orm:"column(PlatformType);size(0);default();"`
	Type         byte      `orm:"column(Type);default(((0)));"`
	Msg          int64     `orm:"column(Msg);default();"`
	Reserve      string    `orm:"column(Reserve);not null;size(0);default();"`
	Remark       string    `orm:"column(Remark);not null;size(0);default();"`
	Creator      int64     `orm:"column(Creator);default();"`
	CreateTime   time.Time `orm:"column(CreateTime);auto_now_add;type(datetime);default((getdate()));"`
	Auditor      int64     `orm:"column(Auditor);not null;default();"`
	AuditTime    time.Time `orm:"column(AuditTime);not null;auto_now_add;type(datetime);default();"`
	Canceler     int64     `orm:"column(Canceler);not null;default();"`
	CancelTime   time.Time `orm:"column(CancelTime);not null;auto_now_add;type(datetime);default();"`
}

func (this *OpenPlatformDefaultMsg) TableName() string {
	return "OpenPlatformDefaultMsg"
}