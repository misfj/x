package process

import (
	"coredx/log"
	"fmt"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

type ConnType int

const (
	TCP ConnType = iota + 1
	UDP
)

func (ct ConnType) String() string {
	switch ct {
	case TCP:
		return "TCP"
	case UDP:
		return "UDP"
	default:
		return "unknown"
	}
}

type ProInfo struct {
	PName      string  `json:"pName"`
	Pid        int32   `json:"pid"`
	Ppid       int32   `json:"ppid"`
	UserName   string  `json:"userName"`
	CreateTime int64   `json:"createTime"`
	NumThreads int32   `json:"numThreads"`
	NumConns   int     `json:"numConns"`
	Status     string  `json:"status"`
	CpuPercent float64 `json:"cpuPercent"`
	MemPercent float32 `json:"memPercent"`
	MemUsed    uint64  `json:"memUsed"` //MB为单位
}

type NetConnInfo struct {
	PName      string `json:"pName"`
	Pid        int32  `json:"pid"`
	ConnType   string `json:"connType"`
	LocalAddr  string `json:"localAddr"`
	RemoteAddr string `json:"RemoteAddr"`
	Status     string `json:"status"`
}

func GetProcessInfo() []*ProInfo {
	pids, err := process.Pids()
	if err != nil {
		log.Error("Failed to get PIDs:", err)
		return nil
	}

	var processes []*ProInfo

	for _, pid := range pids {
		p, err := process.NewProcess(pid)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to create process for PID %d:", pid), err)
			continue
		}

		var (
			name, status, userName string
			ppid, numThreads       int32
			createTime             int64
			conns                  []net.ConnectionStat
			cpuUsage               float64
			memPercent             float32
			memInfo                *process.MemoryInfoStat
		)

		if name, err = p.Name(); err != nil {
			log.Error(err)
			continue
		}

		if s, err := p.IsRunning(); err != nil {
			log.Error(err)
			continue
		} else {
			status = "running"
			if !s {
				status = "not running"
			}
		}

		if ppid, err = p.Ppid(); err != nil {
			log.Error(err)
		}

		if userName, err = p.Username(); err != nil {
			log.Error(err)
		}

		if createTime, err = p.CreateTime(); err != nil {
			log.Error(err)
		}

		if numThreads, err = p.NumThreads(); err != nil {
			log.Error(err)
		}

		if conns, err = p.Connections(); err != nil {
			log.Error(err)
		}

		if cpuUsage, err = p.CPUPercent(); err != nil {
			log.Error(err)
		}

		if memPercent, err = p.MemoryPercent(); err != nil {
			log.Error(err)
		}

		if memInfo, err = p.MemoryInfo(); err != nil {
			log.Error(err)
		}

		proInfo := &ProInfo{
			PName:      name,
			Pid:        pid,
			Ppid:       ppid,
			UserName:   userName,
			CreateTime: createTime,
			NumThreads: numThreads,
			NumConns:   len(conns),
			Status:     status,
			CpuPercent: cpuUsage,
			MemPercent: memPercent,
			MemUsed:    memInfo.RSS / 1024 / 1024,
		}

		processes = append(processes, proInfo)
	}

	return processes
}

func KillProcess(pid int32) error {
	pro, err := process.NewProcess(pid)
	if err != nil {
		return err
	}

	return pro.Kill()
}

func GetProcessNetInfo() []*NetConnInfo {
	pids, err := process.Pids()
	if err != nil {
		log.Error(err)
	}

	var netConns []*NetConnInfo

	for _, pid := range pids {
		p, err := process.NewProcess(pid)
		if err != nil {
			log.Error(err)
			continue
		}

		name, err := p.Name()
		if err != nil {
			log.Error(err)
			continue
		}

		conns, err := p.Connections()
		if err != nil {
			log.Error(err)
			continue
		}

		for _, conn := range conns {
			localAddr := formatAddr(conn.Laddr)
			remoteAddr := formatAddr(conn.Raddr)

			netConn := &NetConnInfo{
				PName:      name,
				Pid:        pid,
				ConnType:   ConnType(conn.Type).String(),
				LocalAddr:  localAddr,
				RemoteAddr: remoteAddr,
				Status:     conn.Status,
			}
			netConns = append(netConns, netConn)
		}
	}
	return netConns
}

func formatAddr(addr net.Addr) string {
	if addr.Port == 0 {
		return addr.IP
	}
	return fmt.Sprintf("%s:%d", addr.IP, addr.Port)
}
