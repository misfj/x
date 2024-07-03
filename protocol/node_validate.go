package protocol

import (
	"errors"
	"fmt"
	"math/rand"
)

func (n *NodeLoginRequest) Validate() error {
	if n.Command == "" {
		return errors.New("command field is null")
	}
	if n.NodeManager == "" {
		n.NodeManager = "dev_manager"
	}
	if n.NodeDept == "" {
		n.NodeDept = "dev_dept"
	}
	if n.NodePhoneType == "" {
		n.NodePhoneType = "+86"
	}
	if n.NodePhone == "" {
		n.NodePhone = "188888888888"
	}
	if n.NodeAddr == "" {
		n.NodeAddr = "成大花园"
	}
	if n.NodeOper == "" {
		n.NodeOper = "移动运营商"
	}
	if n.LoginCode == "" {
		n.LoginCode = "123456"
	}
	if n.NodeID == "" {
		int31n := rand.Int31n(9999)
		n.NodeID = fmt.Sprintf("%s:%d", "nodeID", int31n)
	}
	if n.NodeName == "" {
		int31n := rand.Int31n(9999)
		n.NodeName = fmt.Sprintf("%s:%d", "nodeName", int31n)
	}
	if n.Public == "" {
		int31n := rand.Int31n(9999)
		n.Public = fmt.Sprintf("%s:%d", "public", int31n)
	}
	if n.PublicMD5 == "" {
		int31n := rand.Int31n(9999)
		n.PublicMD5 = fmt.Sprintf("%s:%d", "publicMD5", int31n)
	}
	return nil
}
