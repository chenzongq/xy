package mian

import (
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

	global.XY_DB = initialize.InitGorm()

}
