package request

import (
	"go-admin/internal/app/model/common/request"
	"go-admin/internal/app/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
