package wsnode

import (
	"context"
	"coredx/config"
	"coredx/core/wsnode/handle"
	"coredx/log"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WsNode struct {
	// //节点ID
	// NodeID string
	// //节点名称
	// NodeName string
	// //节点管理员
	// NodeManager string
	// //节点部门信息
	// NodeDept string
	// //节点联系人电话前缀
	// NodeManagerPhoneType string
	// //节点联系人电话号码
	// NodeManagerPhone string
	// //节点地址
	// NodeAddr string
	// //节点运营商
	// NodeOper string
	// //节点登陆码
	// NodeLoginCode string
	// //节点启动时间
	// NodeStart time.Time
	// //启动断开时间
	// NodeExit time.Time
	// //断开原因
	// NodeExitReason string
	// //节点公钥
	// NodePubkey string
	// //节点公钥MD5
	// NodePubKeyMD5 string
	// //保活时间                 (默认单位是分钟)
	// NodeHealthyInterval int
	server *http.Server
	ctx    context.Context
}

func NewWsNode() *WsNode {
	return &WsNode{}
}
func (wsNode *WsNode) Init() error {
	return nil
}

func (wsNode *WsNode) Name() string {
	return "WsNode Server"
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type handler struct {
}

// ServeHTTP 实现 http.Handler 接口
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/v1/ws/1":
		h.Plat_handler(w, r)
	default:
		http.NotFound(w, r)
	}
}
func (s *handler) Plat_handler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("upgrade error:", err)
		return
	}
	defer c.Close()
	handle.Plat_Handler(c)
}

func (wsNode *WsNode) Startup(ctx context.Context) error {
	wsNode.ctx = ctx
	nodePlatform := ctx.Value("node-platform")
	ws := ctx.Value("ws")
	nodePlatformImpl := nodePlatform.(string)
	wsImpl := ws.(string)
	log.Info(config.GetNode().NodePrivateMd5)
	/*
		平台:node-platform == "", ws !=""
		超级节点:node-platform != "", ws ==""
	*/

	fmt.Println(nodePlatformImpl, wsImpl)
	if nodePlatformImpl == "" && wsImpl != "" {

		fmt.Println("平台")
		wsHandler := &handler{}
		wsNode.server = &http.Server{
			Addr:           wsImpl,
			Handler:        wsHandler,
			ReadTimeout:    10 * time.Minute,
			WriteTimeout:   10 * time.Minute,
			MaxHeaderBytes: 1 << 20,
		}
		//平台 开启websocket服务
		log.Infof("websocket service run:%s\n", wsImpl)
		log.Info(wsNode.server.ListenAndServe())
	}
	if nodePlatformImpl != "" && wsImpl == "" {
		//节点 主动连接平台
		fmt.Println("节点")
		timeoutCtx, _ := context.WithTimeout(context.Background(), time.Second*30)
		dial := websocket.Dialer{}
		conn, _, err := dial.DialContext(timeoutCtx, nodePlatformImpl, nil)
		if err != nil {
			log.Error(err)
			return err
		}
		handle.Node_handler(conn)

	}
	return nil

}

func (wsNode *WsNode) Close() error {
	if wsNode.server == nil {
		return nil
	}
	wsNode.server.Shutdown(wsNode.ctx)
	return nil
}
