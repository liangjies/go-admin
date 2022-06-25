package system

import (
	"go-admin/internal/app/global"
	"go-admin/internal/app/model/common/request"
	"go-admin/internal/app/model/common/response"
	"go-admin/internal/app/model/system"
	systemReq "go-admin/internal/app/model/system/request"
	"go-admin/internal/app/service"
	"go-admin/internal/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysJobsApi struct {
}

var sysJobsService = service.ServiceGroupApp.SystemServiceGroup.SysJobsService

// CreateSysJobs 创建SysJobs
// @Tags SysJobs
// @Summary 创建SysJobs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysJobs true "创建SysJobs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysJobs/createSysJobs [post]
func (sysJobsApi *SysJobsApi) CreateSysJobs(c *gin.Context) {
	var sysJobs system.SysJob
	_ = c.ShouldBindJSON(&sysJobs)
	if err := utils.Verify(sysJobs, utils.JobVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysJobsService.CreateSysJobs(sysJobs); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysJobs 删除SysJobs
// @Tags SysJobs
// @Summary 删除SysJobs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysJob true "删除SysJobs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysJobs/deleteSysJobs [delete]
func (sysJobsApi *SysJobsApi) DeleteSysJobs(c *gin.Context) {
	var sysJobs system.SysJob
	_ = c.ShouldBindJSON(&sysJobs)
	if err := sysJobsService.DeleteSysJobs(sysJobs); err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysJobsByIds 批量删除SysJobs
// @Tags SysJobs
// @Summary 批量删除SysJobs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysJobs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysJobs/deleteSysJobsByIds [delete]
func (sysJobsApi *SysJobsApi) DeleteSysJobsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sysJobsService.DeleteSysJobsByIds(IDS); err != nil {
		global.SYS_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysJobs 更新SysJobs
// @Tags SysJobs
// @Summary 更新SysJobs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysJob true "更新SysJobs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysJobs/updateSysJobs [put]
func (sysJobsApi *SysJobsApi) UpdateSysJobs(c *gin.Context) {
	var sysJobs system.SysJob
	_ = c.ShouldBindJSON(&sysJobs)
	if err := utils.Verify(sysJobs, utils.JobVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysJobsService.UpdateSysJobs(sysJobs); err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysJobs 用id查询SysJobs
// @Tags SysJobs
// @Summary 用id查询SysJobs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysJob true "用id查询SysJobs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysJobs/findSysJobs [get]
func (sysJobsApi *SysJobsApi) FindSysJobs(c *gin.Context) {
	var sysJobs system.SysJob
	_ = c.ShouldBindQuery(&sysJobs)
	if err, resysJobs := sysJobsService.GetSysJobs(sysJobs.ID); err != nil {
		global.SYS_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysJobs": resysJobs}, c)
	}
}

// GetSysJobsList 分页获取SysJobs列表
// @Tags SysJobs
// @Summary 分页获取SysJobs列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query testReq.SysJobsSearch true "分页获取SysJobs列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysJobs/getSysJobsList [get]
func (sysJobsApi *SysJobsApi) GetSysJobsList(c *gin.Context) {
	var pageInfo systemReq.SearchJobParams
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sysJobsService.GetSysJobsInfoList(pageInfo); err != nil {
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
