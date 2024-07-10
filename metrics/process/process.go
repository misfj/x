package process

import (
	"coredx/log"
	"github.com/shirou/gopsutil/v4/process"
)

type Process struct {
	Name       string  `json:"name"`
	Pid        int32   `json:"pid"`
	Status     string  `json:"status"`
	CpuPercent float64 `json:"cpuPercent"`
	MemPercent float32 `json:"memPercent"`
}

func GetProcess() []*Process {
	pids, err := process.Pids()
	if err != nil {
		log.Error(err)
	}

	var processes []*Process

	for _, pid := range pids {
		newProcess, err := process.NewProcess(pid)
		if err != nil {
			log.Error("Failed to create process:", err)
			continue
		}

		name, err := newProcess.Name()
		if err != nil {
			log.Error("Failed to get process name:", err)
			continue
		}

		s, err := newProcess.IsRunning()
		if err != nil {
			log.Error("Failed to get process status:", err)
			continue
		}

		status := "running"
		if !s {
			status = "not running"
		}

		cpuUsage, err := newProcess.CPUPercent()
		if err != nil {
			log.Error("Failed to get CPU usage:", err)
		}

		memPercent, err := newProcess.MemoryPercent()
		if err != nil {
			log.Error("Failed to get memory percent:", err)
		}

		p := &Process{
			Name:       name,
			Pid:        pid,
			Status:     status,
			CpuPercent: cpuUsage,
			MemPercent: memPercent,
		}

		processes = append(processes, p)
	}
	return processes
}
