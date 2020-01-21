// @Author: Perry
// @Date  : 2020/1/13
// @Desc  : UBER公司的一个高性能log库

package zap_test

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	var url = "www.baidu.com"
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL:%s", url)
}

func Test2(t *testing.T) {
	var url = "www.baidu.com"
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func Test3(t *testing.T) {
	url := "Hello"
	logger, _ := zap.NewProduction()
	defer func() {
		recover()
	}()
	//Sync刷新任何缓冲的日志条目。
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	logger.Warn("debug log", zap.String("level", url))
	logger.Error("Error Message", zap.String("error", url))
	logger.Panic("Panic log", zap.String("level", url))
}

