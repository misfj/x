package model

import (
	"coredx/log"
	"errors"
	"gorm.io/gorm"
	"time"
)

const TableNameUserPdc = "user_pdc"

// UserPdc 用户数据空间变化记录表
type UserPdc struct {
	ID             uint           `gorm:"primary"`
	CreatedAt      time.Time      `gorm:"column:create_at"`
	UpdatedAt      time.Time      `gorm:"column:update_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:delete_at"`
	UserID         int64          `gorm:"column:user_id;comment:用户id" json:"user_id"`                     // 用户id
	SpaceTotal     int32          `gorm:"column:space_total;comment:默认单位MB" json:"space_total"`           // 默认单位MB
	SpaceUse       float64        `gorm:"column:space_use;comment:使用空间单位字节" json:"space_use"`             // 使用空间单位字节
	SpaceAvailable float64        `gorm:"column:space_available;comment:剩余空间单位字节" json:"space_available"` // 剩余空间单位字节
	Status         string         `gorm:"column:status;comment:1正常 2禁用" json:"status"`                    // 1正常 2禁用
}

// TableName UserPdc's table name
func (*UserPdc) TableName() string {
	return TableNameUserPdc
}

var UserPdcQuery *UserPdcDao

type UserPdcDao struct {
	db *gorm.DB
}

func NewUserPdcDao(db *gorm.DB) {
	UserPdcQuery = &UserPdcDao{db: db}
}
func (updc *UserPdcDao) CloneDb() *gorm.DB {
	return updc.db
}
func (updc *UserPdcDao) Create(pdc *UserPdc) error {
	if err := updc.db.Model(&UserInfo{}).Where("user_id = ?", pdc.UserID).First(&UserPdc{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return updc.db.Create(pdc).Error
		} else {
			return err
		}
	}
	return errors.New("数据库已存在信息")
}
func (updc *UserPdcDao) FindByUserID(userID int) (*UserPdc, error) {
	var userPdc UserPdc
	if err := updc.db.Model(&UserPdc{}).Where("user_id = ? ", userID).First(&userPdc).Error; err != nil {
		return nil, err
	}
	return &userPdc, nil
}

func (updc *UserPdcDao) ExpandCapacity(userID int64, size float64) (*UserPdc, error) {
	//todo 在创建用户的时候请补全创建用pdc记录
	var userPdc UserPdc
	if err := updc.db.Model(&UserPdc{}).Where("user_id = ?", userID).First(&userPdc).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	//更新空间
	userPdc.SpaceAvailable = userPdc.SpaceAvailable + size
	userPdc.SpaceTotal = userPdc.SpaceTotal + int32(size)
	userPdc.UpdatedAt = time.Now()
	if err := updc.db.Model(&UserPdc{}).Where("user_id = ?", userID).Save(&userPdc).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return &userPdc, nil
}
