package v1

import (
	"go-admin/internal/app/api/v1/example"
	"go-admin/internal/app/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
