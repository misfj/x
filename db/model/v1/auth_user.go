package v1

import (
	"coredx/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/rand/v2"
	"time"
)

type AuthUser struct {
	Id           uint64    `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;comment:id" json:"id"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime(6);comment:数据创建时间;NOT NULL" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime(6);comment:数据修改时间" json:"update_time"`
	IsDelete     string    `gorm:"column:is_delete;type:char(1);comment:1表示删除,0表示不删除" json:"is_delete"`
	Remark       string    `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark"`
	Status       string    `gorm:"column:status;type:char(4);comment:状态" json:"status"`
	PhoneTyp     string    `gorm:"column:phone_typ;type:varchar(10);comment:国际化(中国+86) 等,电话号码前缀" json:"phone_typ"`
	PhoneNum     string    `gorm:"column:phone_num;type:varchar(20);comment:电话号码" json:"phone_num"`
	Did          string    `gorm:"column:did;type:varchar(50);comment:数权账号唯一标识did" json:"did"`
	Level        string    `gorm:"column:level;type:char(1);comment:用户级别(类似于v1,v2,v3...)" json:"level"`
	NickName     string    `gorm:"column:nick_name;type:varchar(30);comment:用户名;NOT NULL" json:"nick_name"`
	AuthTime     time.Time `gorm:"column:auth_time;type:datetime(6);comment:认证时间，默认不会去认证.前端通过API触发" json:"auth_time"`
	UserSourceID string    `gorm:"column:user_sourceID;type:varchar(3);comment:用户来源(来自于平台(散户 id=0x01,资产平台id=0x02 ))" json:"user_sourceID"`
	Email        string    `gorm:"column:email;type:varchar(50);comment:用户邮箱" json:"email"`
	Password     string    `gorm:"column:password;type:varchar(255);comment:用户密码" json:"password"`
	IdCard       string    `gorm:"column:id_card;type:varchar(50);comment:身份证号" json:"id_card"`
	UserName     string    `gorm:"column:user_name;type:varchar(50);comment:真实用户名" json:"user_name"`
	BankType     string    `gorm:"column:bank_type;type:varchar(4);comment:银行卡类型" json:"bank_type"`
	BankNum      string    `gorm:"column:bank_num;type:varchar(30);comment:银行卡号" json:"bank_num"`
}

func (m *AuthUser) TableName() string {
	return "auth_user"
}
func AuthUserCreate(db *gorm.DB,
	nickName string, userName string,
	password string, phone string, PhoneType string, email string, idCard string,
	bankType string, bankNum string, userSourceID string) (uint64, error) {
	var authUser AuthUser
	authUser.CreateTime = time.Now()
	authUser.UpdateTime = time.Now()
	authUser.IsDelete = notDelete
	authUser.Remark = nullRemark
	authUser.Status = AuthUserNormalStatus
	authUser.NickName = nickName
	authUser.UserName = userName
	authUser.Password = password
	authUser.PhoneTyp = PhoneType
	authUser.PhoneNum = phone
	authUser.Email = email
	authUser.IdCard = idCard
	authUser.BankType = bankType
	authUser.BankNum = bankNum
	authUser.UserSourceID = userSourceID
	rand := rand.Int32()
	authUser.Did = fmt.Sprintf("did:%d", rand)
	authUser.AuthTime = utils.TimeParse()
	tx := db.Model(AuthUser{}).Create(&authUser)
	if tx.Error != nil {
		return 0, tx.Error
	}
	tx = db.Model(AuthUser{}).Where("nick_name =?", nickName).First(&authUser)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return authUser.Id, nil
}
func AuthUserVerifyByNickNameAndPassword(db *gorm.DB, nickName, password string) (uint64, error) {
	var authUser AuthUser
	tx := db.Model(AuthUser{}).Where("nick_name =? and is_delete = ?", nickName, notDelete).First(&authUser)
	if tx.Error != nil {
		return 0, tx.Error
	}
	x := utils.BcryptVerify(authUser.Password, password)
	if x {
		return authUser.Id, nil
	}
	return 0, errors.New("password error")
}

func AuthUserModifyByUserId(db *gorm.DB,
	phoneType string,
	PhoneNum string,
	Email string,
	Password string,
	IdCard string,
	BankType string,
	BankNum string,
	nickName string,
	userID uint64) error {
	var authUser AuthUser
	if phoneType != "" {
		authUser.PhoneTyp = phoneType
	}
	if PhoneNum != "" {
		authUser.PhoneNum = PhoneNum
	}
	if Email != "" {
		authUser.Email = Email
	}
	if Password != "" {
		bcrypt := utils.Bcrypt(Password)
		authUser.Password = bcrypt
	}
	if IdCard != "" {
		authUser.IdCard = IdCard
	}
	if BankType != "" {
		authUser.BankType = BankType
	}
	if BankNum != "" {
		authUser.BankNum = BankNum
	}
	if nickName != "" {
		authUser.NickName = nickName
	}
	return db.Model(AuthUser{}).Where("id =? and is_delete =?", userID, notDelete).Updates(&authUser).Error
}
