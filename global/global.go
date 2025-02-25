package global

import (
	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"github.com/xyxy/admin-service/config"
	"github.com/xyxy/admin-service/utils/timer"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	XY_VP      *viper.Viper
	XY_CONFIG  config.Server
	XY_LOG     *zap.Logger
	XY_DB      *gorm.DB
	Snowflake  *snowflake.Node
	XY_Timer   timer.Timer = timer.NewTimerTask()
	XY_REDIS   *redis.Client
	BlackCache local_cache.Cache
)
