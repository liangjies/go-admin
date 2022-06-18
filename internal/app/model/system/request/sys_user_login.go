package request

import (
	"go-admin/internal/app/model/common/request"
	"go-admin/internal/app/model/system"
)

type SysUserLoginSearch struct {
	DateRange []string `json:"dateRange[]" form:"dateRange[]"` //日期范围
	system.SysUserLogin
	request.PageInfo
}
