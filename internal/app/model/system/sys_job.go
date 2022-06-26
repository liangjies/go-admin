package system

import (
	"go-admin/internal/app/global"
)

type SysJob struct {
	global.SYS_MODEL
	JobName        string `json:"jobName" form:"jobName" gorm:"size:255;comment:名称"`                    // 名称
	JobGroup       string `json:"jobGroup" form:"jobGroup" gorm:"size:255;comment:任务分组"`                // 任务分组
	JobType        int    `json:"jobType" form:"jobType" gorm:"size:1;comment:任务类型"`                    // 任务类型
	CronExpression string `json:"cronExpression" form:"cronExpression" gorm:"size:255;comment:cron表达式"` // cron表达式
	InvokeTarget   string `json:"invokeTarget" form:"invokeTarget" gorm:"size:255;comment:调用目标"`        // 调用目标
	Args           string `json:"args" form:"args" gorm:"size:255;comment:目标参数"`                        // 目标参数
	Concurrent     int    `json:"concurrent" form:"concurrent" gorm:"size:1;comment:是否并发"`              // 是否并发
	Status         int    `json:"status" form:"status" gorm:"size:1;comment:状态"`                        // 状态
	Description    string `json:"description" form:"description" gorm:"size:255;comment:描述"`            // 描述
	EntryId        int    `json:"entryId" form:"entryId" gorm:"size:11;comment:job启动时返回的id"`            // job启动时返回的id
}
