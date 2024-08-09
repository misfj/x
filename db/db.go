package db

import (
	"coredx/config"
	"coredx/db/dal/model"
	"coredx/global"
	"coredx/log"
	"fmt"
	stdlogger "log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GDB *gorm.DB

func Init(conf *config.Db) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name)
	GDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger: logger.New(
			stdlogger.New(os.Stdout, "\r\n", stdlogger.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,          // Don't include params in the SQL log
				Colorful:                  false,         // Disable color
			},
		),
	})
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	//是否开启连接数
	sqlDb, err := GDB.DB()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	//自动迁移表结构
	sqlDb.SetMaxOpenConns(conf.MaximumPoolSize)
	sqlDb.SetMaxIdleConns(conf.MaximumIdleSize)
	log.Info("initial mysql  success")
	model.NewAppInfoDao(GDB)
	model.NewAppLogDao(GDB)
	model.NewUserInfoDao(GDB)
	model.NewUserCaDao(GDB)
	model.NewAppTokenDao(GDB)
	model.NewUserAccountDao(GDB)
	model.NewDataInfoDao(GDB)
	model.NewDataDirectDao(GDB)
	model.NewAppDirectDao(GDB)
	model.NewServiceInfoDao(GDB)
	model.NewAppAllowDao(GDB)
	model.NewDataEncDao(GDB)
	model.NewUserPdcDao(GDB)
	model.NewDataSafeDao(GDB)
	//插入默认的数据
	model.ServiceInfoQuery.Create(&model.ServiceInfo{

		ServiceName:    "水印服务",
		ServiceExecDir: "",
		ServiceVersion: "v1.0.0",
		ServiceType:    "1",
		Status:         "1",
		VisitURL:       "http:0.0.0.0:8081/api/v1/water",
		Remark:         global.WaterRemark,
		NodeID:         config.GetNode().NodePrivateMd5,
	})
	model.ServiceInfoQuery.Create(&model.ServiceInfo{

		ServiceName:    "AI服务",
		ServiceExecDir: "",
		ServiceVersion: "v1.0.0",
		ServiceType:    "1",
		Status:         "1",
		VisitURL:       "http:0.0.0.0:8081/api/v1/ai",
		Remark:         global.AIRemark,
		NodeID:         config.GetNode().NodePrivateMd5,
	})
}
