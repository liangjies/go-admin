package middleware

import (
	"go-admin/internal/app/global"
	"go-admin/internal/app/model/common/response"
	"go-admin/internal/app/service"
	"go-admin/internal/app/utils"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, _ := utils.GetClaims(c)
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		var e *casbin.SyncedEnforcer

		if global.SYS_Enforcer != nil {
			e = global.SYS_Enforcer
		}
		/*
			// Redis缓存角色权限
			// 判断Redis是否开启
			if global.SYS_CONFIG.System.UseRedis {
				// 从Redis获取权限
				e = casbinService.CasbinRedis()
			} else {
				// 从数据库获取权限
				e = casbinService.Casbin()
			}
		*/
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.SYS_CONFIG.System.Env == "develop" || success {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
