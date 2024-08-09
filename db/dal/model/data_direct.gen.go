package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

const TableNameDataDirect = "data_direct"

// DataDirect 数据目录
type DataDirect struct {
	ID         uint           `gorm:"primary"`
	CreatedAt  time.Time      `gorm:"column:create_at"`
	UpdatedAt  time.Time      `gorm:"column:update_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:delete_at"`                                   // 删除时间
	DirectID   int64          `gorm:"column:direct_id;comment:目录id" json:"direct_id"`     // 目录id
	ParentID   int64          `gorm:"column:parent_id;comment:父目录id" json:"parent_id"`    // 父目录id
	DirectName string         `gorm:"column:direct_name;comment:目录名称" json:"direct_name"` // 目录名称
	Tag        string         `gorm:"column:tag;comment:标签" json:"tag"`                   // 标签
	Remark     string         `gorm:"column:remark;comment:备注" json:"remark"`             // 目录名称
}

// TableName DataDirect's table name
func (*DataDirect) TableName() string {
	return TableNameDataDirect
}

var DataDirectQuery *DataDirectDao

type DataDirectDao struct {
	db *gorm.DB
}

func NewDataDirectDao(db *gorm.DB) {
	DataDirectQuery = &DataDirectDao{
		db: db,
	}
}

func (dd *DataDirectDao) Create(data *DataDirect) error {
	if err := dd.db.Model(&DataDirect{}).Where("direct_id = ? and parent_id = ? and direct_name = ?",
		data.DirectID, data.ParentID, data.DirectName).First(&DataDirect{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err1 := dd.db.Model(&DataDirect{}).Where("direct_name = ?", data.DirectName).First(&DataDirect{}).Error; err1 != nil {
				if errors.Is(err1, gorm.ErrRecordNotFound) {
					return dd.db.Create(data).Error
				} else {
					return err1
				}
			}
			return errors.New("目录已经存在")

		} else {
			return err
		}
	}
	return errors.New("目录已经存在")
}
func (dd *DataDirectDao) Exist(directId int) bool {
	if err := dd.db.Model(&DataDirect{}).Where("direct_id = ?", directId).First(&DataDirect{}).Error; err != nil {
		return false
	}
	return true
}
