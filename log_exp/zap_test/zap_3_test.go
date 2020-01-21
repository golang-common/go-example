// @Author: Perry
// @Date  : 2020/1/13
// @Desc  : 生产上用的测试例

package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"testing"
)

/*日志序列化为文件*/
func Test6(t *testing.T) {
	logger := initLogger2("/Users/daipengyuan/code/go/src/dpy/exp/log_exp/zap_test/all.log", "info")
	logger.Info("test log", zap.Int("line", 47))
	logger.Warn("test warn", zap.Int("line", 47))
}

func initLogger2(logpath string, loglevel string) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径
		MaxSize:    1024,    // 文件最大多少M,默认100M
		MaxAge:     3,       // 最多保留多少天,默认不根据日期删除
		MaxBackups: 7,       // 最多保留多少个备份,
		LocalTime:  false,   // 是否使用本地时间,默认使用UTC时间
		Compress:   true,    // 是否gzip压缩
	}

	w := zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook), os.Stdout)
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), w, level)
	core2 := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), os.Stdout, level)
	tree := zapcore.NewTee(core, core2)
	logger := zap.New(tree)
	logger.Info("defaultLogger init success")
	return logger
}