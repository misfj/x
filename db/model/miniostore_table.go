package model

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	//addr muti cache endpoint
	Addr string `gorm:"column:addr;type:varchar(255)"`
	//bucket muti cache bucket
	Buckets    string `gorm:"column:buckets;type:varchar(255)"`
	UploadTime int64  `gorm:"column:upload_time;type:bigint(20)"`
	DID        string `gorm:"column:did;type:varchar(30)"`
	TID        string `gorm:"column:tid;type:varchar(60)"`
	File       string `gorm:"column:file;type:longtext"`
	Name       string `gorm:"column:name;type:varchar(50)"`
	Class      string `gorm:"column:class;type:varchar(50)"`
	Demo       string `gorm:"column:demo;type:varchar(255)"`
	Size       int64  `gorm:"column:size;type:bigint(20)"`
	Type       string `gorm:"column:type;type:varchar(2)"`
	Format     string `gorm:"column:format;type:varchar(10)"`
	Md5        string `gorm:"column:md5;type:varchar(50)"`
}

//func (c *Store) TableName() string {
//	return "cored_store"
//}

func (c *Store) TableName() string {
	return "cored_cache"
}
