package middleware

import (
	"bytes"
	"go-admin/internal/app/global"
	"go-admin/internal/app/model/system"
	systemReq "go-admin/internal/app/model/system/request"
	"go-admin/internal/app/service"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

var sysUserLoginService = service.ServiceGroupApp.SystemServiceGroup.SysUserLoginService

func LoginLogRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var l systemReq.Login
		_ = c.ShouldBindJSON(&l)
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		c.Next()
		// 登录信息
		var sysUserLogin system.SysUserLogin
		sysUserLogin.Username = l.Username
		sysUserLogin.IP = c.ClientIP()
		sysUserLogin.Agent = c.Request.UserAgent()
		ua := user_agent.New(c.Request.UserAgent())
		sysUserLogin.LoginLocation, _, _ = global.SYS_IPQuery.QueryIP(c.ClientIP())
		sysUserLogin.OS = ua.OS()
		sysUserLogin.Browser, _ = ua.Browser()
		sysUserLogin.ErrorMessage = writer.body.String()
		sysUserLoginService.CreateSysUserLogin(sysUserLogin)

	}
}
