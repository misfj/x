package handler

import (
	"coredx/core/api/ws/pkg"
	//"coredx/core/node"
	"coredx/db"
	v1 "coredx/db/model/v1"
	"coredx/log"
	"coredx/protocol"
	"coredx/utils"
	"encoding/json"
	"net/http"
	"time"
)

func Login(msg []byte, remoteIP string) (*pkg.Node, []byte, error) {
	var req protocol.NodeLoginRequest
	err := json.Unmarshal(msg, &req)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	err = req.Validate()
	if err != nil {
	}
	log.Debug(utils.TimeParse())
	//插入数据到node_regist表,然后插入数据到node_info
	err = v1.NodeRegistCreate(db.GDB, time.Now(), time.Now(), "0", "1", "正常", req.NodeID, req.NodeName, remoteIP, time.Now(), utils.TimeParse())
	if err != nil {
		return nil, nil, err
	}
	err = v1.NodeInfoCreate(db.GDB, time.Now(), time.Now(), "0", "正常", "1",
		req.NodeManager, req.NodeDept, req.NodePhoneType, req.NodePhone, req.NodeAddr,
		req.NodeOper, req.LoginCode, req.NodeID, req.NodeName, req.Public, "", req.PublicMD5, "")
	if err != nil {
		return nil, nil, err
	}
	return pkg.Convert(&req), protocol.NodeResponse("login", http.StatusOK, "success", ""), nil
}
