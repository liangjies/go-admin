package request

import (
	"go-admin/internal/app/model/common/request"
	"go-admin/internal/app/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
