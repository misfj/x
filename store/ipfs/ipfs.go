package ipfs

import (
	"coredx/config"
	"coredx/log"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
	"time"
)

var ipfs *IPFS

type IPFS struct {
	listenAddr string
	httpClient *resty.Client
	peers      []string
	//peerMap    map[string]PeerInfo
	//pwd        string
}

func Init(conf *config.BackUp) error {
	client := resty.New()
	client.SetTimeout(time.Second * 60 * 3)
	ipfs = &IPFS{
		listenAddr: conf.Addr,
		httpClient: client,
		peers:      make([]string, 0),
	}
	err := ipfs.findPeers()
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debug("initial backup store success")
	return nil
}

func (i *IPFS) findPeers() error {
	api := "/api/v0/swarm/peers"
	api = fmt.Sprintf("%s%s", i.listenAddr, api)
	resp, err := i.httpClient.R().Post(api)
	if err != nil {
		return err
	}
	type ResStruct struct {
		Peers []struct {
			Addr     string `json:"Addr"`
			Peer     string `json:"Peer"`
			Identify struct {
				ID           string `json:"ID"`
				PublicKey    string `json:"PublicKey"`
				Addresses    any    `json:"Addresses"`
				AgentVersion string `json:"AgentVersion"`
				Protocols    any    `json:"Protocols"`
			} `json:"Identify"`
		} `json:"Peers"`
	}
	var peers ResStruct
	err = json.Unmarshal(resp.Body(), &peers)
	if err != nil {
		return err
	}
	for _, v := range peers.Peers {
		split := strings.Split(v.Addr, "/")
		x := fmt.Sprintf("http://%s%s", split[2], ":5001")
		i.peers = append(i.peers, x)
	}
	i.peers = append(i.peers, i.listenAddr)
	log.Debugf("ipfs node:%v", i.peers)

	return nil
}
