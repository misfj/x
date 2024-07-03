package node

import (
	"coredx/core/api/ws/pkg"
	"coredx/log"
	"sync"
)

var GlobalNodes Nodes

type Nodes struct {
	Nodes map[string]*pkg.Node
	mux   sync.Mutex
}

func NewGlobalNodes() {
	GlobalNodes.Nodes = make(map[string]*pkg.Node)
}

// AddOnlineNodeInfo 添加登录上来的超级节点信息
func (gln *Nodes) AddOnlineNodeInfo(node *pkg.Node) {

	defer gln.mux.Unlock()
	gln.mux.Lock()
	//便利当前map,如果存在节点名称但是添加进来,说明系统出现问题
	nodeName := node.NodeName
	if gln.Nodes[nodeName] != nil {
		log.Error("系统出现严重问题,节点名称重复")
		return
	}
	gln.Nodes[nodeName] = node
	log.Debugf("添加节点信息:%v", node)
}

// ClearOfflineNode 清楚离线超级节点信息
func (gln *Nodes) ClearOfflineNode(nodeName string) {
	if gln.Nodes[nodeName] == nil {
		log.Errorf("%s:系统出现严重问题,节点不存在\n", nodeName)
		return
	}
	defer gln.mux.Unlock()
	gln.mux.Lock()
	log.Debugf("删除节点信息:%v", gln.Nodes[nodeName])

	delete(gln.Nodes, nodeName)
}
