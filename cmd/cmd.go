package cmd

import (
	"go-admin/internal/app/core"
	"go-admin/internal/app/global"
	"go-admin/internal/app/initialize"
	"go-admin/internal/app/service"
	"time"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

func Execute() {
	// 初始化Viper
	global.SYS_VIP = core.Viper()
	// 初始化zap日志库
	global.SYS_LOG = core.Zap()
	// 注册全局logger
	zap.ReplaceGlobals(global.SYS_LOG)
	// 初始化GORM连接
	global.SYS_DB = initialize.Gorm()
	// 初始化定时器
	initialize.Timer()
	// 记录开始时间
	global.SYS_StartTime = time.Now()
	// 初始化数据库
	if global.SYS_DB != nil {
		initialize.RegisterTables(global.SYS_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.SYS_DB.DB()
		defer db.Close()
	}
	// 初始化LocalCache
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(global.SYS_CONFIG.JWT.ExpiresTime)),
	)
	// 初始化IP data
	global.SYS_IPQuery = initialize.QQWRT()
	// 初始化casbin
	casbinService.CasbinInit()
	/*
		// 初始化casbin redis
		if err := casbinService.CasbinInitRedis(); err != nil {
			global.SYS_LOG.Error("casbin redis failed", zap.Error(err))
			os.Exit(0)
		}
	*/

	core.RunWindowsServer()

}
