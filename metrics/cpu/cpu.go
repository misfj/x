package cpu

import (
	"github.com/shirou/gopsutil/v3/process"
)

func CPUInfo() {
	proc, err := process.NewProcess(int32(1))
}
