package handler

import (
	"coredx/protocol"
	"encoding/json"
)

func Dirty(msg []byte) ([]byte, error) {
	return json.Marshal(protocol.NodeDirtyRequest{Command: protocol.CMD_NODE_DIRTY_RES})

}
