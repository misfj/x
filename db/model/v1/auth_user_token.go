package v1

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type AuthUserToken struct {
	Id              uint64    `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;comment:数据id" json:"id"`
	CreateTime      time.Time `gorm:"column:create_time;type:datetime(6);comment:创建时间" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time;type:datetime(6);comment:更新时间" json:"update_time"`
	IsDelete        string    `gorm:"column:is_delete;type:char(1);comment:是否删除(1表示删除,0不删除)" json:"is_delete"`
	Remark          string    `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark"`
	Status          string    `gorm:"column:status;type:char(1);comment:状态(1表示正常,0表示异常)" json:"status"`
	UserId          uint64    `gorm:"column:user_id;type:bigint(20);comment:用户id" `
	AccessKey       string    `gorm:"column:access_key;type:varchar(255);comment:访问凭证(放header里面)" json:"access_key"`
	AccessKeyExpire time.Time `gorm:"column:access_key_expire;type:datetime(6);comment:到期时间" json:"access_key_expire"`
	AccessKeyStatus string    `gorm:"column:access_key_status;type:varchar(10);comment:状态(正常和禁用)" json:"access_key_status"`
}

func (m *AuthUserToken) TableName() string {
	return "auth_user_token"
}
func AuthUserTokenCreate(db *gorm.DB, userID uint64, accessKey string, accessKeyExpire time.Time, accessKeyStatus string) (uint64, error) {
	var authUserToken AuthUserToken
	authUserToken.CreateTime = time.Now()
	authUserToken.UpdateTime = time.Now()
	authUserToken.IsDelete = notDelete
	authUserToken.Remark = nullRemark
	authUserToken.Status = AuthUserTokenNormalStatus
	authUserToken.UserId = userID
	authUserToken.AccessKey = accessKey
	authUserToken.AccessKeyExpire = accessKeyExpire
	authUserToken.AccessKeyStatus = accessKeyStatus
	tx := db.Model(&authUserToken).Create(&authUserToken)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return authUserToken.Id, tx.Error
}

func AuthUserTokenFindUserAccessByUserId(db *gorm.DB, userId uint64) (string, error) {
	var authUserToken AuthUserToken
	tx := db.Model(&authUserToken).Where("user_id =?  and is_delete =?", userId, notDelete).First(&authUserToken)

	return authUserToken.AccessKey, tx.Error
}

func AuthUserTokenIsExpireByAccessKey(db *gorm.DB, accessKey string) error {
	var authUserToken AuthUserToken
	tx := db.Model(&authUserToken).Where("access_key =?  and is_delete =?", accessKey, notDelete).First(&authUserToken)
	if tx.Error != nil {
		//errors.IS 速度较慢
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return errors.New("illegal accessKey")
		} else {
			return tx.Error
		}
	}
	if authUserToken.AccessKeyStatus != AuthUserTokenNormalStatus {
		return errors.New("accessKey is forbid")
	}
	if authUserToken.AccessKeyExpire.After(time.Now()) {
		return nil
	} else {
		return errors.New("accessKey is expire")
	}
}

//func AuthUserTokenGetNickNameByAccessKey(db *gorm.DB, accessKey string) (uint64, string, error) {
//	var authUserToken AuthUserToken
//	tx := db.Model(&authUserToken).Where("access_key =?  and is_delete =?", accessKey, notDelete).First(&authUserToken)
//
//	return authUserToken.Id, authUserToken.NickName, tx.Error
//
//}

func AuthUserTokenGetUserIdByAccessKey(db *gorm.DB, accessKey string) (uint64, error) {
	var authUserToken AuthUserToken
	tx := db.Model(&authUserToken).Where("access_key =? and is_delete = ?", accessKey, notDelete).First(&authUserToken)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return authUserToken.UserId, tx.Error
}
