package model

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"time"
)

const TableNameAppToken = "app_token"

// AppToken mapped from table <app_token>
type AppToken struct {
	ID        uint         `gorm:"primary"`
	CreatedAt time.Time    `gorm:"column:create_at"`
	UpdatedAt time.Time    `gorm:"column:update_at"`
	DeletedAt sql.NullTime `gorm:"column:delete_at"`                               // 删除时间
	AppID     int64        `gorm:"column:app_id;comment:业务应用id" json:"app_id"`     // 业务应用id
	Token     string       `gorm:"column:token;comment:token" json:"token"`        // token
	TokenExp  time.Time    `gorm:"column:token_exp;comment:过期时间" json:"token_exp"` // 过期时间
}

// TableName AppToken's table name
func (*AppToken) TableName() string {
	return TableNameAppToken
}

var AppTokenQuery *AppTokenDao

type AppTokenDao struct {
	db *gorm.DB
}

func NewAppTokenDao(db *gorm.DB) {
	AppTokenQuery = &AppTokenDao{db: db}
}

func (token *AppTokenDao) Create(tk *AppToken) error {
	appID := tk.AppID
	if err := token.db.Model(&AppToken{}).Where("app_id = ?", appID).First(&AppToken{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//token 不存在创建token
			return token.db.Model(&AppToken{}).Create(tk).Error
		} else {
			return err
		}
	} else {
		return nil
	}
	//return nil
}

const (
	TokenNotFound = "TokenNotFound"
)

func (token *AppTokenDao) FindByAppID(appID int) (string, error) {
	var tk AppToken
	if err := token.db.Where("app_id =?", appID).First(&tk).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", gorm.ErrRecordNotFound
		} else {
			return "", err
		}
	}
	return tk.Token, nil
}
