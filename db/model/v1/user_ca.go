package v1

import (
	"gorm.io/gorm"
	"time"
)

type UserCa struct {
	Id         uint64    `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;comment:记录ID" json:"id"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime(6);comment:创建时间;NOT NULL" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime(6);comment:更新时间" json:"update_time"`
	IsDelete   string    `gorm:"column:is_delete;type:char(1);comment:1表示删除,0表示不删除" json:"is_delete"`
	Remark     string    `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark"`
	Status     string    `gorm:"column:status;type:char(1);comment:状态(1表示可用,0表示禁用)" json:"status"`
	//NickName   string    `gorm:"column:nick_name;type:varchar(50);comment:用户名;NOT NULL" json:"nick_name"`
	UserId     uint64 `gorm:"column:user_id;type:bigint(20) unsigned;comment:用户ID(auth_user表里的id)"`
	Public     string `gorm:"column:public;type:text;comment:用户公钥" json:"public"`
	Private    string `gorm:"column:private;type:text;comment:用户私钥" json:"private"`
	PublicMd5  string `gorm:"column:public_md5;type:varchar(32);comment:公钥md5" json:"public_md5"`
	PrivateMd5 string `gorm:"column:private_md5;type:varchar(32);comment:私钥md5" json:"private_md5"`
	StoreKey   string `gorm:"column:store_key;type:varchar(50);comment:sm2存储密码(rsa不需要存空)" json:"store_key"`
	Algorithm  string `gorm:"column:algorithm;type:varchar(10);comment:算法" json:"algorithm"`
	TimeStamp  int64  `gorm:"column:time_stamp;type:bigint(20);comment:证书时间戳"`
}

func (m *UserCa) TableName() string {
	return "user_ca"
}

func UserCaCreate(db *gorm.DB, userID uint64,
	public string,
	private string, publicMd5 string,
	privateMd5 string, storeKey string, algorithm string, timeStamp int64) error {
	var userCa UserCa
	userCa.CreateTime = time.Now()
	userCa.UpdateTime = time.Now()
	userCa.IsDelete = notDelete
	userCa.Remark = nullRemark
	userCa.Status = normalStatus
	//userCa.NickName = nickName
	userCa.UserId = userID
	userCa.Public = public
	userCa.Private = private
	userCa.PublicMd5 = publicMd5
	userCa.PrivateMd5 = privateMd5
	userCa.StoreKey = storeKey
	userCa.Algorithm = algorithm
	userCa.TimeStamp = timeStamp
	return db.Create(&userCa).Error
}

func UserCaDeleteByUserIds(db *gorm.DB, userId uint64) error {
	return db.Model(&UserCa{}).Where("user_id = ? and is_delete = ? ", userId, notDelete).Updates(map[string]interface{}{"is_delete": delete, "update_time": time.Now(),
		"remark": "API删除"}).Error
}

//func UserCaFindPrivateKeyByUserId(db *gorm.DB, userId uint64) (string, error) {
//	return "", fmt.Errorf("not found private key")
//}
