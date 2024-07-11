package process

import (
	"coredx/log"
	"github.com/shirou/gopsutil/v4/process"
	"strconv"
)

type Process struct {
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

type NetConn struct {
	PName      string `json:"pName"`
	Pid        int32  `json:"pid"`
	ConnType   string `json:"connType"`
	LocalAddr  string `json:"localAddr"`
	RemoteAddr string `json:"RemoteAddr"`
	Status     string `json:"status"`
}

func GetProcessInfo() []*Process {
	pids, err := process.Pids()
	if err != nil {
		log.Error(err)
	}

	var processes []*Process

	for _, pid := range pids {
		newProcess, err := process.NewProcess(pid)
		if err != nil {
			log.Error(err)
			continue
		}

		name, err := newProcess.Name()
		if err != nil {
			log.Error(err)
			continue
		}

		s, err := newProcess.IsRunning()
		if err != nil {
			log.Error(err)
			continue
		}

		status := "running"
		if !s {
			status = "not running"
		}

		ppid, err := newProcess.Ppid()
		if err != nil {
			log.Error(err)
		}

		userName, err := newProcess.Username()
		if err != nil {
			log.Error(err)
		}

		createTime, err := newProcess.CreateTime()
		if err != nil {
			log.Error(err)
		}
		numThreads, err := newProcess.NumThreads()
		if err != nil {
			log.Error(err)
		}

		conns, err := newProcess.Connections()
		if err != nil {
			log.Error(err)
		}

		cpuUsage, err := newProcess.CPUPercent()
		if err != nil {
			log.Error("Failed to get CPU usage:", err)
		}

		memPercent, err := newProcess.MemoryPercent()
		if err != nil {
			log.Error("Failed to get memory percent:", err)
		}

		memInfo, err := newProcess.MemoryInfo()
		if err != nil {
			log.Error("Failed to get memory used:", err)
		}

		p := &Process{
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

		processes = append(processes, p)
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

func GetProcessNetInfo() []*NetConn {
	pids, err := process.Pids()
	if err != nil {
		log.Error(err)
	}

	var netConns []*NetConn

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
			connType := "unknown"
			if conn.Type == 1 {
				connType = "TCP"
			} else if conn.Type == 2 {
				connType = "UDP"
			}

			netConn := &NetConn{
				PName:      name,
				Pid:        pid,
				ConnType:   connType,
				LocalAddr:  conn.Laddr.IP + ":" + strconv.Itoa(int(conn.Laddr.Port)),
				RemoteAddr: conn.Raddr.IP + ":" + strconv.Itoa(int(conn.Raddr.Port)),
				Status:     conn.Status,
			}
			netConns = append(netConns, netConn)
		}
	}
	return netConns
}
