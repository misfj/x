package model

import "gorm.io/gorm"

func SysUserFindByUsername(db *gorm.DB, username string) (*SysUser, error) {
	var user SysUser
	return &user, db.Model(&SysUser{}).Where("user_name = ? ", username).First(&user).Error
}
