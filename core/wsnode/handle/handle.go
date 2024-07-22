package handle

import (
	"github.com/gorilla/websocket"
)

type NodeCommand string

const (
	login NodeCommand = "login"
)

func Node_handler(conn *websocket.Conn) {
	node_handler(conn)
}
