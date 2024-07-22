package handle

import (
	"coredx/log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

// Handle 作为超级节点
func node_handler(conn *websocket.Conn) {
	var mutex sync.Mutex

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error(err)
			}
		}()
	loop:
		for {
			defer conn.Close()
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Infof("websocket read message error: %v", err)
				break loop
			}
			log.Debugf("recv:")
			//TODO 解密
			result := gjson.Get(string(msg), "command")
			switch result.Str {
			case string(login):
				response := login_handler(msg)
				mutex.Lock()
				conn.WriteMessage(websocket.TextMessage, response)
				mutex.Unlock()
			}

		}
	}()
}
