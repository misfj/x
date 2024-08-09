package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

const TableNameAppDirect = "app_direct"

// AppDirect mapped from table <app_direct>
type AppDirect struct {
	ID        uint           `gorm:"primary"`
	CreatedAt time.Time      `gorm:"column:create_at"`
	UpdatedAt time.Time      `gorm:"column:update_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:delete_at"`                                 // 删除时间
	AppID     int64          `gorm:"column:app_id;comment:业务引用id" json:"app_id"`       // 业务引用id
	DirectID  int64          `gorm:"column:direct_id;comment:数据目录id" json:"direct_id"` // 数据目录id
	Status    string         `gorm:"column:status;comment:1正常访问 2禁止" json:"status"`    // 1正常访问 2禁止
}

// TableName AppDirect's table name
func (*AppDirect) TableName() string {
	return TableNameAppDirect
}

var AppDirectQuery *AppDirectDao

type AppDirectDao struct {
	db *gorm.DB
}

func NewAppDirectDao(db *gorm.DB) {
	AppDirectQuery = &AppDirectDao{
		db: db,
	}
}

func (ad *AppDirectDao) Create(data *AppDirect) error {
	if err := ad.db.Model(&AppDirect{}).Where("direct_id = ?", data.DirectID).First(&AppDirect{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ad.db.Create(data).Error
		} else {
			return err
		}
	}
	return errors.New("数据存在记录")
}
