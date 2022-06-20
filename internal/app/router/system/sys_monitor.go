package system

import (
	v1 "go-admin/internal/app/api/v1"

	"github.com/gin-gonic/gin"
)

type SysMonitorRouter struct{}

func (s *SysMonitorRouter) InitSysMonitorRouter(Router *gin.RouterGroup) {
	sysMonitorRouter := Router.Group("sysMonitor")
	sysMonitorApi := v1.ApiGroupApp.SystemApiGroup.SysMonitorApi
	{
		sysMonitorRouter.GET("getRedisInfoCur", sysMonitorApi.GetRedisInfoCur) // 获取实时redis信息
		sysMonitorRouter.GET("getRedisInfo", sysMonitorApi.GetRedisInfo)       // 获取redis信息

	}
}
