package v1

import (
	"time"
)

type AppGrant struct {
	Id          int64     `gorm:"column:id;type:bigint(20);primary_key;comment:记录id" json:"id"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime(6);comment:创建时间" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime(6);comment:更新时间" json:"update_time"`
	IsDelete    string    `gorm:"column:is_delete;type:char(1);comment:是否删除" json:"is_delete"`
	Remark      string    `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark"`
	Status      string    `gorm:"column:status;type:char(1);comment:状态" json:"status"`
	AppId       string    `gorm:"column:app_id;type:varchar(50);comment:业务应用程序ID" json:"app_id"`
	GrantLevel  string    `gorm:"column:grant_level;type:char(1);comment:账号级别(v1,v2,v3)" json:"grant_level"`
	GrantTime   time.Time `gorm:"column:grant_time;type:datetime(6);comment:授权时间" json:"grant_time"`
	GrantExpire time.Time `gorm:"column:grant_expire;type:datetime(6);comment:授权期限" json:"grant_expire"`
	GrantCode   time.Time `gorm:"column:grant_code;type:varchar(50);comment:授权码" json:"grant_code"`
}

func (m *AppGrant) TableName() string {
	return "app_grant"
}
