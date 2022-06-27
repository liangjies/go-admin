package system

import (
	v1 "go-admin/internal/app/api/v1"
	"go-admin/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

type SysJobsRouter struct{}

// InitSysJobsRouter 初始化 SysJobs 路由信息
func (s *SysJobsRouter) InitSysJobsRouter(Router *gin.RouterGroup) {
	sysJobsRouter := Router.Group("sysJobs").Use(middleware.OperationRecord())
	sysJobsRouterWithoutRecord := Router.Group("sysJobs")
	var sysJobsApi = v1.ApiGroupApp.SystemApiGroup.SysJobsApi
	{
		sysJobsRouter.POST("createSysJobs", sysJobsApi.CreateSysJobs)             // 新建SysJobs
		sysJobsRouter.DELETE("deleteSysJobs", sysJobsApi.DeleteSysJobs)           // 删除SysJobs
		sysJobsRouter.DELETE("deleteSysJobsByIds", sysJobsApi.DeleteSysJobsByIds) // 批量删除SysJobs
		sysJobsRouter.PUT("updateSysJobs", sysJobsApi.UpdateSysJobs)              // 更新SysJobs
		sysJobsRouter.POST("runSysJob", sysJobsApi.RunSysJob)                     // 运行定时任务
	}
	{
		sysJobsRouterWithoutRecord.GET("findSysJobs", sysJobsApi.FindSysJobs)       // 根据ID获取SysJobs
		sysJobsRouterWithoutRecord.GET("getSysJobsList", sysJobsApi.GetSysJobsList) // 获取SysJobs列表
	}
}
