package initialize

import (
	"go-admin/internal/app/service"
)

var sysJobsService = service.ServiceGroupApp.SystemServiceGroup.SysJobsService

func Timer() {
	sysJobsService.InitTimer()
}
