package model

import "gorm.io/gorm"

type Back struct {
	gorm.Model
	CID       string `gorm:"column:cid;type:varchar(50);unique;not null"`
	TID       string `gorm:"column:tid;type:varchar(100)"`
	FileName  string `gorm:"column:file_name;type:varchar(255);not null"`
	OriginUrl string `gorm:"column:origin_url;type:longtext"`
	EndPoints string `gorm:"column:end_points;type:varchar(255)"`
	Size      string `gorm:"column:bytes;type:varchar(255)"`
}

func (d *Back) TableName() string {
	return "cored_back"
}
