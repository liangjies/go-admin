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

// 解析redis信息
func parseInfo(info string, key string) (res string) {
	leftIndex := strings.Index(info, key)
	midIndex := strings.Index(info[leftIndex:], ":")
	rightIndex := strings.Index(info[leftIndex:], "\r\n")

	return info[leftIndex+midIndex+1 : leftIndex+rightIndex]
}
