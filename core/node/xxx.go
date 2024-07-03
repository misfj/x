package node

import (
	"coredx/log"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

func NewWrapWsConn(conn *websocket.Conn, interval int) *WrapWsConn {
	return &WrapWsConn{
		Conn:                  conn,
		internal:              interval,
		LatestHealthyDeadline: time.Now().Add(time.Duration(interval) * time.Second),
	}
}
func (w *WrapWsConn) FlushDeadLine() {
	defer func() {
		if e := recover(); e != nil {
			log.Errorf("超级大错误,请联系超级节点负责人:%v", e)
			return
		}
	}()
	log.Debugf("客户端剩余到期时间:%v", w.LatestHealthyDeadline.Sub(time.Now()))
	if w.LatestHealthyDeadline.Sub(time.Now()) < time.Second*20 {
		log.Debug("这个分支s")
		deadline := w.LatestHealthyDeadline.Add(time.Duration(w.internal) * time.Second)
		err := w.Conn.SetReadDeadline(deadline)
		if err != nil {
			log.Error(err)
			return
		}
		w.LatestHealthyDeadline = deadline
		log.Debugf("客户端剩余到期时间XXXX:%v", w.LatestHealthyDeadline.Sub(time.Now()))
	} else {
		log.Debug("粉煤灰")
	}
}

type WrapWsConn struct {
	//gorilla websocket 包非并发安全,所以加锁
	Conn     *websocket.Conn
	mux      sync.Mutex
	NodeID   string
	NodeName string
	//保活时间
	internal              int
	LatestHealthyDeadline time.Time
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
