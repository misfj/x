package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

const TableNameAppLog = "app_log"

// AppLog 服务的访问日志
type AppLog struct {
	//ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:数据id" json:"id"` // 数据id
	//CreateAt    time.Time `gorm:"column:create_at;comment:创建时间" json:"create_at"`                 // 创建时间
	//UpdateAt    time.Time `gorm:"column:update_at;comment:更新时间" json:"update_at"`                 // 更新时间
	//DeleteAt    time.Time `gorm:"column:delete_at;comment:删除时间" json:"delete_at"`
	//gorm.Model         // 删除时间
	ID          uint
	CreatedAt   time.Time    `gorm:"column:create_at"`
	UpdatedAt   time.Time    `gorm:"column:update_at"`
	DeletedAt   sql.NullTime `gorm:"column:delete_at"`
	AppID       int64        `gorm:"column:app_id;comment:app用户id" json:"app_id"`          // app用户id
	ServiceType string       `gorm:"column:service_type;comment:服务类别" json:"service_type"` // 服务类别
	ReqURL      string       `gorm:"column:req_url;comment:请求路径" json:"req_url"`           // 请求路径
	Method      string       `gorm:"column:method;comment:请求方法" json:"method"`             // 请求方法
	RemoteIP    string       `gorm:"column:remote_ip;comment:访问ip" json:"remote_ip"`       // 访问ip
	ReqContent  string       `gorm:"column:req_content;comment:请求内容" json:"req_content"`   // 请求内容
	ResMsg      string       `gorm:"column:res_msg;comment:响应消息" json:"res_msg"`           // 响应消息
	ResSize     int32        `gorm:"column:res_size;comment:响应字节数" json:"res_size"`        // 响应字节数
	ResErr      string       `gorm:"column:res_err;comment:响应错误" json:"res_err"`
	Cost        int64        `gorm:"column:cost;comment:耗时(ms)" json:"cost"` // 响应错误
}

// TableName AppLog's table name
func (*AppLog) TableName() string {
	return TableNameAppLog
}

var AppLogQuery *AppLogDao

// NewAppLogDao  NewAppLogDao
func NewAppLogDao(db *gorm.DB) {
	AppLogQuery = &AppLogDao{db: db}
}

type AppLogDao struct {
	db *gorm.DB
}

func (log *AppLogDao) Create(appLog *AppLog) error {
	return log.db.Model(&AppLog{}).Create(appLog).Error
}
