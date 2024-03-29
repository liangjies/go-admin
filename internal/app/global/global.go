package global

import (
	"go-admin/internal/app/config"
	"go-admin/internal/app/utils/qqwry"
	"go-admin/internal/app/utils/timer"
	"time"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	SYS_DB        *gorm.DB
	SYS_REDIS     *redis.Client
	SYS_CONFIG    config.Server
	SYS_VIP       *viper.Viper
	SYS_LOG       *zap.Logger
	SYS_Timer     timer.Timer = timer.NewTimerTask()
	SYS_IPQuery   qqwry.IPQuery
	BlackCache    local_cache.Cache
	SYS_Enforcer  *casbin.SyncedEnforcer
	SYS_StartTime time.Time
)
