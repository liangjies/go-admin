package system

import (
	"go-admin/internal/app/global"
	"go-admin/internal/app/model/common/request"
	"go-admin/internal/app/model/common/response"
	"go-admin/internal/app/model/system"
	systemReq "go-admin/internal/app/model/system/request"
	"go-admin/internal/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysUserLoginApi struct{}

// @Tags SysUserLogin
// @Summary 创建SysUserLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysUserLogin true "创建SysUserLogin"
// @Success 200 {object} response.Response{msg=string} "创建SysUserLogin"
// @Router /sysUserLogin/createSysUserLogin [post]
func (s *SysUserLoginApi) CreateSysUserLogin(c *gin.Context) {
	var SysUserLogin system.SysUserLogin
	_ = c.ShouldBindJSON(&SysUserLogin)
	if err := sysUserLoginService.CreateSysUserLogin(SysUserLogin); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysUserLogin
// @Summary 删除SysUserLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysUserLogin true "SysUserLogin模型"
// @Success 200 {object} response.Response{msg=string} "删除SysUserLogin"
// @Router /sysUserLogin/deleteSysUserLogin [delete]
func (s *SysUserLoginApi) DeleteSysUserLogin(c *gin.Context) {
	var SysUserLogin system.SysUserLogin
	_ = c.ShouldBindJSON(&SysUserLogin)
	if err := sysUserLoginService.DeleteSysUserLogin(SysUserLogin); err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysUserLogin
// @Summary 批量删除SysUserLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysUserLogin"
// @Success 200 {object} response.Response{msg=string} "批量删除SysUserLogin"
// @Router /sysUserLogin/deleteSysUserLoginByIds [delete]
func (s *SysUserLoginApi) DeleteSysUserLoginByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sysUserLoginService.DeleteSysUserLoginByIds(IDS); err != nil {
		global.SYS_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags SysUserLogin
// @Summary 用id查询SysUserLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysUserLogin true "Id"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "用id查询SysUserLogin"
// @Router /sysUserLogin/findSysUserLogin [get]
func (s *SysUserLoginApi) FindSysUserLogin(c *gin.Context) {
	var SysUserLogin system.SysUserLogin
	_ = c.ShouldBindQuery(&SysUserLogin)
	if err := utils.Verify(SysUserLogin, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, reSysUserLogin := sysUserLoginService.GetSysUserLogin(SysUserLogin.ID); err != nil {
		global.SYS_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"reSysUserLogin": reSysUserLogin}, "查询成功", c)
	}
}

// @Tags SysUserLogin
// @Summary 分页获取SysUserLogin列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysUserLoginSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取SysUserLogin列表,返回包括列表,总数,页码,每页数量"
// @Router /sysUserLogin/getSysUserLoginList [get]
func (s *SysUserLoginApi) GetSysUserLoginList(c *gin.Context) {
	var pageInfo systemReq.SysUserLoginSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sysUserLoginService.GetSysUserLoginInfoList(pageInfo); err != nil {
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
