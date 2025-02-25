package mian

import (
	"github.com/bwmarrin/snowflake"
	"github.com/xyxy/admin-service/core"
	"github.com/xyxy/admin-service/global"
	"github.com/xyxy/admin-service/initialize"
	"go.uber.org/zap"
)

func main() {

	global.XY_VP = core.Viper() // 初始化viper
	initialize.OtherInit()

	global.XY_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.XY_LOG)

	global.XY_DB = initialize.InitGorm(global.XY_CONFIG.Mysql) // 初始化数据库

	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	global.Snowflake = node //初始话雪花算法实例

	initialize.Timer()

	if global.XY_DB != nil {
		db, _ := global.XY_DB.DB()
		defer db.Close()
	}

	core.RunWindowsServer()
}
