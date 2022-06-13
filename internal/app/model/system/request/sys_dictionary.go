package request

import (
	"go-admin/internal/app/model/common/request"
	"go-admin/internal/app/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
