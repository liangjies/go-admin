package system

import (
	"go-admin/internal/app/global"
	"go-admin/internal/app/model/common/request"
	"go-admin/internal/app/model/system"
	systemReq "go-admin/internal/app/model/system/request"
)

type SysJobsService struct {
}

// CreateSysJobs 创建SysJobs记录
func (sysJobsService *SysJobsService) CreateSysJobs(sysJobs system.SysJob) (err error) {
	err = global.SYS_DB.Create(&sysJobs).Error
	return err
}

// DeleteSysJobs 删除SysJobs记录
func (sysJobsService *SysJobsService) DeleteSysJobs(sysJobs system.SysJob) (err error) {
	err = global.SYS_DB.Delete(&sysJobs).Error
	return err
}

// DeleteSysJobsByIds 批量删除SysJobs记录
func (sysJobsService *SysJobsService) DeleteSysJobsByIds(ids request.IdsReq) (err error) {
	err = global.SYS_DB.Delete(&[]system.SysJob{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSysJobs 更新SysJobs记录
func (sysJobsService *SysJobsService) UpdateSysJobs(sysJobs system.SysJob) (err error) {
	err = global.SYS_DB.Save(&sysJobs).Error
	return err
}

// GetSysJobs 根据id获取SysJobs记录
// Author [liangjies]
func (sysJobsService *SysJobsService) GetSysJobs(id uint) (err error, sysJobs system.SysJob) {
	err = global.SYS_DB.Where("id = ?", id).First(&sysJobs).Error
	return
}

// GetSysJobsInfoList 分页获取SysJobs记录
func (sysJobsService *SysJobsService) GetSysJobsInfoList(info systemReq.SearchJobParams) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.SYS_DB.Model(&system.SysJob{})
	var sysJobss []system.SysJob
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.JobName != "" {
		db = db.Where("Job_name LIKE ?", "%"+info.JobName+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&sysJobss).Error
	return err, sysJobss, total
}
