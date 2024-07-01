package v1

import "time"

type UserCA struct {
	ID         int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	UserID     int64     `gorm:"column:user_id;type:bigint(20);unique"`
	UserName   string    `gorm:"column:user_name;type:varchar(50);not null;unique"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime(6)"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime(6)"`
	CreateBy   string    `gorm:"column:create_by;type:varchar(50);not null"`
	//1表示删除
	IsDelete   string `gorm:"column:is_delete;type:char(1);not null;default:0"`
	Remark     string `gorm:"column:remark;type:varchar(255)"`
	PublicKey  string `gorm:"column:public_key;type:text;not null"`
	PrivateKey string `gorm:"column:private_key;type:text;not null"`
	PublicMd5  string `gorm:"column:public_md5;type:varchar(32);not null"`
	PrivateMd5 string `gorm:"column:private_md5;type:varchar(32);not null"`
	TimeStamp  int64  `gorm:"column:time_stamp;type:bigint(20);not null"`
	StoreKey   string `gorm:"column:store_key;type:varchar(50);not null"`
	Algorithm  string `gorm:"column:algorithm;type:char(10);not null"`
	//
	Status string `gorm:"column:status;type:char(1);not null;default:0"`
}

func (u UserInfo) name() {

}
