package system

import (
	v1 "go-admin/internal/app/api/v1"
	"go-admin/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

type SysJobLogsRouter struct {
}

// InitSysJobLogsRouter 初始化 SysJobLogs 路由信息
func (s *SysJobLogsRouter) InitSysJobLogsRouter(Router *gin.RouterGroup) {
	sysJobLogsRouter := Router.Group("sysJobLogs").Use(middleware.OperationRecord())
	sysJobLogsRouterWithoutRecord := Router.Group("sysJobLogs")
	var sysJobLogsApi = v1.ApiGroupApp.SystemApiGroup.SysJobLogsApi
	{
		sysJobLogsRouter.POST("createSysJobLogs", sysJobLogsApi.CreateSysJobLogs)             // 新建SysJobLogs
		sysJobLogsRouter.DELETE("deleteSysJobLogs", sysJobLogsApi.DeleteSysJobLogs)           // 删除SysJobLogs
		sysJobLogsRouter.DELETE("deleteSysJobLogsByIds", sysJobLogsApi.DeleteSysJobLogsByIds) // 批量删除SysJobLogs
		sysJobLogsRouter.DELETE("clearSysJobLogs", sysJobLogsApi.ClearSysJobLogs)             // 清理运行日志
	}
	{
		sysJobLogsRouterWithoutRecord.GET("findSysJobLogs", sysJobLogsApi.FindSysJobLogs)       // 根据ID获取SysJobLogs
		sysJobLogsRouterWithoutRecord.GET("getSysJobLogsList", sysJobLogsApi.GetSysJobLogsList) // 获取SysJobLogs列表
	}
}
