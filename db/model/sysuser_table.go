package model

import "time"

// 用户信息表
type SysUser struct {
	UserID      int64     `gorm:"column:user_id;type:bigint(20);primary_key;AUTO_INCREMENT"` // 用户ID
	DeptID      int64     `gorm:"column:dept_id;type:bigint(20)"`                            // 部门ID
	UserName    string    `gorm:"column:user_name;type:varchar(30);NOT NULL"`                // 用户账号
	NickName    string    `gorm:"column:nick_name;type:varchar(30);NOT NULL"`                // 用户昵称
	UserType    string    `gorm:"column:user_type;type:varchar(2);default:00"`               // 用户类型（00系统用户）
	Email       string    `gorm:"column:email;type:varchar(50)"`                             // 用户邮箱
	PhoneNumber string    `gorm:"column:phone_number;type:varchar(11)"`                      // 手机号码
	Sex         string    `gorm:"column:sex;type:char(1);default:0"`                         // 用户性别（0男 1女 2未知）
	Avatar      string    `gorm:"column:avatar;type:varchar(100)"`                           // 头像地址
	Password    string    `gorm:"column:password;type:varchar(100)"`                         // 密码
	Status      string    `gorm:"column:status;type:char(1);default:0"`                      // 帐号状态（0正常 1停用）
	DelFlag     string    `gorm:"column:del_flag;type:char(1);default:0"`                    // 删除标志（0代表存在 2代表删除）
	LoginIP     string    `gorm:"column:login_ip;type:varchar(128)"`                         // 最后登录IP
	LoginDate   time.Time `gorm:"column:login_date;type:datetime"`                           // 最后登录时间
	CreateBy    string    `gorm:"column:create_by;type:varchar(64)"`                         // 创建者
	CreateTime  time.Time `gorm:"column:create_time;type:datetime"`                          // 创建时间
	UpdateBy    string    `gorm:"column:update_by;type:varchar(64)"`                         // 更新者
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime"`                          // 更新时间
	Remark      string    `gorm:"column:remark;type:varchar(500)"`                           // 备注
	PlatformID  int64     `gorm:"column:platform_id;type:bigint(20)"`                        // 用户所属节点ID
	OpenID      string    `gorm:"column:open_id;type:varchar(255)"`                          // 小程序用户的唯一标识
}

func (m *SysUser) TableName() string {
	return "sys_user"
}
