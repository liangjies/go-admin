package request

import (
	"go-admin/internal/app/model/{{.Package}}"
	"go-admin/internal/app/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
