package resource

import "github.com/shirou/gopsutil/disk"

type DiskInfo struct {
	Device     string `json:"device"`
	MountPoint string `json:"mount_point"`

	// Gbyte
	Total int `json:"total"`
	Used  int `json:"used"`
	Free  int `json:"free"`

	UsagePercent int `json:"usage_percent"`
}

type DiskInfos struct {
	Infos []DiskInfo `json:"disk_infos"`
}

func NewDiskInfos() *DiskInfos {
	return &DiskInfos{}
}

func (d *DiskInfos) GetDiskInfos() error {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return err
	}

	for _, partition := range partitions {
		diskInfo := DiskInfo{
			Device:     partition.Device,
			MountPoint: partition.Mountpoint,
		}

		diskUsageStat, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			return err
		}

		diskInfo.Total = int(diskUsageStat.Total / 1024 / 1024 / 1024)
		diskInfo.Used = int(diskUsageStat.Used / 1024 / 1024 / 1024)
		diskInfo.Free = int(diskUsageStat.Free / 1024 / 1024 / 1024)
		diskInfo.UsagePercent = int(diskUsageStat.UsedPercent)

		d.Infos = append(d.Infos, diskInfo)
	}

	return nil
}

func (d *DiskInfos) Dump() {
	for _, info := range d.Infos {
		println("Device:", info.Device)
		println("MountPoint:", info.MountPoint)
		println("Total:", info.Total, "Gb")
		println("Used:", info.Used, "Gb")
		println("Free:", info.Free, "Gb")
		println("UsagePercent:", info.UsagePercent, "%")
	}
}
