package internal

import (
	"github.com/xyxy/admin-service/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type DBBASE interface {
	GetLogMode() string
}

type _gorm struct{}

var Gorm = new(_gorm)

// Config gorm 自定义配置
// Author [chenzongqing]
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      false,
	})
	//var logMode DBBASE
	//switch global.XY_CONFIG.System.DbType {
	//case "mysql":
	//	logMode = &global.XY_CONFIG.Mysql
	//case "pgsql":
	//	logMode = &global.XY_CONFIG.Pgsql
	//default:
	//	logMode = &global.XY_CONFIG.Mysql
	//}

	var logMode DBBASE
	logMode = &global.XY_CONFIG.Mysql

	switch logMode.GetLogMode() {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}
