package v1

import (
	"gorm.io/gorm"
	"time"
)

type NodeRegist struct {
	Id         uint64    `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;comment:记录id" json:"id"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime(6);comment:创建时间" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime(6);comment:更新时间" json:"update_time"`
	IsDelete   string    `gorm:"column:is_delete;type:char(1);comment:是否删除 1删除,0不删除" json:"is_delete"`
	Remark     string    `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark"`
	Status     string    `gorm:"column:status;type:char(1);comment:状态" json:"status"`
	NodeId     string    `gorm:"column:node_id;type:varchar(50);comment:节点id" json:"node_id"`
	NodeName   string    `gorm:"column:node_name;type:varchar(50);comment:节点名称" json:"node_name"`
	NodeIp     string    `gorm:"column:node_ip;type:varchar(50);comment:登录ip" json:"node_ip"`
	NodeStart  time.Time `gorm:"column:node_start;type:datetime(6);comment:节点接入时间" json:"node_start"`
	NodeExit   time.Time `gorm:"column:node_exit;type:datetime(6);comment:节点断开时间" json:"node_exit"`
}

func (m *NodeRegist) TableName() string {
	return "node_regist"
}

func NodeRegistCreate(db *gorm.DB, createTime time.Time, updateTime time.Time,
	IsDelete string, Status string, remark string, nodeId string, nodeName string, nodeIP string,
	nodeStart time.Time, nodeExit time.Time) error {
	var ng NodeRegist
	ng.CreateTime = createTime
	ng.UpdateTime = updateTime
	ng.IsDelete = IsDelete
	ng.Status = Status
	ng.Remark = remark
	ng.NodeId = nodeId
	ng.NodeName = nodeName
	ng.NodeIp = nodeIP
	ng.NodeStart = nodeStart
	ng.NodeExit = nodeExit
	return db.Model(&NodeRegist{}).Create(&ng).Error
}
