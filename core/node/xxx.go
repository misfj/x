package node

import (
	"coredx/log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func NewWrapWsConn(conn *websocket.Conn, interval int) *WrapWsConn {
	return &WrapWsConn{

		Conn:     conn,
		internal: interval,
		//LatestHealthyTime: time.Now().Add(time.Duration(interval) * time.Second),
	}
}

type WrapWsConn struct {
	//gorilla websocket 包非并发安全,所以加锁
	Conn     *websocket.Conn
	mux      sync.Mutex
	NodeID   string
	NodeName string
	//保活时间

	internal          int
	LatestHealthyTime time.Time
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
