package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

const TableNameUserCa = "user_ca"

// UserCa 用用户公钥证书 日期 算法 存储密钥
type UserCa struct {
	//ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:数据id" json:"id"` // 数据id
	//CreateAt   time.Time `gorm:"column:create_at;comment:创建时间" json:"create_at"`                 // 创建时间
	//UpdateAt   time.Time `gorm:"column:update_at;comment:更新时间" json:"update_at"`                 // 更新时间
	//DeleteAt   time.Time `gorm:"column:delete_at;comment:删除时间" json:"delete_at"`                 // 删除时间
	ID         uint           `gorm:"primary"`
	CreatedAt  time.Time      `gorm:"column:create_at"`
	UpdatedAt  time.Time      `gorm:"column:update_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:delete_at"`
	UserID     int64          `gorm:"column:user_id;comment:创建时间" json:"user_id"`          // 创建时间
	Public     string         `gorm:"column:public;comment:公钥d" json:"public"`             // 公钥d
	PublicMd5  string         `gorm:"column:public_md5;comment:公钥md5" json:"public_md5"`   // 公钥md5
	Private    string         `gorm:"column:private;comment:私钥" json:"private"`            // 私钥
	PrivateMd5 string         `gorm:"column:private_md5;comment:私钥md5" json:"private_md5"` // 私钥md5
	Algo       string         `gorm:"column:algo;comment:算法" json:"algo"`                  // 算法
	StoreKey   string         `gorm:"column:store_key;comment:存储密码" json:"store_key"`      // 存储密码
	TimeStamp  int64          `gorm:"column:time_stamp;comment:时间戳" json:"time_stamp"`     // 时间戳
}

// TableName UserCa's table name
func (*UserCa) TableName() string {
	return TableNameUserCa
}

var UserCaQuery *UserCaDao

type UserCaDao struct {
	db *gorm.DB
}

func NewUserCaDao(db *gorm.DB) {
	UserCaQuery = &UserCaDao{db: db}
}
func (ca *UserCaDao) Create(caa *UserCa) (err error) {
	//查询数据库是否存在密钥队
	err = ca.db.Model(&UserCa{}).Where("public_md5 = ? and private_md5 = ?", caa.PublicMd5, caa.PrivateMd5).First(&UserCa{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ca.db.Model(&UserCa{}).Create(caa).Error
		} else {
			return err
		}
	}
	return errors.New("数据库已经存在该密钥对")
}

func (ca *UserCaDao) FindByUserID(userID uint) (*UserCa, error) {
	var uca UserCa
	return &uca, ca.db.Where("user_id =?", userID).First(&uca).Error
}
