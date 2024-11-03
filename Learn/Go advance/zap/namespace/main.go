package main

import "go.uber.org/zap"

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	//1、直接作为字段传入Debug/Info等方法
	logger.Info("tracked some metrics",
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
	)

	//2、调用With()创建一个新的Logger
	//新的Logger记录日志时总是带上预设的字段
	logger2 := logger.With(
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
	)
	logger2.Info("tracked some metrics")
}
