package global

import (
	"github.com/spf13/viper"
	"github.com/xyxy/admin-service/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB    *gorm.DB
	XY_VP     *viper.Viper
	XY_CONFIG config.Server
	XY_LOG    *zap.Logger
	XY_DB     *gorm.DB
)
