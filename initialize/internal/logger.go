package internal

import (
	"fmt"
	"github.com/xyxy/admin-service/global"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
// Author [linx]
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
// Author [linx]
func (w *writer) Printf(message string, data ...interface{}) {
	//var logZap bool
	//switch global.XY_CONFIG.System.DbType {
	//case "mysql":
	//	logZap = global.XY_CONFIG.Mysql.LogZap
	//case "pgsql":
	//	logZap = global.XY_CONFIG.Pgsql.LogZap
	//}
	logZap := global.XY_CONFIG.Mysql.LogZap
	if logZap {
		global.XY_LOG.WithOptions(zap.WithCaller(false)).
			Sugar().Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
