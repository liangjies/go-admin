package global

import (
	"go-admin/internal/app/config"
	"go-admin/internal/app/utils/timer"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	SYS_DB     *gorm.DB
	SYS_REDIS  *redis.Client
	SYS_CONFIG config.Server
	SYS_VIP    *viper.Viper
	SYS_LOG    *zap.Logger
	SYS_Timer  timer.Timer = timer.NewTimerTask()
	SYS_IPQuery 
	BlackCache local_cache.Cache
)
