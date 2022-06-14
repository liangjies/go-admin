package system

import (
	"go-admin/internal/app/global"
)

// 用户登录日志
type SysUserLogin struct {
	global.SYS_MODEL
	Username     string `json:"userName" gorm:"comment:用户登录名"`                                               // 用户登录名
	Ip           string `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`                                  // 请求ip
	Status       int    `json:"status" form:"status" gorm:"column:status;comment:登录状态"`                      // 登录状态
	Agent        string `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`                           // 代理
	ErrorMessage string `json:"error_message" form:"error_message" gorm:"column:error_message;comment:错误信息"` // 错误信息
}
