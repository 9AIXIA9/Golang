package main

import (
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	//程序退出之前，把缓冲区日志刷进磁盘
	defer logger.Sync()
	simpleHttpGet("www.google.com")
	simpleHttpGet2("http://www.google.com")

}

func InitLogger() {
	logger, _ = zap.NewProduction()
	//logger, _ = zap_demo.NewDevelopment()
	//NewProduction是json格式
	//NewDevelopment是终端形式的日志
	sugarLogger = logger.Sugar()
}
func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url ..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info(
			"Success",
			zap.String("status code", resp.Status),
			zap.String("url", url),
		)
		resp.Body.Close()
	}
}
func simpleHttpGet2(url string) {
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Error(
			"Error fetching url ..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		sugarLogger.Info(
			"Success",
			zap.String("status code", resp.Status),
			zap.String("url", url),
		)
		resp.Body.Close()
	}
}
