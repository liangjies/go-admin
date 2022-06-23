package system

import (
	"context"
	"errors"
	"go-admin/internal/app/global"
	"go-admin/internal/app/model/system/response"
	"strings"
)

type SysMonitorService struct{}

func (sysMonitorService *SysMonitorService) GetRedisInfoCur() (res response.RedisInfoCur, err error) {
	if global.SYS_CONFIG.System.UseRedis {
		client := global.SYS_REDIS
		ctx := context.Background()
		// Key数量
		res.DBSize = client.DBSize(ctx).String()[strings.Index(client.DBSize(ctx).String(), ":")+2:]
		// 内存信息
		res.UsedMemory = parseInfo(client.Info(ctx, "Memory").Val(), "used_memory")
	} else {
		err = errors.New("redis未配置")
	}
	return
}

func (sysMonitorService *SysMonitorService) GetRedisInfo() (res response.RedisInfor, err error) {

	if !global.SYS_CONFIG.System.UseRedis {
		return res, errors.New("redis未配置")
	}
	if global.SYS_REDIS == nil {
		return res, errors.New("redis连接失败")
	}

	client := global.SYS_REDIS
	ctx := context.Background()
	// 获取 Redis 状态
	server := client.Info(ctx, "Server").Val()
	clients := client.Info(ctx, "Clients").Val()
	stats := client.Info(ctx, "Stats").Val()
	menory := client.Info(ctx, "Memory").Val()
	// Redis版本
	res.Redis_version = parseInfo(server, "redis_version")
	// Redis模式
	res.Redis_mode = parseInfo(server, "redis_mode")
	//  Redis TCP端口
	res.Tcp_port = parseInfo(server, "tcp_port")
	// Redis运行天数
	res.Uptime_in_days = parseInfo(server, "uptime_in_days")
	// Key数量
	res.DBSize = client.DBSize(ctx).String()[strings.Index(client.DBSize(ctx).String(), ":")+2:]
	// 内存信息
	res.UsedMemory = parseInfo(menory, "used_memory")
	// 当前连接数
	res.Connected_clients = parseInfo(clients, "connected_clients")
	// 已执行命令数
	res.Total_commands_processed = parseInfo(stats, "total_commands_processed")
	return
}

// 解析redis信息
func parseInfo(info string, key string) (res string) {
	leftIndex := strings.Index(info, key)
	midIndex := strings.Index(info[leftIndex:], ":")
	rightIndex := strings.Index(info[leftIndex:], "\r\n")

	return info[leftIndex+midIndex+1 : leftIndex+rightIndex]
}
