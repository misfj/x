package config

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

//	func Init(db *DB) (string, error) {
//		gormDB := connect(db)
//		scfg := config(gormDB)
//
//		closure(gormDB)
//	}
var cfg *Cfg

func Init(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(cfg)
}

func connect(db *DB) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.User,
		db.Password,
		db.Host, db.Port, db.Name)
	cli, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return cli
}
func config(db *gorm.DB) string {
	var cnf SystemServer
	if db.Model(&SystemServer{}).Where("name = ?", "cored_config").First(&cnf).Error != nil {
		fmt.Println(db.Model(&SystemServer{}).Where("name = ?", "cored_config").First(&cnf).Error)
		os.Exit(1)
	}
	return cnf.ConfigInfo

}
func closure(wrapDb *gorm.DB) {
	//关闭DB
	raw, err := wrapDb.DB()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = raw.Close()
	if err != nil {
		fmt.Println(err)
	}
}
