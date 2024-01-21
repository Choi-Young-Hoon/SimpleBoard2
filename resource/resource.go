package resource

import "sync"

func NewSystemResource() *SystemResource {
	return &SystemResource{
		CPU:     NewCPUInfos(),
		Disk:    NewDiskInfos(),
		Memory:  NewMemoryInfo(),
		Network: NewNetworkInfo(),
	}
}

type SystemResource struct {
	CPU     *CPUInfos    `json:"cpu_infos"`
	Disk    *DiskInfos   `json:"disk_infos"`
	Memory  *MemoryInfo  `json:"memory_info"`
	Network *NetworkInfo `json:"network_info"`
}

func (s *SystemResource) GetSystemResource() {
	wg := &sync.WaitGroup{}
	wg.Add(4)
	s.getCPUUsage(wg)
	s.getDiskInfos(wg)
	s.getMemoryInfo(wg)
	s.getNetworkInfo(wg)
	wg.Wait()
}

func (s *SystemResource) Dump() {
	s.CPU.Dump()
	s.Disk.Dump()
	s.Memory.Dump()
	s.Network.Dump()
}

func (s *SystemResource) getCPUUsage(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		s.CPU.GetCPUUsage()
	}()
}

func (s *SystemResource) getDiskInfos(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		s.Disk.GetDiskInfos()
	}()
}

func (s *SystemResource) getMemoryInfo(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		s.Memory.GetMemoryInfo()
	}()
}

func (s *SystemResource) getNetworkInfo(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		s.Network.GetNetworkInfo()
	}()
}
