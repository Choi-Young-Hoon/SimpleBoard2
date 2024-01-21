package resource

import (
	"encoding/json"
	"github.com/shirou/gopsutil/mem"
	"math"
)

func NewMemoryInfo() *MemoryInfo {
	return &MemoryInfo{}
}

type MemoryInfo struct {
	// Gbyte
	Total float64 `json:"total"`
	Used  float64 `json:"used"`
	Free  float64 `json:"free"`

	UsedPercent int `json:"used_percent"`
}

func (m *MemoryInfo) GetMemoryInfo() error {
	memUsage, err := mem.VirtualMemory()
	if err != nil {
		return err
	}

	m.Total = float64(memUsage.Total) / (1024 * 1024 * 1024)
	m.Used = float64(memUsage.Used) / (1024 * 1024 * 1024)
	m.Free = float64(memUsage.Free) / (1024 * 1024 * 1024)

	m.Total = math.Round(m.Total*10) / 10
	m.Used = math.Round(m.Used*10) / 10
	m.Free = math.Round(m.Free*10) / 10

	m.UsedPercent = int(memUsage.UsedPercent)

	return nil
}

func (m *MemoryInfo) Dump() {
	jsonString, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		println(err)
		return
	}

	println(string(jsonString))
}
