package db

import (
	"coredx/config"
	"coredx/log"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GDB *gorm.DB

func Init(conf *config.Db) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name)
	GDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: log.DefaultLogger,
		PrepareStmt: true,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//是否开启连接数
	sqlDb, err := GDB.DB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//s.SetMaxIdleConns()
	//自动迁移表结构
	sqlDb.SetMaxOpenConns(conf.MaximumPoolSize)
	sqlDb.SetMaxIdleConns(conf.MaximumIdleSize)
	log.Info("initial mysql  success")

}
