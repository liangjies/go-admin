package router

import (
	"go-admin/internal/app/router/example"
	"go-admin/internal/app/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
