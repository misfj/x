package v1

import (
	"gorm.io/gorm"
	"time"
)

type NodeInfo struct {
	Id            int64     `gorm:"column:id;type:bigint(20);primary_key;comment:记录id" json:"id"`
	CreateTime    time.Time `gorm:"column:create_time;type:datetime(6);comment:创建时间" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime(6);comment:更新时间" json:"update_time"`
	IsDelete      string    `gorm:"column:is_delete;type:char(1);comment:是否删除 1表示删除 0表示不删除" json:"is_delete"`
	Remark        string    `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark"`
	Status        string    `gorm:"column:status;type:char(1);comment:状态1表示正常,0不正常" json:"status"`
	NodeManager   string    `gorm:"column:node_manager;type:varchar(30);comment:节点管理员名字" json:"node_manager"`
	NodeDept      string    `gorm:"column:node_dept;type:varchar(50);comment:节点所属部门" json:"node_dept"`
	NodePhoneType string    `gorm:"column:node_phone_type;type:varchar(10);comment:电话前缀,区分国际" json:"node_phone_type"`
	NodePhone     string    `gorm:"column:node_phone;type:varbinary(20);comment:联系电话" json:"node_phone"`
	NodeAddr      string    `gorm:"column:node_addr;type:varchar(255);comment:节点地址" json:"node_addr"`
	NodeOper      string    `gorm:"column:node_oper;type:varchar(50);comment:节点所在运营商" json:"node_oper"`
	LoginCode     string    `gorm:"column:login_code;type:varchar(50);comment:登录码" json:"login_code"`
	NodeId        string    `gorm:"column:node_id;type:varchar(50);comment:节点id;NOT NULL" json:"node_id"`
	NodeName      string    `gorm:"column:node_name;type:varchar(50);comment:节点名称" json:"node_name"`
	Public        string    `gorm:"column:public;type:text;comment:节点公钥;NOT NULL" json:"public"`
	Private       string    `gorm:"column:private;type:text;comment:节点私钥(一般不存储)" json:"private"`
	PublicMd5     string    `gorm:"column:public_md5;type:varchar(32);comment:公钥md5;NOT NULL" json:"public_md5"`
	PrivateMd5    string    `gorm:"column:private_md5;type:varchar(32);comment:私钥md5" json:"private_md5"`
}

func (m *NodeInfo) TableName() string {
	return "node_info"
}

func NodeInfoCreate(db *gorm.DB, createTime time.Time, updateTime time.Time, isDelete string,
	remark string, status string, nodeManager string, nodeDept string, nodePhoneType string,
	nodePhone string, nodeAddr string, nodeOper string, loginCode string, nodeId string, nodeName string, public string,
	private string, publicMd5 string, privateMd5 string) error {
	var ni NodeInfo
	ni.CreateTime = createTime
	ni.UpdateTime = updateTime
	ni.IsDelete = isDelete
	ni.Remark = remark
	ni.Status = status
	ni.NodeManager = nodeManager
	ni.NodeDept = nodeDept
	ni.NodePhoneType = nodePhoneType
	ni.NodePhone = nodePhone
	ni.NodeAddr = nodeAddr
	ni.NodeOper = nodeOper
	ni.LoginCode = loginCode
	ni.NodeId = nodeId
	ni.NodeName = nodeName
	ni.Public = public
	ni.Private = private
	ni.PublicMd5 = publicMd5
	ni.PrivateMd5 = privateMd5
	return db.Model(&NodeInfo{}).Create(&ni).Error
}
