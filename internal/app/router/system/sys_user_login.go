package system

import (
	v1 "go-admin/internal/app/api/v1"

	"github.com/gin-gonic/gin"
)

func (s *OperationRecordRouter) InitSysUserLoginRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysUserLogin")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.SysUserLoginApi
	{
		operationRecordRouter.POST("createSysUserLogin", authorityMenuApi.CreateSysUserLogin)             // 新建SysUserLogin
		operationRecordRouter.DELETE("deleteSysUserLogin", authorityMenuApi.DeleteSysUserLogin)           // 删除SysUserLogin
		operationRecordRouter.DELETE("deleteSysUserLoginByIds", authorityMenuApi.DeleteSysUserLoginByIds) // 批量删除SysUserLogin
		operationRecordRouter.GET("findSysUserLogin", authorityMenuApi.FindSysUserLogin)                  // 根据ID获取SysUserLogin
		operationRecordRouter.GET("getSysUserLoginList", authorityMenuApi.GetSysUserLoginList)            // 获取SysUserLogin列表
	}
}
