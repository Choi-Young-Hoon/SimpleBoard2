package resource

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type CPUInfo struct {
	Cores     int32   `json:"cores"`
	Ghz       float64 `json:"ghz"`
	ModelName string  `json:"model_name"`
	VendorID  string  `json:"vendor_id"`
}

func NewCPUInfos() *CPUInfos {
	cpuInterfaceInfos, err := cpu.Info()
	if err != nil {
		return nil
	}

	cpuInfos := &CPUInfos{}
	for _, cpuInterfaceInfo := range cpuInterfaceInfos {
		cpuInfo := CPUInfo{
			Cores:     cpuInterfaceInfo.Cores,
			Ghz:       cpuInterfaceInfo.Mhz / 1024,
			ModelName: cpuInterfaceInfo.ModelName,
			VendorID:  cpuInterfaceInfo.VendorID,
		}
		cpuInfos.Infos = append(cpuInfos.Infos, cpuInfo)
	}

	return cpuInfos
}

type CPUInfos struct {
	Infos []CPUInfo `json:"cpu_infos"`

	UsagePercent int `json:"usage_percent"`
}

func (c *CPUInfos) GetCPUUsage() error {

	cpuUsage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return err
	}

	c.UsagePercent = int(cpuUsage[0])

	return nil
}

func (c *CPUInfos) Dump() {
	for _, info := range c.Infos {
		println("CPU Cores:", info.Cores)
		println("CPU Ghz:", info.Ghz)
		println("CPU Model name:", info.ModelName)
		println("CPU Vendor ID:", info.VendorID)
	}
	println("CPU Usage:", c.UsagePercent)
}
