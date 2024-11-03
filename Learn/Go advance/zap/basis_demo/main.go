package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	log := zap.NewExample()
	defer log.Sync() //将缓存同步到文件中

	url := "https://example.org/api"
	//默认的Logger只支持强类型的、结构化的日志
	//必须使用zap提供的方法记录字段
	log.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	sugarLog := log.Sugar()
	//加糖易操作
	//无需使用zap.int等来传入对应类型的变量
	sugarLog.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugarLog.Info("failed to fetch URL: %s", url) //极简版
}
