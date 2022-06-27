package system

import (
	"bytes"
	"encoding/json"
	"go-admin/internal/app/global"
	"go-admin/internal/app/model/common/request"
	"go-admin/internal/app/model/system"
	systemReq "go-admin/internal/app/model/system/request"
	systemTask "go-admin/internal/app/task/system"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type SysJobsService struct {
}

var sysJobsService SysJobsService

// CreateSysJobs 创建SysJobs记录
func (sysJobsService *SysJobsService) CreateSysJobs(sysJobs system.SysJob) (err error) {
	err = global.SYS_DB.Create(&sysJobs).Error
	if err == nil {
		if sysJobs.Status == 1 {
			// 初始化定时任务
			InitOneTimer(sysJobs)
		}
	}
	return err
}

// DeleteSysJobs 删除SysJobs记录
func (sysJobsService *SysJobsService) DeleteSysJobs(sysJob system.SysJob) (err error) {
	err = global.SYS_DB.Delete(&sysJob).Error
	if err == nil {
		// 清空定时任务
		global.SYS_Timer.Clear(strconv.FormatUint(uint64(sysJob.ID), 10))
	}
	return err
}

// DeleteSysJobsByIds 批量删除SysJobs记录
func (sysJobsService *SysJobsService) DeleteSysJobsByIds(ids request.IdsReq) (err error) {
	var sysJobs []system.SysJob
	err = global.SYS_DB.Delete(&[]system.SysJob{}, "id in ?", ids.Ids).Error
	if err == nil {
		global.SYS_DB.Unscoped().Where("id in ?", ids.Ids).Find(&sysJobs)
		// 清空定时任务
		for _, sysJob := range sysJobs {
			global.SYS_Timer.Clear(strconv.FormatUint(uint64(sysJob.ID), 10))
		}
	}
	return err
}

// UpdateSysJobs 更新SysJobs记录
func (sysJobsService *SysJobsService) UpdateSysJobs(sysJobs system.SysJob) (err error) {
	err = global.SYS_DB.Save(&sysJobs).Error
	if err == nil {
		// 清空定时任务
		global.SYS_Timer.Clear(strconv.FormatUint(uint64(sysJobs.ID), 10))
		if sysJobs.Status == 1 {
			// 更新定时任务
			InitOneTimer(sysJobs)
		}
	}
	return err
}

// 系统 更新SysJobs记录
func (sysJobsService *SysJobsService) SysUpdateSysJobs(sysJobs system.SysJob) (err error) {
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

// 获取状态为正常的所有定时任务
func (sysJobsService *SysJobsService) GetSysJobsList() (err error, sysJobs []system.SysJob) {
	err = global.SYS_DB.Where("status = ?", 1).Find(&sysJobs).Error
	return
}

func (sysJobsService *SysJobsService) InitTimer() {
	if global.SYS_CONFIG.Timer.Start {
		// 获取所有定时任务
		err, sysJobs := sysJobsService.GetSysJobsList()
		if err != nil {
			global.SYS_LOG.Error("获取定时任务失败!", zap.Error(err))
			return
		}
		_ = sysJobs

		// 定义定时任务
		for _, sysJob := range sysJobs {
			InitOneTimer(sysJob)
		}
	}
}

// 初始化一个定时任务
func InitOneTimer(sysJob system.SysJob) {
	// 判断任务类型
	switch sysJob.JobType {
	case 1:
		go func(sysJob system.SysJob) {
			EntryId, err := global.SYS_Timer.AddTaskByFunc(strconv.FormatUint(uint64(sysJob.ID), 10), sysJob.CronExpression, func() {
				ExecuteRESTful(sysJob)
			})
			// 运行错误更改任务状态
			if err != nil {
				sysJob.Status = 3
			}
			sysJob.EntryId = int(EntryId)
			sysJobsService.SysUpdateSysJobs(sysJob)
		}(sysJob)
	case 2:
		go func(sysJob system.SysJob) {
			EntryId, err := global.SYS_Timer.AddTaskByFunc(strconv.FormatUint(uint64(sysJob.ID), 10), sysJob.CronExpression, func() {
				ExecuteMethod(sysJob)
			})
			// 运行错误更改任务状态
			if err != nil {
				sysJob.Status = 3
			}
			sysJob.EntryId = int(EntryId)
			sysJobsService.SysUpdateSysJobs(sysJob)
		}(sysJob)
	default:
		global.SYS_LOG.Error("定时任务类型错误，任务名称：" + sysJob.JobName)
	}
}

// 执行运行方法
func ExecuteMethod(sysJob system.SysJob) {
	var sysJobLogsService SysJobLogsService
	var sysJobLog system.SysJobLog
	// 开始时间
	startTime := time.Now()
	var count = 0
	var err error
	var res string
	// 发送请求
	res, err = systemTask.ScheduleTaskRun(sysJob.InvokeTarget)
	// 结束时间
	endTime := time.Now()
	// 执行时间
	sysJobLog.Latency = endTime.Sub(startTime)
	if err != nil {
		sysJobLog.JobMessage = err.Error()
		sysJobLog.Status = 2
	} else {
		count++
		sysJobLog.JobMessage = res
		sysJobLog.Status = 1
	}
	// 保存日志
	sysJobLog.JobId = sysJob.ID
	sysJobLog.JobName = sysJob.JobName
	sysJobLog.JobType = sysJob.JobType
	sysJobLog.InvokeTarget = sysJob.InvokeTarget
	sysJobLog.CreateTime = time.Now()
	_ = sysJobLogsService.CreateSysJobLogs(sysJobLog)
}

// 执行RESTful请求
func ExecuteRESTful(sysJob system.SysJob) {
	var sysJobLogsService SysJobLogsService
	var sysJobLog system.SysJobLog
	// 开始时间
	startTime := time.Now()
	var count = 0
	var err error
	var str string
	// 发送请求
	str, err = Get(sysJob.InvokeTarget)
	// 结束时间
	endTime := time.Now()
	// 执行时间
	sysJobLog.Latency = endTime.Sub(startTime)
	if err != nil {
		sysJobLog.JobMessage = err.Error()
		sysJobLog.Status = 2
	} else {
		count++
		sysJobLog.JobMessage = str
		sysJobLog.Status = 1
	}
	// 保存日志
	sysJobLog.JobId = sysJob.ID
	sysJobLog.JobName = sysJob.JobName
	sysJobLog.JobType = sysJob.JobType
	sysJobLog.InvokeTarget = sysJob.InvokeTarget
	sysJobLog.CreateTime = time.Now()
	_ = sysJobLogsService.CreateSysJobLogs(sysJobLog)
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)

	return string(result), nil
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) ([]byte, error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return result, nil

}
