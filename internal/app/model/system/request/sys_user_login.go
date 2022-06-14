package request

import (
	"go-admin/internal/app/model/common/request"
	"go-admin/internal/app/model/system"
)

type SysUserLoginSearch struct {
	system.SysUserLogin
	request.PageInfo
}
