package global

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type SingleConn struct {
	Conn  *websocket.Conn
	mutex sync.Mutex
	//节点ID
	NodeID string
	//节点编号
	NodeNumber string
	//节点名称
	NodeName string
	//节点管理员
	NodeManager string
	//节点部门信息
	NodeDept string
	//节点联系人电话前缀
	NodeManagerPhoneType string
	//节点联系人电话号码
	NodeManagerPhone string
	//节点地址
	NodeAddr string
	//节点运营商
	NodeOper string
	//节点登陆码
	NodeLoginCode string
	//节点启动时间
	NodeStart time.Time
	// //启动断开时间
	// NodeExit time.Time
	// //断开原因
	// NodeExitReason string
	//节点公钥
	NodePubkey string
	//节点公钥MD5
	NodePubKeyMD5 string
	// //保活时间                 (默认单位是分钟)
	// NodeHealthyInterval int
}

// func (c *SingleConn) WriteJson(msg []byte) {
// 	defer func() {
// 		//连接断开,预防panic
// 		if err := recover(); err != nil {
// 			log.Error(err)
// 			return
// 		}
// 	}()
// 	c.mutex.Lock()
// 	defer c.mutex.Unlock()
// 	err := c.ws.WriteMessage(websocket.TextMessage, msg)
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}
// }
