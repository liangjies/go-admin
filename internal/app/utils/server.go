package utils

import (
	"fmt"
	"go-admin/internal/app/global"
	"runtime"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type Server struct {
	Os   Os   `json:"os"`
	Cpu  Cpu  `json:"cpu"`
	Ram  Ram  `json:"ram"`
	Disk Disk `json:"disk"`
}

type Os struct {
	SysComputerName string `json:"sysComputerName"`
	SysOsName       string `json:"sysOsName"`
	SysOsArch       string `json:"sysOsArch"`
	NumCPU          int    `json:"numCpu"`
	Compiler        string `json:"compiler"`
	GoVersion       string `json:"goVersion"`
	NumGoroutine    int    `json:"numGoroutine"`
	RunTime         uint   `json:"runTime"`
}

type Cpu struct {
	CpuPercent float64 `json:"cpuPercent"`
	ModelName  string  `json:"modelName"`
	Cores      int     `json:"cores"`
}

type Ram struct {
	UsedMB      int     `json:"usedMb"`
	TotalMB     int     `json:"totalMb"`
	UsedPercent int     `json:"usedPercent"`
	SwapUsedMB  int     `json:"swapUsedMb"`
	SwapTotalMB int     `json:"swapTotalMb"`
	GoUsed      float64 `json:"goUsed"`
}

type Disk struct {
	UsedMB      int              `json:"usedMb"`
	UsedGB      int              `json:"usedGb"`
	TotalMB     int              `json:"totalMb"`
	TotalGB     int              `json:"totalGb"`
	UsedPercent int              `json:"usedPercent"`
	DiskList    []disk.UsageStat `json:"diskList"`
}

//@function: InitCPU
//@description: OS信息
//@return: o Os, err error

func InitOS() (o Os) {
	sysInfo, err := host.Info()
	if err == nil {
		// 系统名称
		o.SysComputerName = sysInfo.Hostname
		// 操作系统
		o.SysOsName = sysInfo.OS
		// 系统架构
		o.SysOsArch = sysInfo.KernelArch
	}
	// 逻辑内核数
	o.NumCPU = runtime.NumCPU()
	o.Compiler = runtime.Compiler
	o.GoVersion = runtime.Version()
	o.NumGoroutine = runtime.NumGoroutine()
	// 程序运行时间
	// 计算开始和现在时间差
	o.RunTime = uint(time.Since(global.SYS_StartTime).Minutes())
	return o
}

//@function: InitCPU
//@description: CPU信息
//@return: c Cpu, err error

func InitCPU() (c Cpu, err error) {
	// 物理内核数
	if cores, err := cpu.Counts(false); err != nil {
		return c, err
	} else {
		c.Cores = cores
	}
	// CPU型号
	if infos, err := cpu.Info(); err != nil {
		return c, err
	} else {
		c.ModelName = infos[0].ModelName
	}
	// CPU使用率
	if cpuPercent, err := cpu.Percent(time.Second, false); err != nil {
		return c, err
	} else {
		c.CpuPercent, _ = strconv.ParseFloat(strconv.FormatFloat(cpuPercent[0], 'f', 2, 64), 64)
	}
	return c, nil
}

//@function: InitRAM
//@description: RAM信息
//@return: r Ram, err error

func InitRAM() (r Ram, err error) {
	// 物理内存
	if u, err := mem.VirtualMemory(); err != nil {
		return r, err
	} else {
		// 使用
		r.UsedMB = int(u.Used) / MB
		// 总量
		r.TotalMB = int(u.Total) / MB
		// 使用率
		r.UsedPercent = int(u.UsedPercent)

	}
	// 虚拟内存
	if u, err := mem.SwapMemory(); err != nil {
		return r, err
	} else {
		// 使用
		r.SwapUsedMB = int(u.Used) / MB
		// 总量
		r.SwapTotalMB = int(u.Total) / MB
	}

	// Golang占用内存
	var gomem runtime.MemStats
	runtime.ReadMemStats(&gomem)
	r.GoUsed = float64(gomem.Sys) / 1024 / 1024

	return r, nil
}

//@function: InitDisk
//@description: 硬盘信息
//@return: d Disk, err error

func InitDisk() (d Disk, err error) {
	/*
		if u, err := disk.Usage("/"); err != nil {
			return d, err
		} else {
			d.UsedMB = int(u.Used) / MB
			d.UsedGB = int(u.Used) / GB
			d.TotalMB = int(u.Total) / MB
			d.TotalGB = int(u.Total) / GB
			d.UsedPercent = int(u.UsedPercent)
		}
	*/
	diskInfo, err := disk.Partitions(true) //所有分区
	if err == nil {
		for _, p := range diskInfo {
			diskDetail, err := disk.Usage(p.Mountpoint)
			if err == nil {
				diskDetail.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", diskDetail.UsedPercent), 64)
				d.DiskList = append(d.DiskList, *diskDetail)
			}
		}
	}

	return d, nil
}
