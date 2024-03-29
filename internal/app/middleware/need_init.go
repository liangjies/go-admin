package middleware

import (
	"go-admin/internal/app/global"
	"go-admin/internal/app/model/common/response"

	"github.com/gin-gonic/gin"
)

// 处理跨域请求,支持options访问
func NeedInit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.SYS_DB == nil {
			response.OkWithDetailed(gin.H{
				"needInit": true,
			}, "前往初始化数据库", c)
			c.Abort()
		} else {
			c.Next()
		}
		// 处理请求
	}
}
