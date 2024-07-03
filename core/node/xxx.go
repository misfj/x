package node

import (
	"coredx/log"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

func NewWrapWsConn(conn *websocket.Conn, interval int) *WrapWsConn {
	return &WrapWsConn{
		Conn:     conn,
		internal: interval,
	}
}
func (w *WrapWsConn) FlushDeadLine() {
	defer func() {
		if e := recover(); e != nil {
			log.Errorf("超级大错误,请联系超级节点负责人:%v", e)
			return
		}
	}()
	err := w.Conn.SetReadDeadline(time.Now().Add(time.Duration(w.internal) * time.Second))
	if err != nil {
		log.Errorf("超级大错误,请联系超级节点负责人:%v", err)
		return
	}
}

type WrapWsConn struct {
	//gorilla websocket 包非并发安全,所以加锁
	Conn     *websocket.Conn
	mux      sync.Mutex
	NodeID   string
	NodeName string
	//保活时间
	internal int
}

func (w *WrapWsConn) Write(typ int, p []byte) {
	defer func() {
		if e := recover(); e != nil {
			log.Errorf("超级大错误,请联系超级节点负责人:%v", e)
			return
		}
	}()
	defer w.mux.Unlock()
	w.mux.Lock()
	err := w.Conn.WriteMessage(typ, p)
	if err != nil {
		log.Error(err)
	}
}

// NodeService func Establish(w http.ResponseWriter, r *http.Request) {
//func NodeService(ctx *gin.Context) {
//	upgrader := websocket.Upgrader{
//		CheckOrigin: func(r *http.Request) bool {
//			return true // 允许所有CORS请求，实际使用时应该更严格
//		},
//		ReadBufferSize:   4096,
//		WriteBufferSize:  4096,
//		HandshakeTimeout: 30 * time.Second,
//	}
//	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
//	if err != nil {
//		log.Error(err)
//		return
//	}
//	wrapConn := NewWrapWsConn(c, 30)
//	wrapConn.FlushDeadLine()
//	handleNode(wrapConn)
//}

//func handleNode(wrap *WrapWsConn) {
//	for {
//
//		_, p, err := wrap.Conn.ReadMessage()
//		if err != nil {
//			GlobalNodes.ClearOfflineNode(wrap.NodeName)
//			log.Error(err)
//			return
//		}
//		log.Debugf("read ws msg:%s", string(p))
//		cmd, err := utils.ExtractCmdValue(string(p))
//		if err != nil {
//			log.Error(err)
//			continue
//		}
//		remoteIP := wrap.Conn.RemoteAddr().String()
//		switch cmd {
//		case protocol.CMD_NODE_LOGIN_REQ:
//			req, login, err := handler.Login(p, remoteIP)
//			if err != nil {
//				log.Error(err)
//				continue
//			}
//			wrap.Write(websocket.TextMessage, login)
//			//登录成功之后将信息保存到全局的超级节点map当中
//			wrap.NodeID = req.NodeID
//			wrap.NodeName = req.NodeName
//			GlobalNodes.AddOnlineNodeInfo(req)
//		case protocol.CMD_NODE_HEALTHY_REQ:
//			//todo 使用平台数据库的私钥进行验签,验签不通过就断开连接,节点进行公私钥检查
//			//刷新到期时间
//			wrap.FlushDeadLine()
//
//			//todo 后续业务
//
//		}
//
//	}
//}
