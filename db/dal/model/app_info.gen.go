package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAppInfo = "app_info"

// AppInfo 业务应用的基本信息
type AppInfo struct {
	//gorm.Model
	ID        uint           `gorm:"primary"`
	CreatedAt time.Time      `gorm:"column:create_at"`
	UpdatedAt time.Time      `gorm:"column:update_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:delete_at"`
	// 删除时间
	AppName       string    `gorm:"column:app_name;comment:业务应用名称" json:"app_name"`                // 业务应用名称
	AppUser       string    `gorm:"column:app_user;comment:应用联系人" json:"app_user"`                 // 应用联系人
	AppPassword   string    `gorm:"column:app_password;comment:应用密码" json:"app_password"`          //应用密码
	AppPhone      string    `gorm:"column:app_phone;comment:联系电话" json:"app_phone"`                // 联系电话
	AppType       string    `gorm:"column:app_type;comment:业务应用类型(1文旅2电商3资产4交易)" json:"app_type"`  // 业务应用类型(1文旅2电商3资产4交易)
	AppPublic     string    `gorm:"column:app_public;comment:应用公钥" json:"app_public"`              // 应用公钥
	AppPublicMd5  string    `gorm:"column:app_public_md5;comment:应用公钥md5" json:"app_public_md5"`   // 应用公钥md5
	AppPrivate    string    `gorm:"column:app_private;comment:应用私钥" json:"app_private"`            // 应用私钥
	AppPrivateMd5 string    `gorm:"column:app_private_md5;comment:应用私钥md5" json:"app_private_md5"` // 应用私钥md5
	Algo          string    `gorm:"column:algo;comment:算法" json:"algo"`                            // 算法
	StoreKey      string    `gorm:"column:store_key;comment:应用私钥" json:"store_key"`                // 应用私钥
	StartTime     time.Time `gorm:"column:start_time;comment:开始时间" json:"start_time"`              // 开始时间
	Status        string    `gorm:"column:status;comment:1表示正常 2禁用 3暂停" json:"status"`             // 1表示正常 2禁用 3暂停
}

// TableName AppInfo's table name
func (*AppInfo) TableName() string {
	return TableNameAppInfo
}

var AppInfoQuery *AppInfoDao

func NewAppInfoDao(db *gorm.DB) {
	AppInfoQuery = &AppInfoDao{db: db}
}

type AppInfoDao struct {
	db *gorm.DB
}

func (info *AppInfoDao) Create(appInfo *AppInfo) error {
	return info.db.Create(appInfo).Error
}
func (info *AppInfoDao) FindByAppName(appName string) (app *AppInfo, err error) {
	return app, info.db.Where("app_name = ?", appName).First(&app).Error
}
func (info *AppInfoDao) VerifyUserPassword(appName string, appPassword string) bool {
	return info.db.Where("app_name = ?  and password = ?", appName, appPassword).Error == nil
}

func (info *AppInfoDao) List(pageNum int, pageSize int) ([]AppInfo, error) {
	var appInfos []AppInfo
	return appInfos, info.db.Model(&AppInfo{}).
		Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&appInfos).Error
}
