package db

import (
	"coredx/config"
	v1 "coredx/db/model/v1"
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
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	err = GDB.AutoMigrate(&v1.UserInfo{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	//是否开启连接数
	//s, err := db.DB()
	//s.SetMaxIdleConns()
	//自动迁移表结构
}
