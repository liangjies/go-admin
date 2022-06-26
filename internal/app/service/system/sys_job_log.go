package system

import (
	"go-admin/internal/app/global"
	"go-admin/internal/app/model/common/request"
	"go-admin/internal/app/model/system"
	systemReq "go-admin/internal/app/model/system/request"
	"time"
)

type SysJobLogsService struct {
}

// CreateSysJobLogs 创建SysJobLogs记录
func (sysJobLogsService *SysJobLogsService) CreateSysJobLogs(sysJobLogs system.SysJobLog) (err error) {
	err = global.SYS_DB.Create(&sysJobLogs).Error
	return err
}

// DeleteSysJobLogs 删除SysJobLogs记录
func (sysJobLogsService *SysJobLogsService) DeleteSysJobLogs(sysJobLogs system.SysJobLog) (err error) {
	err = global.SYS_DB.Delete(&sysJobLogs).Error
	return err
}

// DeleteSysJobLogsByIds 批量删除SysJobLogs记录
func (sysJobLogsService *SysJobLogsService) DeleteSysJobLogsByIds(ids request.IdsReq) (err error) {
	err = global.SYS_DB.Delete(&[]system.SysJobLog{}, "id in ?", ids.Ids).Error
	return err
}

// GetSysJobLogs 根据id获取SysJobLogs记录
func (sysJobLogsService *SysJobLogsService) GetSysJobLogs(id uint) (err error, sysJobLogs system.SysJobLog) {
	err = global.SYS_DB.Where("id = ?", id).First(&sysJobLogs).Error
	return
}

// GetSysJobLogsInfoList 分页获取SysJobLogs记录
func (sysJobLogsService *SysJobLogsService) GetSysJobLogsInfoList(info systemReq.SearchJobLogParams) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.SYS_DB.Model(&system.SysJobLog{})
	var sysJobLogss []system.SysJobLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.JobName != "" {
		db = db.Where("job_name LIKE ?", "%"+info.JobName+"%")
	}
	if info.JobId != 0 {
		db = db.Where("Job_id = ?", info.JobId)
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&sysJobLogss).Error
	return err, sysJobLogss, total
}

// 清空日志
func (sysJobLogsService *SysJobLogsService) ClearSysJobLogs(info systemReq.CleanJobLogParams) (err error, total int64) {
	// 创建db
	db := global.SYS_DB.Model(&system.SysJobLog{})
	var sysJobLogss []system.SysJobLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.JobName != "" {
		db = db.Where("job_name LIKE ?", "%"+info.JobName+"%")
	}
	if info.JobId != 0 {
		db = db.Where("Job_id = ?", info.JobId)
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	nowTime := time.Now()
	if info.Time == "week" {
		// 清理一周前的数据
		time := nowTime.AddDate(0, 0, -7)
		db = db.Where("create_time < ?", time)
	} else if info.Time == "month" {
		// 清理一月前的数据
		time := nowTime.AddDate(0, -1, 0)
		db = db.Where("create_time < ?", time)
	}

	res := db.Delete(&sysJobLogss)
	return res.Error, res.RowsAffected
}
