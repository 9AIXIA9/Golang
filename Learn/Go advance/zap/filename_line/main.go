package main

import "go.uber.org/zap"

//让zap输出文件名和行号

func main() {
	//调用zap.AddCaller()返回的选项设置输出文件名和行号
	logger, _ := zap.NewProduction(zap.AddCaller())
	defer logger.Sync()

	logger.Info("hello world")
}
