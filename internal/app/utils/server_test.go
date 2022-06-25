package utils

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/v3/cpu"
)

func TestServer(t *testing.T) {
	// version, _ := host.KernelVersion()
	// fmt.Println(version)

	// platform, family, version, _ := host.PlatformInformation()
	// fmt.Println("platform:", platform) // 操作系统信息
	// fmt.Println("family:", family)
	// fmt.Println("version:", version)

	// hInfo, _ := host.Info()
	// fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)

	// hnet := GetOutboundIP()
	// fmt.Println("net info:%v ", hnet)
	// info, _ := mem.VirtualMemory()
	// fmt.Println(info)
	// info2, _ := mem.SwapMemory()
	// fmt.Println(info2)

	// CPU使用率
	info4, _ := cpu.Percent(time.Duration(time.Second), false)
	fmt.Println(info4)

	info3, _ := load.Avg()
	fmt.Println(info3)
	/*
		for _, info := range infos {
			fmt.Println(info.ModelName)
				data, _ := json.MarshalIndent(info, "", " ")
				fmt.Print(string(data))
		}
	*/
}

func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

// Get preferred outbound ip of this machine
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	return localAddr.IP.String()
}
