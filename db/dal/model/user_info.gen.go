package model

import (
	"coredx/utils"
	"errors"
	"gorm.io/gorm"
	"time"
)

const TableNameUserInfo = "user_info"

// UserInfo 用户的详细信息
type UserInfo struct {
	ID             uint           `gorm:"primary"`
	CreatedAt      time.Time      `gorm:"column:create_at"`
	UpdatedAt      time.Time      `gorm:"column:update_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:delete_at"`
	UserName       string         `gorm:"column:user_name;comment:用户名称" json:"user_name"` // 用户名称
	NickName       string         `gorm:"column:nick_name;comment:用户昵称" json:"nick_name"`
	Password       string         `gorm:"column:password;comment:用户密码" json:"password"`                                  // 用户昵称
	UserLevel      string         `gorm:"column:user_level;comment:用户级别" json:"user_level"`                              // 用户级别
	UserPhone      string         `gorm:"column:user_phone;comment:用户手机号" json:"user_phone"`                             // 用户手机号
	UserSourceType string         `gorm:"column:user_source_type;comment:1表示超级节点 2表示平台注册 3业务应用" json:"user_source_type"` // 1表示超级节点 2表示平台注册 3业务应用
	UserSource     string         `gorm:"column:user_source;comment:存储用户来源的详细信息" json:"user_source"`                     // 存储用户来源的详细信息
	UserType       string         `gorm:"column:user_type;comment:1表示个人 2表示企业 3表示政府部门" json:"user_type"`                 // 1表示个人 2表示企业 3表示政府部门
	Did            string         `gorm:"column:did;comment:用户did" json:"did"`                                           // 用户did
}

// TableName UserInfo's table name
func (*UserInfo) TableName() string {
	return TableNameUserInfo
}

var UserInfoQuery *UserInfoDao

type UserInfoDao struct {
	db *gorm.DB
}

func NewUserInfoDao(db *gorm.DB) {
	UserInfoQuery = &UserInfoDao{db: db}

}
func (info *UserInfoDao) Create(userInfo *UserInfo) error {
	nickName := userInfo.NickName
	err := info.db.Where("nick_name = ?", nickName).First(&UserInfo{}).Error
	if err != nil {
		if errors.Is(info.db.Where("nick_name = ?", nickName).First(&UserInfo{}).Error, gorm.ErrRecordNotFound) {
			return info.db.Debug().Create(userInfo).Error
		} else {
			return info.db.Where("nick_name = ?", nickName).First(&UserInfo{}).Error
		}
	} else {
		return errors.New("当前的昵称已经被占用,请重新填写")
	}
}

func (info *UserInfoDao) FindUserIDByNickName(nickName string) (uint, error) {
	var user UserInfo
	err := info.db.Model(&UserInfo{}).Where("nick_name = ?", nickName).First(&user).Error
	return user.ID, err
}

func (info *UserInfoDao) ExistByDID(did string) (uint, bool) {
	var user UserInfo
	err := info.db.Model(&UserInfo{}).Where("did =?", did).First(&user).Error
	return user.ID, err == nil
}
func (info *UserInfoDao) FindByID(id uint) (*UserInfo, error) {
	var user UserInfo
	return &user, info.db.Model(&UserInfo{}).Where("id = ?", id).First(&user).Error
}

func (info *UserInfoDao) VerifyIDPassword(id uint, password string) (bool, error) {
	if err := info.db.Model(&UserInfo{}).Where("id  = ? and password = ?", id, password).First(&UserInfo{}).Error; err != nil {
		return true, nil
	} else {
		return false, err
	}
}
func (info *UserInfoDao) VerifyDIDPassword(did string, password string) error {
	var userInfo UserInfo
	err := info.db.Model(&UserInfo{}).Where("did  =?", did).First(&userInfo).Error
	if err != nil {
		return err
	}
	verify := utils.BcryptVerify(userInfo.Password, password)
	if verify {
		return nil
	} else {
		return errors.New("did或者密码不正确")
	}
}
