package node

import (
	"context"
	"coredx/log"
	"coredx/protocol"
	"coredx/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type Node struct {
	//节点ID
	NodeID string
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
	//启动断开时间
	NodeExit time.Time
	//断开原因
	NodeExitReason string
	//节点公钥
	NodePubkey string
	//节点公钥MD5
	NodePubKeyMD5 string
	//保活时间                 (默认单位是分钟)
	NodeHealthyInterval int
}

type VNodeServer struct {
	platNodeAddr string
	ctx          context.Context
}

func NewVNodeServer(platNodeAddr string) *VNodeServer {
	platWsAddr := fmt.Sprintf("ws://%s/api/v1/node/1", platNodeAddr)
	return &VNodeServer{
		platNodeAddr: platWsAddr,
	}
}
func (v *VNodeServer) Init() error {
	return nil
}
func (v *VNodeServer) Name() string {
	return "VNodeServer"
}
func (v *VNodeServer) Startup(ctx context.Context) (err error) {
	log.Info("VNodeServer startup...")
	v.ctx = ctx
	//保证让api server 先启动
	time.Sleep(time.Second * 5)
	var conn *websocket.Conn
	timeoutCtx, _ := context.WithTimeout(context.Background(), time.Second*30)
	dial := websocket.Dialer{}
	conn, _, err = dial.DialContext(timeoutCtx, v.platNodeAddr, nil)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debugf("连接平台核心服务器成功:%s 平台地址:%s\n", conn.RemoteAddr().String(), v.platNodeAddr)
	wrapConn := NewWrapWsConn(conn, 20)
	defer func() {
		//连接断开,预防panic
		if err := recover(); err != nil {
			log.Error(err)
			return
		}
	}()
	go func(c *websocket.Conn) {
		timer := time.NewTimer(time.Second * 5)
		defer c.Close()
		defer func() {
			if err := recover(); err != nil {
				log.Error(err)
				return
			}
		}()
	loop:
		for {
			select {
			case <-v.ctx.Done():
				break loop
			case <-timer.C:
				dirty := `{"command":"Dirty"}`
				err = c.WriteMessage(websocket.TextMessage, []byte(dirty))
				timer.Reset(time.Second * 2)

			}

		}
	}(wrapConn.Conn)
	for {
		//检验是否外部cancel
		if v.ctx.Err() != nil {
			fmt.Println("VNodeServer exit")
			return nil
		}
		//err := wrapConn.Conn.SetReadDeadline(time.Now().Add(time.Second * 60))
		//log.Error(err)

		_, p, err := wrapConn.Conn.ReadMessage()
		if err != nil {
			log.Error(err)
			//log.Debugf("客户端断开的最近更新时间:%v", wrapConn.LatestHealthyTime)
			break
		}
		log.Debugf("ws client read ws msg:%s", string(p))
		cmd, err := utils.ExtractCmdValue(string(p))
		if err != nil {
			log.Error(err)
			break
		}
		//为了不协程泄露和严格控制退出速度,下级节点5s发一个与业务无关的消息
		switch cmd {
		case protocol.CMD_NODE_DIRTY_RES:

			//wrapConn.FlushDeadLine()

		}

	}
	return nil
}

func (v *VNodeServer) Close() error {
	//log.Info("VNodeServer close..")
	return nil
}
