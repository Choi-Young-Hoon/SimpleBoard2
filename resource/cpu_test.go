package resource

import "testing"

func TestCPUInfos(test *testing.T) {
	cpuInfos := NewCPUInfos()

	cpuInfos.GetCPUUsage()
	cpuInfos.Dump()
}
