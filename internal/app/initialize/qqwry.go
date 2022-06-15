package initialize

import (
	"go-admin/internal/app/global"
	"go-admin/internal/app/utils/qqwry"

	"go.uber.org/zap"
)

func QQWRT() (ipQuery qqwry.IPQuery) {
	if err := ipQuery.LoadFile("assets\\qqwry.dat"); err != nil {
		global.SYS_LOG.Error("load qqwry failed", zap.Error(err))
	}
	return
}
