package protocol

import (
	"encoding/json"
)

type NodeLoginRequest struct {
	Command       string `json:"command"`
	NodeManager   string `json:"nodeManager"`
	NodeDept      string `json:"nodeDept"`
	NodePhoneType string `json:"nodePhoneType"`
	NodePhone     string `json:"nodePhone"`
	NodeAddr      string `json:"nodeAddr"`
	NodeOper      string `json:"nodeOper"`
	LoginCode     string `json:"loginCode"`
	NodeID        string `json:"nodeID"`
	NodeName      string `json:"nodeName"`
	Public        string `json:"public"`
	PublicMD5     string `json:"publicMD5"`
}
type NodeLoginResponse struct {
	Command string `json:"command"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    string `json:"data"`
}

func NodeResponse(typ string, code int, msg string, data string) []byte {
	switch typ {
	case "login":
		var res NodeLoginResponse
		res.Command = CMD_NODE_LOGIN_RES
		res.Code = code
		res.Msg = msg
		res.Data = data
		marshal, _ := json.Marshal(res)
		return marshal
	}
	return nil
}

type NodeHealthyRequest struct {
	Command  string `json:"command"`
	NodeName string `json:"nodeName"`
	Sign     string `json:"sign"`
}
type NodeDirtyRequest struct {
	Command string `json:"command"`
}
