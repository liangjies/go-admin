package system

import (
	"go-admin/internal/app/global"
	"go-admin/internal/app/model/common/request"
	"go-admin/internal/app/model/common/response"
	"go-admin/internal/app/model/system"
	systemReq "go-admin/internal/app/model/system/request"
	"go-admin/internal/app/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysJobLogsApi struct{}

var sysJobLogsService = service.ServiceGroupApp.SystemServiceGroup.SysJobLogsService

// CreateSysJobLogs 创建SysJobLogs
// @Tags SysJobLogs
// @Summary 创建SysJobLogs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysJobLog true "创建SysJobLogs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysJobLogs/createSysJobLogs [post]
func (sysJobLogsApi *SysJobLogsApi) CreateSysJobLogs(c *gin.Context) {
	var sysJobLogs system.SysJobLog
	_ = c.ShouldBindJSON(&sysJobLogs)
	if err := sysJobLogsService.CreateSysJobLogs(sysJobLogs); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysJobLogs 删除SysJobLogs
// @Tags SysJobLogs
// @Summary 删除SysJobLogs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysJobLog true "删除SysJobLogs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysJobLogs/deleteSysJobLogs [delete]
func (sysJobLogsApi *SysJobLogsApi) DeleteSysJobLogs(c *gin.Context) {
	var sysJobLogs system.SysJobLog
	_ = c.ShouldBindJSON(&sysJobLogs)
	if err := sysJobLogsService.DeleteSysJobLogs(sysJobLogs); err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysJobLogsByIds 批量删除SysJobLogs
// @Tags SysJobLogs
// @Summary 批量删除SysJobLogs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysJobLogs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysJobLogs/deleteSysJobLogsByIds [delete]
func (sysJobLogsApi *SysJobLogsApi) DeleteSysJobLogsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sysJobLogsService.DeleteSysJobLogsByIds(IDS); err != nil {
		global.SYS_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// FindSysJobLogs 用id查询SysJobLogs
// @Tags SysJobLogs
// @Summary 用id查询SysJobLogs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysJobLog true "用id查询SysJobLogs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysJobLogs/findSysJobLogs [get]
func (sysJobLogsApi *SysJobLogsApi) FindSysJobLogs(c *gin.Context) {
	var sysJobLogs system.SysJobLog
	_ = c.ShouldBindQuery(&sysJobLogs)
	if err, resysJobLogs := sysJobLogsService.GetSysJobLogs(sysJobLogs.ID); err != nil {
		global.SYS_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysJobLogs": resysJobLogs}, c)
	}
}

// GetSysJobLogsList 分页获取SysJobLogs列表
// @Tags SysJobLogs
// @Summary 分页获取SysJobLogs列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query systemReq.SysJobLogsSearch true "分页获取SysJobLogs列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysJobLogs/getSysJobLogsList [get]
func (sysJobLogsApi *SysJobLogsApi) GetSysJobLogsList(c *gin.Context) {
	var pageInfo systemReq.SearchJobLogParams
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sysJobLogsService.GetSysJobLogsInfoList(pageInfo); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags SysJobLogs
// @Summary 清空日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysJobLog true "清空日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"清空成功"}"
// @Router /sysJobLogs/clearSysJobLogs [delete]
func (sysJobLogsApi *SysJobLogsApi) ClearSysJobLogs(c *gin.Context) {
	var cleanJobLogParams systemReq.CleanJobLogParams
	_ = c.ShouldBindJSON(&cleanJobLogParams)
	if err, count := sysJobLogsService.ClearSysJobLogs(cleanJobLogParams); err != nil {
		global.SYS_LOG.Error("清空失败!", zap.Error(err))
		response.FailWithMessage("清空失败", c)
	} else {
		response.OkWithMessage("清理成功,共清理"+strconv.Itoa(int(count))+"行", c)
	}
}
