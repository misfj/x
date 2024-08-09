package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameDataSafe = "data_safe"

// DataSafe 记录数据安全相关信息
type DataSafe struct {
	ID        uint           `gorm:"primary"`
	CreatedAt time.Time      `gorm:"column:create_at"`
	UpdatedAt time.Time      `gorm:"column:update_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:delete_at"`                                              // 删除时间
	DataID    string         `gorm:"column:data_id;comment:数据id" json:"data_id"`                    // 数据id
	SafeLevel string         `gorm:"column:safe_level;comment:安全登记(1表示确权2授权3密码)" json:"safe_level"` // 安全登记(1表示确权2授权3密码)
	SafeInfo  string         `gorm:"column:safe_info;comment:授权信息确权信息密码信息" json:"safe_info"`        // 授权信息确权信息密码信息
	UserID    int64          `gorm:"column:user_id;comment:用户id" json:"user_id"`
}

// TableName DataSafe's table name
func (*DataSafe) TableName() string {
	return TableNameDataSafe
}

var DataSafeQuery *DataSafeDao

type DataSafeDao struct {
	db *gorm.DB
}

func NewDataSafeDao(db *gorm.DB) {
	DataSafeQuery = &DataSafeDao{
		db: db,
	}
}
func (ds *DataSafeDao) CloneDb() *gorm.DB {
	return ds.db
}
