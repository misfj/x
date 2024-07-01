package v1

import (
	"gorm.io/gorm"
	"time"
)

// UserInfo 用户信息表结构体
type UserInfo struct {
	ID           int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"` // 用户ID自增
	Type         string    `gorm:"column:type;type:char(2);not null"`                    // 个人或企业
	UserName     string    `gorm:"column:user_name;type:varchar(50);unique;not null"`    // 个人或企业名称，不可重复
	Email        string    `gorm:"column:email;type:varchar(30);not null"`               // 邮箱
	PhoneType    string    `gorm:"column:phone_type;type:varchar(10);not null"`          //手机号码前缀+86等
	PhoneNumber  string    `gorm:"column:phone_number;type:varchar(20);not null;unique"` // 手机号码
	Password     string    `gorm:"column:password;type:varchar(255);not null"`           // 用户密码
	CreateTime   time.Time `gorm:"column:create_time;type:datetime(6)" `                 // 用户注册时间
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime(6)" `                 // 信息更新时间
	Remark       string    `gorm:"column:remark;type:varchar(500)"`                      // 备注信息
	IsDelete     string    `gorm:"column:is_delete;type:char(1)" `                       // 账号是否注销
	Status       string    `gorm:"column:status;type:char(1)"`                           // 是否禁用
	StatusReason string    `gorm:"column:status_reason;type:varchar(255)"`               // 用户禁用原因
	UserDigital  string    `gorm:"column:user_digital;type:char(1)"`                     // 是否是数权用户
}

// 自动迁移数据库表结构
func (u *UserInfo) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

// GORM 自动迁移函数
func (u *UserInfo) TableName() string {
	return "cored_user_info"
}
