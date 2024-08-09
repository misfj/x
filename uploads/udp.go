package global

import (
	"coredx/log"
	"net"
)

func SendUDPMessage(remoteAddr string, msg []byte) error {
	//targetAddr := "localhost:9090"
	// 创建一个 UDP 地址
	udpAddr, err := net.ResolveUDPAddr("udp", remoteAddr)
	if err != nil {
		log.Error(err)
		return err
	}
	// 创建一个 UDP 连接（socket）
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Error(err)
		return err
	}
	defer conn.Close()
	// 发送消息
	_, err = conn.Write(msg)
	if err != nil {
		log.Error(err)
		return err
	}
	buf := make([]byte, 4096)

	n, err := conn.Read(buf)
	if err != nil {
		log.Error(err)
		return err
	}
	Global.UdpCh <- buf[:n]
	return nil
}
