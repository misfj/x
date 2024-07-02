package ws

import (
	"context"
	"coredx/log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var Node node

type node struct {
	conn *websocket.Conn
	mux  sync.Mutex

	//节点ID
	nodeID string
	//节点名称
	nodeName string
	//节点管理员
	nodeManager string
	//节点联系人电话前缀
	nodeManagerPhoneType string
	//节点联系人电话号码
	nodeManagerPhone string
	//节点启动时间
	nodeStart time.Time
	//启动断开时间
	nodeExit time.Time
	//节点公钥
	nodePubkey string
	//节点公钥MD5
	nodePubKeyMD5 string
	//节点保活时间
	nodeKeepLiveInternal int
	//节点授权码
	authCode string
}

func (c *node) SyncWrite(msg []byte) {
	defer c.mux.Lock()
	err := c.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Error(err)
		return
	}
}

func Connect(platAddr string) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	//连接平台核心服务器地址
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, platAddr, nil)
	if err != nil {
		// log.Error("websocket dial:", err)
		// fmt.Println(err)
		log.Errorf("websocket dial error:%s", err.Error())
		return
	}
	defer conn.Close()
	wsC.conn = conn
}
