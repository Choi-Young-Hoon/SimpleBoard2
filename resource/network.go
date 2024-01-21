package resource

import (
	"errors"
	psnet "github.com/shirou/gopsutil/net"
	"time"
)

type NetworkInterfaceInfo struct {
	Name  string   `json:"name"`
	MTU   int      `json:"mtu"`
	Addrs []string `json:"addrs"`
}

func NewNetworkInfo() *NetworkInfo {
	netInterfaces, err := psnet.Interfaces()
	if err != nil {
		return nil
	}

	networkInfo := &NetworkInfo{}
	for _, netInterface := range netInterfaces {
		netInfo := NetworkInterfaceInfo{
			Name: netInterface.Name,
			MTU:  netInterface.MTU,
		}

		for _, ipAddr := range netInterface.Addrs {
			netInfo.Addrs = append(netInfo.Addrs, ipAddr.Addr)
		}

		networkInfo.NetworkInterfaceInfos = append(networkInfo.NetworkInterfaceInfos, netInfo)
	}

	return networkInfo
}

type NetworkInfo struct {
	NetworkInterfaceInfos []NetworkInterfaceInfo `json:"network_interface_infos"`

	// Byte
	BytesSent int `json:"bytes_sent"`
	BytesRecv int `json:"bytes_recv"`

	PacketsSent int `json:"packets_sent"`
	PacketsRecv int `json:"packets_recv"`
}

func (n *NetworkInfo) GetNetworkInfo() error {
	ioCounters1, err := psnet.IOCounters(false)
	if err != nil {

		return err
	}

	time.Sleep(time.Second)

	ioCounters2, err := psnet.IOCounters(false)
	if err != nil {
		return err
	}

	if len(ioCounters1) < 1 || len(ioCounters2) < 1 {
		return errors.New("ioCounters is empty")
	}

	n.BytesSent = int(ioCounters2[0].BytesSent-ioCounters1[0].BytesSent) / 1024
	n.BytesRecv = int(ioCounters2[0].BytesRecv-ioCounters1[0].BytesRecv) / 1024
	n.PacketsSent = int(ioCounters2[0].PacketsSent - ioCounters1[0].PacketsSent)
	n.PacketsRecv = int(ioCounters2[0].PacketsRecv - ioCounters1[0].PacketsRecv)

	return nil
}

func (n *NetworkInfo) Dump() {
	for _, info := range n.NetworkInterfaceInfos {
		println("Name:", info.Name)
		println("MTU:", info.MTU)
		for _, addr := range info.Addrs {
			println("Addr:", addr)
		}
	}

	println("BytesSent:", n.BytesSent)
	println("BytesRecv:", n.BytesRecv)
	println("PacketsSent:", n.PacketsSent)
	println("PacketsRecv:", n.PacketsRecv)
}
