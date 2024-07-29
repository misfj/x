package config

import "sync"

func GetDb() *Db {
	mutex.RLock()
	defer mutex.RUnlock()
	return &Config.Db
}
func GetJwt() *Jwt {
	mutex.RLock()
	defer mutex.RUnlock()
	return &Config.Jwt
}
func GetCache() *Cache {
	mutex.RLock()
	defer mutex.RUnlock()
	return &Config.Cache
}
func GetStore() *Store {
	mutex.RLock()
	defer mutex.RUnlock()
	return &Config.Store
}
func GetBack() *Backup {
	mutex.RLock()
	defer mutex.RUnlock()
	return &Config.Backup
}
func GetLogging() *Logging {
	mutex.RLock()
	defer mutex.RUnlock()
	return &Config.Logging
}
func GetDr() *Dr {
	mutex.RLock()
	defer mutex.RUnlock()
	return &Config.Dr
}
func GetApi() *API {
	mutex.RLock()
	defer mutex.RUnlock()
	return &Config.API
}
func GetWs() *Ws {
	mutex.RLock()
	defer mutex.RUnlock()
	return &Config.Ws
}
func GetWater() *Water {
	mutex.RLock()
	defer mutex.RUnlock()
	return &Config.Water
}
func GetManage() *Manage {
	mutex.RLock()
	defer mutex.RUnlock()
	return &Config.Manage
}
func GetNode() *Node {
	mutex.RLock()
	defer mutex.RUnlock()
	return &Config.Node
}

// func
var mutex sync.RWMutex

type config struct {
	Db  Db  `yaml:"db"`
	Jwt Jwt `yaml:"jwt"`

	Cache    Cache   `yaml:"cache"`
	Store    Store   `yaml:"store"`
	Backup   Backup  `yaml:"backup"`
	Logging  Logging `yaml:"logging"`
	Dr       Dr      `yaml:"dr"`
	API      API     `yaml:"api"`
	Ws       Ws      `yaml:"ws"`
	Water    Water   `yaml:"water"`
	Manage   Manage  `yaml:"manage`
	Node     Node    `yaml:"node"`
	Watertwo Water2  `yaml:"waterx"`
}

type Db struct {
	Host            string `yaml:"host"`
	Name            string `yaml:"name"`
	Port            int    `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	MaximumPoolSize int    `yaml:"maximum_pool_size"`
	MaximumIdleSize int    `yaml:"maximum_idle_size"`
}

type Jwt struct {
	Exp    int    `yaml:"exp"`
	Iss    string `yaml:"iss"`
	Sub    string `yaml:"sub"`
	Secret string `yaml:"secret"`
}

type Cache struct {
	Expire          int    `yaml:"expire"`
	UseSsl          bool   `yaml:"use_ssl"`
	AccessID        string `yaml:"access_id"`
	EndPoint        string `yaml:"end_point"`
	SecretAccessKey string `yaml:"secret_accessKey"`
}

type Store struct {
	Expire          int    `yaml:"expire"`
	UseSsl          bool   `yaml:"use_ssl"`
	AccessID        string `yaml:"access_id"`
	EndPoint        string `yaml:"end_point"`
	DefaultSize     int    `yaml:"default_size"`
	SecretAccessKey string `yaml:"secret_accessKey"`
}

type Backup struct {
	Addr string `yaml:"addr"`
}

type Logging struct {
	File      string `yaml:"file"`
	Level     string `yaml:"level"`
	MaxAge    int    `yaml:"max_age"`
	MaxSize   int    `yaml:"max_size"`
	MaxBackup int    `yaml:"max_backups"`
}

type Dr struct {
	DelData          string `yaml:"del-data"`
	Encrypt          string `yaml:"encrypt"`
	Endpoint         string `yaml:"endpoint"`
	UploadData       string `yaml:"upload-data"`
	DrUserCreate     string `yaml:"dr-user-create"`
	DrUserModify     string `yaml:"dr-user-modify"`
	DrAccountInfo    string `yaml:"dr-account-info"`
	DrAccountSync    string `yaml:"dr-account-sync"`
	GetDrUserInfo    string `yaml:"get-dr-user-info"`
	Authentication   string `yaml:"authentication"`
	SelectUploadData string `yaml:"select-upload-data"`
}

type API struct {
	KeyFile       string `yaml:"key-file"`
	CertFile      string `yaml:"cert_file"`
	EnableTLS     bool   `yaml:"enable-tls"`
	ListenAddress string `yaml:"listen-address"`
}
type Ws struct {
	ListenAddress string `yaml:"listen-address"`
}
type Water struct {
	Grid                                string `yaml:"grid"`
	First                               string `yaml:"first"`
	Plain                               string `yaml:"plain"`
	Third                               string `yaml:"third"`
	Protocol                            string `yaml:"protocol"`
	EndPoint                            string `yaml:"end_point"`
	GenerateStamps                      string `yaml:"generate-stamps"`
	StampCircleExtract                  string `yaml:"stamp-circle-extract"`
	StampElipseExtract                  string `yaml:"stamp-elipse-extract"`
	GenerateStampElipse                 string `yaml:"generate-stamp-elipse"`
	GenerateBsWaterMark                 string `yaml:"generate-bs-water-mark"`
	WeakWaterMarkInsert                 string `yaml:"weak-water-mark-insert"`
	WeakWaterMarkExtract                string `yaml:"weak-water-mark-extract"`
	ElectronicWatermarkBright           string `yaml:"electronic-watermark-bright"`
	ElectronicWatermarkUnderflow        string `yaml:"electronic-watermark-underflow"`
	ElectronicWatermarkUnderflowExtract string `yaml:"electronic-watermark-underflow-extract"`
}

type Manage struct {
	ListenAddr string `yaml:"listen-addr"`
	ListenPort int    `yaml:"listen-port"`
}

type Node struct {
	NodeNumber     string `yaml:"node-number"`
	NodeUnit       string `yaml:"node-unit"`
	NodeDeploy     string `yaml:"node-deploy"`
	NodeOper       string `yaml:"node-oper"`
	NodeLoginCode  string `yaml:"node-login-code"`
	NodePublicMd5  string `yaml:"node-public-md5"`
	NodePrivateMd5 string `yaml:"node-private-md5"`
	NodeAlgo       string `yaml:"node-algo"`
	NodeStoreKey   string `yaml:"node-store-key"`
	NodePlatform   string `yaml:"node-platform"`
}

type Water2 struct {
	Name string `yaml:"name"`
	Args string `yaml:"args"`
}
