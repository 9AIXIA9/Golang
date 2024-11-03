package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	//程序退出之前，把缓冲区日志刷进磁盘
	defer logger.Sync()
	//simpleHttpGet("https://www.liwenzhou.com/posts/Go/zap-in-gin/#c-0-0-1")
	//simpleHttpGet("https://www.liwenzhou.com")
	//测试分割日志
	for i := 0; i < 4000; i++ {
		logger.Info("test.......")
	}
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller())
	//zap.AddCaller()用于添加函数调用信息
	sugarLogger = logger.Sugar()
}
func getEncoder() zapcore.Encoder {
	//json格式编码
	//return zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	//空格分割
	//return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)

}
func getLogWriter() zapcore.WriteSyncer {
	//file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	//return zapcore.AddSync(file)
	//日志切割归档功能
	//zap本身不支持要借助Lumberjack
	//目前只支持按文件大小切割，原因是按时间切割效率低且不能保证日志数据不被破坏。
	//详情戳https://github.com/natefinch/lumberjack/issues/54。
	//想按日期切割可以使用github.com/lestrrat-go/file-rotatelogs这个库，虽然目前不维护了，但也够用了。
	lumberJackLogger := &lumberjack.Logger{

		Filename:   "./text.log",
		MaxSize:    1,     //单位 M
		MaxBackups: 5,     // 备份数量
		MaxAge:     30,    // 备份天数
		Compress:   false, //是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
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
