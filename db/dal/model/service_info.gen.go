package model

import (
	"coredx/log"
	"errors"
	"gorm.io/gorm"
	"time"
)

const TableNameServiceInfo = "service_info"

// ServiceInfo 服务代理的基本信息,包含程序
type ServiceInfo struct {
	ID             uint           `gorm:"primary"`
	CreatedAt      time.Time      `gorm:"column:create_at"`
	UpdatedAt      time.Time      `gorm:"column:update_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:delete_at"`                                                  // 删除时间
	ServiceName    string         `gorm:"column:service_name;comment:服务名称" json:"service_name"`              // 服务名称
	ServiceExecDir string         `gorm:"column:service_exec_dir;comment:程序运行目录" json:"service_exec_dir"`    // 程序运行目录
	ServiceVersion string         `gorm:"column:service_version;comment:服务版本" json:"service_version"`        // 服务版本
	ServiceType    string         `gorm:"column:service_type;comment:服务类别(1水印 2智能 3数权)" json:"service_type"` // 服务类别(1水印 2智能 3数权)
	Status         string         `gorm:"column:status;comment:1表示正常 2表示异常(不能用)" json:"status"`              // 1表示正常 2表示异常(不能用)
	VisitURL       string         `gorm:"column:visit_url;comment:访问地址" json:"visit_url"`                    // 访问地址
	Remark         string         `gorm:"column:remark;comment:备注" json:"remark"`
	NodeID         string         `gorm:"column:node_id;comment:节点id" json:"node_id"` // 节点id
}

// TableName ServiceInfo's table name
func (*ServiceInfo) TableName() string {
	return TableNameServiceInfo
}

var ServiceInfoQuery *ServiceInfoDao

type ServiceInfoDao struct {
	db *gorm.DB
}

func NewServiceInfoDao(db *gorm.DB) {
	ServiceInfoQuery = &ServiceInfoDao{
		db: db,
	}
}
func (s *ServiceInfoDao) Create(dta *ServiceInfo) error {
	if err := s.db.Model(&ServiceInfo{}).
		Where("service_name = ?", dta.ServiceName).First(&ServiceInfo{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return s.db.Create(dta).Error
		} else {
			return err
		}
	}
	return errors.New("服务名称在数据库已经存在")
}
func (s *ServiceInfoDao) List() ([]ServiceInfo, error) {
	var serviceInfos []ServiceInfo
	if err := s.db.Model(&ServiceInfo{}).Find(&serviceInfos).Error; err != nil {
		return nil, err
	}
	log.Debug(serviceInfos)
	return serviceInfos, nil
}
func (info *ServiceInfoDao) ExistServiceID(serviceID int) bool {
	var service ServiceInfo
	if err := info.db.Debug().Model(&ServiceInfo{}).Where("id = ?", serviceID).
		First(&service).Error; err != nil {

		return false
	}
	return true

}
