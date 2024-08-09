package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

const TableNameAppAllow = "app_allow"

// AppAllow 业务应用许可的服务项目和期限
type AppAllow struct {
	ID           uint           `gorm:"primary"`
	CreatedAt    time.Time      `gorm:"column:create_at"`
	UpdatedAt    time.Time      `gorm:"column:update_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:delete_at"`
	AppID        int64          `gorm:"column:app_id;comment:应用id" json:"app_id"`                  // 应用id
	ServiceID    int32          `gorm:"column:service_id;comment:服务id" json:"service_id"`          // 服务id
	ServiceAllow time.Time      `gorm:"column:service_allow;comment:服务许可时间" json:"service_allow"`  // 服务许可时间
	ServiceExp   time.Time      `gorm:"column:service_exp;comment:服务过期时间" json:"service_exp"`      // 服务过期时间
	UseType      string         `gorm:"column:use_type;comment:使用方式(1次数 2流量 3期限)" json:"use_type"` // 使用方式(1次数 2流量 3期限)
	Status       string         `gorm:"column:status;comment:1表示过期2禁用3正常4暂停" json:"status"`        // 1表示过期2禁用3正常4暂停
}

// TableName AppAllow's table name
func (*AppAllow) TableName() string {
	return TableNameAppAllow
}

var AppAllowQuery *AppAllowDao

type AppAllowDao struct {
	db *gorm.DB
}

func NewAppAllowDao(db *gorm.DB) {
	AppAllowQuery = &AppAllowDao{db: db}
}

func (a *AppAllowDao) Create(dta *AppAllow) error {
	if err := a.db.Model(&AppAllow{}).Where("app_id = ? and service_id = ?", dta.AppID, dta.ServiceID).First(&AppAllow{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return a.db.Create(dta).Error
		} else {
			return err
		}
	}
	return errors.New("服务已经购买")
}
