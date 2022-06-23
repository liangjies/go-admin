package system

import (
	v1 "go-admin/internal/app/api/v1"
	"go-admin/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system")
	sysRouterRecore := Router.Group("system").Use(middleware.OperationRecord())
	systemApi := v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		sysRouter.POST("getServerInfo", systemApi.GetServerInfo)           // 获取服务器信息
		sysRouterRecore.POST("getSystemConfig", systemApi.GetSystemConfig) // 获取配置文件内容
		sysRouterRecore.POST("setSystemConfig", systemApi.SetSystemConfig) // 设置配置文件内容
		sysRouterRecore.POST("reloadSystem", systemApi.ReloadSystem)       // 重启服务
	}

}
