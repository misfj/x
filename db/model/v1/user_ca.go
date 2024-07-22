package v1

import (
	"time"

	"gorm.io/gorm"
)

type UserCa struct {
	Id         uint64    `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"` // 数据id
	CreateAt   time.Time `gorm:"column:create_at;type:datetime(3)" json:"create_at"`                      // 创建时间
	UpdateAt   time.Time `gorm:"column:update_at;type:datetime(3)" json:"update_at"`                      // 更新时间
	DeleteAt   time.Time `gorm:"column:delete_at;type:datetime(3)" json:"delete_at"`                      // 删除时间
	UserId     time.Time `gorm:"column:user_id;type:bigint(20)" json:"user_id"`                           // 创建时间
	Public     time.Time `gorm:"column:public;type:text" json:"public"`                                   // 公钥d
	PublicMd5  time.Time `gorm:"column:public_md5;type:varchar(32)" json:"public_md5"`                    // 公钥md5
	Private    time.Time `gorm:"column:private;type:text" json:"private"`                                 // 私钥
	PrivateMd5 time.Time `gorm:"column:private_md5;type:varchar(32)" json:"private_md5"`                  // 私钥md5
	Algo       time.Time `gorm:"column:algo;type:varchar(10)" json:"algo"`                                // 算法
	StoreKey   time.Time `gorm:"column:store_key;type:varchar(32)" json:"store_key"`                      // 存储密码
	TimeStamp  time.Time `gorm:"column:time_stamp;type:bigint(20)" json:"time_stamp"`                     // 时间戳
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
