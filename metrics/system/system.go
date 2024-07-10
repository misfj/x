package system

import (
	"coredx/log"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"time"
)

func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetDiskPercent() float64 {
	var totalSize, usedSize uint64 = 0, 0

	parts, err := disk.Partitions(true)
	if err != nil {
		log.Error("Error getting disk partitions:", err)
		return 0
	}

	for _, part := range parts {
		diskInfo, err := disk.Usage(part.Mountpoint)
		if err != nil {
			log.Error("Error getting usage for %s: %v\n", part.Mountpoint, err)
			continue
		}

		totalSize += diskInfo.Total
		usedSize += diskInfo.Used
	}

	if totalSize == 0 {
		return 0
	}

	return (float64(usedSize) / float64(totalSize)) * 100
}
