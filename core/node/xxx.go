package node

import (
	"coredx/log"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

func NewWrapWsConn(conn *websocket.Conn, interval int) *WrapWsConn {
	return &WrapWsConn{
		Conn:              conn,
		internal:          interval,
		LatestHealthyTime: time.Now().Add(time.Duration(interval) * time.Second),
	}
}
func (w *WrapWsConn) FlushDeadLine() {
	defer func() {
		if e := recover(); e != nil {
			log.Errorf("超级大错误,请联系超级节点负责人:%v", e)
			return
		}
	}()
	log.Debugf("客户端剩余时间:%v", w.LatestHealthyTime.Sub(time.Now()))
	if w.LatestHealthyTime.Sub(time.Now()) < time.Second*20 {
		log.Debug("保活时间小于20s")
		newt := w.LatestHealthyTime.Add(time.Duration(w.internal) * time.Second)
		err := w.Conn.SetReadDeadline(w.LatestHealthyTime.Add(time.Duration(w.internal) * time.Second))
		if err != nil {
			log.Error(err)
			return
		}
		w.LatestHealthyTime = newt
		log.Debugf("客户端增加之后的时间:%v", newt)
	} else {

	}
	//err := w.Conn.SetReadDeadline(time.Now().Add(time.Duration(w.internal) * time.Second))
	//if err != nil {
	//	log.Errorf("超级大错误,请联系超级节点负责人:%v", err)
	//	return
	//}
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
