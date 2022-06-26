package initialize

import (
	"bytes"
	"encoding/json"
	"go-admin/internal/app/global"
	"go-admin/internal/app/model/system"
	"go-admin/internal/app/service"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"
)

var sysJobsService = service.ServiceGroupApp.SystemServiceGroup.SysJobsService
var sysJobLogsService = service.ServiceGroupApp.SystemServiceGroup.SysJobLogsService

//var timerdWork = timer.TimerdWork

func Timer() {
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
			// 判断任务类型
			switch sysJob.JobType {
			case 1:
				go func(sysJob system.SysJob) {
					EntryId, err := global.SYS_Timer.AddTaskByFunc(sysJob.JobName, sysJob.CronExpression, func() {
						ExecuteRESTful(sysJob)
					})
					// 运行错误更改任务状态
					if err != nil {
						sysJob.Status = 3
					}
					sysJob.EntryId = int(EntryId)
					sysJobsService.UpdateSysJobs(sysJob)
				}(sysJob)
			case 2:
				// TODO: 定时任务
			default:
				global.SYS_LOG.Error("定时任务类型错误，任务名称：" + sysJob.JobName)
			}

			/*
				for i := range global.SYS_CONFIG.Timer.Detail {
					go func(detail config.Detail) {
						global.SYS_Timer.AddTaskByFunc("ClearDB", global.SYS_CONFIG.Timer.Spec, func() {
							err := utils.ClearTable(global.SYS_DB, detail.TableName, detail.CompareField, detail.Interval)
							if err != nil {
								fmt.Println("timer error:", err)
							}
						})
					}(global.SYS_CONFIG.Timer.Detail[i])
				}
			*/
		}
	}
}

// 执行RESTful请求
func ExecuteRESTful(sysJob system.SysJob) {
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
