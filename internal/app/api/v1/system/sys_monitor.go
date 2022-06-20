package system

import (
	"go-admin/internal/app/global"
	"go-admin/internal/app/model/common/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysMonitorApi struct{}

// @Tags GetRedisInfoCur
// @Summary 获取Redis实时信息
// @Security ApiKeyAuth
// @Produce application/json
// @Param
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取Redis实时信息"
// @Router /sysMonitor/getRedisInfoCur [get]
func (s *SysMonitorApi) GetRedisInfoCur(c *gin.Context) {
	if reRedisInfoCur, err := sysMonitorService.GetRedisInfoCur(); err != nil {
		global.SYS_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(reRedisInfoCur, "查询成功", c)
	}
}

// @Tags GetRedisInfo
// @Summary 获取Redis信息
// @Security ApiKeyAuth
// @Produce application/json
// @Param
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取Redis信息"
// @Router /sysMonitor/getRedisInfo [get]
func (s *SysMonitorApi) GetRedisInfo(c *gin.Context) {
	if reRedisInfo, err := sysMonitorService.GetRedisInfo(); err != nil {
		global.SYS_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(reRedisInfo, "查询成功", c)
	}
}
