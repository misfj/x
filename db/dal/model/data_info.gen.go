package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

const TableNameDataInfo = "data_info"

// DataInfo 数据是否共享 是否需要审核 数据价值
type DataInfo struct {
	ID             uint           `gorm:"primary"`
	CreatedAt      time.Time      `gorm:"column:create_at"`
	UpdatedAt      time.Time      `gorm:"column:update_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:delete_at"`                                                 // 删除时间
	UseID          int64          `gorm:"column:use_id;comment:用户id" json:"use_id"`                         // 用户id
	DataID         string         `gorm:"column:data_id;comment:数据散列" json:"data_id"`                       // 数据散列
	DataName       string         `gorm:"column:data_name;comment:数据名称" json:"data_name"`                   // 数据名称
	DataDiectoryID int32          `gorm:"column:data_diectory_id;comment:数据目录" json:"data_diectory_id"`     // 数据目录
	DataType       string         `gorm:"column:data_type;comment:1表示文档2表示图片3表示视频4表示关系数据" json:"data_type"` // 1表示文档2表示图片3表示视频4表示关系数据
	IsOpen         string         `gorm:"column:is_open;comment:1表示开放2表示关闭" json:"is_open"`                 // 1表示开放2表示关闭
	DataSize       int64          `gorm:"column:data_size;comment:数据大小" json:"data_size"`                   // 数据大小
	DataPos        string         `gorm:"column:data_pos;comment:数据位置" json:"data_pos"`                     // 数据位置
	NodeID         string         `gorm:"column:node_id;comment:节点id" json:"node_id"`
	Remark         string         `gorm:"column:remark;comment:介绍" json:"remark"`
}

// TableName DataInfo's table name
func (*DataInfo) TableName() string {
	return TableNameDataInfo
}

var DataInfoQuery *DataInfoDao

type DataInfoDao struct {
	db *gorm.DB
}

func (dinfo *DataInfoDao) CloneDb() *gorm.DB {
	return dinfo.db
}
func NewDataInfoDao(db *gorm.DB) {
	DataInfoQuery = &DataInfoDao{db: db}
}

func (dinfo *DataInfoDao) Create(info *DataInfo) error {
	dataID := info.DataID
	var dataInfo DataInfo
	if err := dinfo.db.Model(&DataInfo{}).Where("data_id = ? and use_id = ?", dataID, info.UseID).First(&dataInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dinfo.db.Create(info).Error
		} else {
			return err
		}
	}
	return nil
}

func (dinfo *DataInfoDao) List(pageNum int, pageSize int, conditions string) ([]DataInfo, error) {
	var dataInfos []DataInfo
	return dataInfos, dinfo.db.
		Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&dataInfos).Error
}

func (dinfo *DataInfoDao) FindByDataID(dataID string) (*DataInfo, error) {
	var dataInfo DataInfo
	err := dinfo.db.Model(&dataInfo).Where("dataID = ? ", dataID).First(&dataInfo).Error
	if err != nil {
		return nil, err
	}
	return &dataInfo, nil
}
