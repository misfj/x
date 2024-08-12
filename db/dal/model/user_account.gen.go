package model

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"time"
)

const TableNameUserAccount = "user_account"

// UserAccount mapped from table <user_account>
type UserAccount struct {
	ID        uint         `gorm:"primary"`
	CreatedAt time.Time    `gorm:"column:create_at"`
	UpdatedAt time.Time    `gorm:"column:update_at"`
	DeletedAt sql.NullTime `gorm:"column:delete_at"`                                 // 删除时间
	UserID    int64        `gorm:"column:user_id;comment:用户id" json:"user_id"`       // 用户id
	Account   string       `gorm:"column:account;comment:json数组用户绑定" json:"account"` // json数组用户绑定
	Approve   string       `gorm:"column:approve;comment:记录验证账号信息" json:"approve"`   // 记录验证账号信息
}

// TableName UserAccount's table name
func (*UserAccount) TableName() string {
	return TableNameUserAccount
}

var UserAccountQuery *UserAccountDao

type UserAccountDao struct {
	db *gorm.DB
}

func NewUserAccountDao(db *gorm.DB) {
	UserAccountQuery = &UserAccountDao{db: db}
}

func (userAccount *UserAccountDao) Create(userAcccount *UserAccount) error {
	if err := userAccount.db.Model(&UserAccount{}).
		Where("user_id = ?", userAcccount.UserID).
		First(&UserAccount{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//创建数据
			return userAccount.db.Model(&UserAccount{}).Create(userAcccount).Error
		} else {
			return err
		}
	}
	return errors.New("数据存在记录,存在绑定记录")
}

//func (userAccount *UserAccountDao)
