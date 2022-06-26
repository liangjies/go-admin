package system

import (
	"time"
)

type SysJobLog struct {
	ID           uint          `gorm:"primarykey"` // 主键ID
	JobId        uint          `gorm:"column:job_id;comment:名称" json:"jobId" form:"jobId"`
	JobName      string        `gorm:"column:job_name;comment:任务名称" json:"jobName" form:"jobName"`
	JobType      int           `gorm:"column:job_type;comment:任务类型" json:"jobType" form:"jobType"`
	InvokeTarget string        `gorm:"column:invoke_target;comment:调用目标字符串" json:"invokeTarget" form:"invokeTarget"`
	JobMessage   string        `gorm:"column:job_message;comment:日志信息" json:"jobMessage" form:"jobMessage"`
	Status       int           `gorm:"column:status;comment:执行状态（0正常 1失败）" json:"status" form:"status"`
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:执行时间" swaggertype:"string"`
	CreateTime   time.Time     `gorm:"column:create_time;comment:创建时间" json:"createTime" form:"createTime"`
}
