package v1

import (
	"gorm.io/gorm"
	"time"
)

type AuthUserSpace struct {
	Id             int64     `gorm:"column:id;type:bigint(20);primary_key;comment:记录id" json:"id"`
	CreateTime     time.Time `gorm:"column:create_time;type:datetime(6);comment:创建时间" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime(6);comment:更新时间" json:"update_time"`
	IsDelete       string    `gorm:"column:is_delete;type:char(1);comment:是否删除" json:"is_delete"`
	Remark         string    `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark"`
	Status         string    `gorm:"column:status;type:char(1);comment:状态" json:"status"`
	UserId         uint64    `gorm:"column:user_id;type:bigint(20)" json:"nick_name"`
	Space          int64     `gorm:"column:space;type:int(11);comment:用户空间大小(默认MB)" json:"space"`
	UseSpace       float64   `gorm:"column:use_space;type:double;comment:已用空间" json:"use_space"`
	AvailableSpace float64   `gorm:"column:available_space;type:double;comment:可用空间" json:"available_space"`
	SpaceStatus    string    `gorm:"column:space_status;type:char(1);comment:0表示正常 1表示欠费 2表示禁用" json:"space_status"`
}

func (m *AuthUserSpace) TableName() string {
	return "auth_user_space"
}

func AuthUserSpaceCreate(db *gorm.DB, userId uint64,
	space int64, useSpace float64, availableSpace float64,
) error {
	var authUserSpace AuthUserSpace
	authUserSpace.CreateTime = time.Now()
	authUserSpace.UpdateTime = time.Now()
	authUserSpace.IsDelete = notDelete
	authUserSpace.Remark = nullRemark
	authUserSpace.Status = normalStatus
	authUserSpace.UserId = userId
	authUserSpace.Space = space
	authUserSpace.UseSpace = useSpace
	authUserSpace.AvailableSpace = availableSpace
	authUserSpace.SpaceStatus = normalSpaceStatus
	return db.Model(&authUserSpace).Create(&authUserSpace).Error
}

func AuthUserSpaceDeleteByUserId(db *gorm.DB, userId uint64) error {
	return db.Model(AuthUserSpace{}).Where("user_id = ? and is_delete = ? ", userId, notDelete).Updates(map[string]interface{}{
		"is_delete":   delete,
		"update_time": time.Now(),
		"remark":      "API删除",
	}).Error
}
