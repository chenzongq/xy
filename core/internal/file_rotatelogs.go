package internal

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/tcxcn/finance-a/global"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
// Author [linx]
func (r *fileRotatelogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {

	maxAge := time.Duration(-1)
	if global.GVA_CONFIG.Zap.MaxAge != -1 {
		maxAge = time.Duration(global.GVA_CONFIG.Zap.MaxAge) * 24 * time.Hour
	}

	fileWriter, err := rotatelogs.New(
		path.Join(global.GVA_CONFIG.Zap.Director, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(maxAge), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if global.GVA_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
