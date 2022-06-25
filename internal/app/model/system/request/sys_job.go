package request

import (
	"go-admin/internal/app/model/common/request"
	"go-admin/internal/app/model/system"
)

// api分页条件查询及排序结构体
type SearchJobParams struct {
	system.SysJob
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
