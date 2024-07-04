package config

import (
	"time"
)

type DB struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Cfg struct {
	Jwt          Jwt          `json:"Jwt"`
	Cache        Cache        `json:"Cache"`
	Store        Store        `json:"Store"`
	BackUp       BackUp       `json:"BackUp"`
	Logging      Logging      `json:"Logging"`
	Db           Db           `json:"Db"`
	ManageServer ManageServer `json:"ManageServer"`
	DRServer     DRServer     `json:"DRServer"`
	APIServer    APIServer    `json:"APIServer"`
	WaterServer  WaterServer  `json:"WaterServer"`
}

// Jwt Token
type Jwt struct {
	Exp    int    `json:"exp"`
	Iss    string `json:"iss"`
	Secret string `json:"secret"`
	Sub    string `json:"sub"`
}
type Cache struct {
	EndPoint  string `json:"end_point"`
	UseSSL    bool   `json:"use_ssl"`
	AccessID  string `json:"access_id"`
	SecretKey string `json:"secret_key"`
	Expire    uint8  `json:"expire"`
}
type Store struct {
	EndPoint    string `json:"end_point"`
	UseSSL      bool   `json:"use_ssl"`
	AccessID    string `json:"access_id"`
	SecretKey   string `json:"secret_key"`
	Expire      uint8  `json:"expire"`
	DefaultSize int64  `json:"default_size"`
}
type BackUp struct {
	Addr string `json:"addr"`
}
type Logging struct {
	File      string `json:"file"`
	Level     string `json:"level"`
	MaxAge    uint16 `json:"max_age"`
	MaxSize   uint16 `json:"max_size"`
	MaxBackUp uint16 `json:"max_backups"`
}
type Db struct {
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// ManageServer UDP服务器配置信息
type ManageServer struct {
	ListenIP   string `json:"ListenIP"`
	ListenPort uint16 `json:"ListenPort"`
}

// DRServer 第三方数权服务器配置信息
type DRServer struct {
	EndPoint         string `json:"endPoint"`
	Authentication   string `json:"authentication"`
	DelData          string `json:"delData"`
	DrAccountInfo    string `json:"drAccountInfo"`
	DrAccountSync    string `json:"drAccountSync"`
	DrUserCreate     string `json:"drUserCreate"`
	DrUserModify     string `json:"drUserModify"`
	Encrypt          string `json:"encrypt"`
	GetDrUserInfo    string `json:"getDrUserInfo"`
	SelectUploadData string `json:"selectUploadData"`
	UploadData       string `json:"uploadData"`
}

// WaterServer 水印服务器配置信息
type WaterServer struct {
	EndPoint                            string `json:"endPoint"`
	Grid                                string `json:"grid"`
	Plain                               string `json:"plain"`
	First                               string `json:"first"`
	Third                               string `json:"third"`
	ElectronicWatermarkUnderflow        string `json:"electronic_watermark_underflow"`
	ElectronicWatermarkBright           string `json:"electronic_watermark_bright"`
	ElectronicWatermarkUnderflowExtract string `json:"electronic_watermark_underflow_extract"`
	GenerateStamps                      string `json:"generate_stamps"`
	GenerateStampElipse                 string `json:"generate_stamp_elipse"`
	StampCircleExtract                  string `json:"stamp_circle_extract"`
	StampElipseExtract                  string `json:"stamp_elipse_extract"`
	GenerateBsWaterMark                 string `json:"generate_bs_water_mark"`

	WeakWaterMarkInsert  string `json:"weak_water_mark_insert"`
	WeakWaterMarkExtract string `json:"weak_water_mark_extract"`
}

// APIServer 应用服务器配置信息
type APIServer struct {
	ListenAddress string `json:"ListenAddress"`
	EnableTLS     bool   `json:"EnableTLS"`
	CertFile      string `json:"CertFile"`
	KeyFile       string `json:"KeyFile"`
}

// SystemServer 服务信息表 here:https://sql2gorm.mccode.info/
type SystemServer struct {
	ID           int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"` // 主键id
	CreateBy     string    `gorm:"column:create_by;type:varchar(255)"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime(6)"`   // 创建时间
	CreateUserID int64     `gorm:"column:create_user_id;type:bigint(20)"` // 创建人id
	IsDelete     int       `gorm:"column:is_delete;type:int(11)"`         // 1标识删除
	Remark       string    `gorm:"column:remark;type:varchar(255)"`       // 备注
	Status       int       `gorm:"column:status;type:int(4);default:0"`   // 状态
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime(6)"`   // 更新时间
	Name         string    `gorm:"column:name;type:varchar(255)"`         // 名称
	ConfigInfo   string    `gorm:"column:config_info;type:json"`
	Type         int       `gorm:"column:type;type:int(11)"`      // 类型，字典system_server_type
	Addr         string    `gorm:"column:addr;type:varchar(255)"` // 地址
}

func JwtConfig() *Jwt {
	return &cfg.Jwt
}
func CacheConfig() *Cache {
	return &cfg.Cache
}

func StoreConfig() *Store {
	return &cfg.Store
}
func BackupConfig() *BackUp {
	return &cfg.BackUp
}
func LoggingConfig() *Logging {
	return &cfg.Logging
}
func DbConfig() *Db {
	return &cfg.Db
}
func ManageServerConfig() *ManageServer {
	return &cfg.ManageServer
}
func DRConfig() *DRServer {
	return &cfg.DRServer
}
func APIConfig() *APIServer {
	return &cfg.APIServer
}
