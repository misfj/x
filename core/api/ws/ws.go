package ws

import (
	"coredx/core/api/ws/handler"
	"coredx/core/node"
	"coredx/log"
	"coredx/protocol"
	"coredx/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

func NewWrapWsConn(conn *websocket.Conn, interval int) *WrapWsConn {
	return &WrapWsConn{
		conn:     conn,
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
	log.Debugf("当前时间:%v", time.Now())
	log.Debugf("ws conn flush dead line:%v", time.Now().Add(time.Duration(w.internal)*time.Second))
	sub := time.Now().Add(time.Duration(w.internal) * time.Second).Sub(time.Now())
	log.Debugf("当前时间Sub:%v", sub)
	err := w.conn.SetReadDeadline(time.Now().Add(time.Duration(w.internal) * time.Second))
	if err != nil {
		log.Errorf("超级大错误,请联系超级节点负责人:%v", err)
		return
	}
}

type WrapWsConn struct {
	//gorilla websocket 包非并发安全,所以加锁
	conn     *websocket.Conn
	mux      sync.Mutex
	nodeID   string
	nodeName string
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
	err := w.conn.WriteMessage(typ, p)
	if err != nil {
		log.Error(err)
	}
}

// NodeService func Establish(w http.ResponseWriter, r *http.Request) {
func NodeService(ctx *gin.Context) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许所有CORS请求，实际使用时应该更严格
		},
		ReadBufferSize:   4096,
		WriteBufferSize:  4096,
		HandshakeTimeout: 30 * time.Second,
	}
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Error(err)
		return
	}
	wrapConn := NewWrapWsConn(c, 20)
	wrapConn.FlushDeadLine()
	handleNode(wrapConn)
}
func handleNode(wrap *WrapWsConn) {
	for {

		_, p, err := wrap.conn.ReadMessage()
		if err != nil {
			node.GlobalNodes.ClearOfflineNode(wrap.nodeName)
			log.Error(err)
			return
		}
		wrap.FlushDeadLine()
		log.Debugf("ws server read  msg:%s", string(p))
		cmd, err := utils.ExtractCmdValue(string(p))
		if err != nil {
			log.Error(err)
			continue
		}
		remoteIP := wrap.conn.RemoteAddr().String()
		switch cmd {
		case protocol.CMD_NODE_LOGIN_REQ:
			req, login, err := handler.Login(p, remoteIP)
			if err != nil {
				log.Error(err)
				continue
			}
			wrap.Write(websocket.TextMessage, login)
			//登录成功之后将信息保存到全局的超级节点map当中
			wrap.nodeID = req.NodeID
			wrap.nodeName = req.NodeName

			node.GlobalNodes.AddOnlineNodeInfo(req)
		case protocol.CMD_NODE_HEALTHY_REQ:
			//todo 使用平台数据库的私钥进行验签,验签不通过就断开连接,节点进行公私钥检查
			//刷新到期时间
			wrap.FlushDeadLine()
		case protocol.CMD_NODE_DIRTY_REQ:
			dirty, _ := handler.Dirty(p)
			wrap.Write(websocket.TextMessage, dirty)
			wrap.FlushDeadLine()
			//todo 后续业务

		}

	}
}
