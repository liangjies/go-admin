package system

import (
	"time"
)

type SysJobLog struct {
	ID            uint      `gorm:"primarykey"` // 主键ID
	JobId         int64     `gorm:"column:job_log_id;comment:名称" json:"job_log_id" form:"job_log_id"`
	JobName       string    `gorm:"column:job_name;comment:任务名称" json:"job_name" form:"job_name"`
	JobGroup      string    `gorm:"column:job_group;comment:任务组名" json:"job_group" form:"job_group"`
	InvokeTarget  string    `gorm:"column:invoke_target;comment:调用目标字符串" json:"invoke_target" form:"invoke_target"`
	JobMessage    string    `gorm:"column:job_message;comment:日志信息" json:"job_message" form:"job_message"`
	Status        string    `gorm:"column:status;comment:执行状态（0正常 1失败）" json:"status" form:"status"`
	ExceptionInfo string    `gorm:"column:exception_info;comment:异常信息" json:"exception_info" form:"exception_info"`
	CreateTime    time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time" form:"create_time"`
}
