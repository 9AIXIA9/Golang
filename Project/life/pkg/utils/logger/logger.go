package logger

import (
	"errors"
	"fmt"
	"life/internal/settings"
	"os"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	timeKey   = "time"
	debugMode = gin.DebugMode
)

var (
	errorLogger = errors.New("logger init failed")
)

// Init 初始化logger
func Init() (err error) {
	config := settings.Config.LogConf
	mode := settings.Config.App.Mode
	writeSyncer := getLogWriter(
		config.Filename,
		config.MaxSize,
		config.MaxBackups,
		config.MaxAge,
	)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(config.Level))
	if err != nil {
		return fmt.Errorf("%w,error:%v", errorLogger, err)
	}

	var core zapcore.Core
	if mode == debugMode {
		//开发模式,日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			//两个core，前一个写到文件，后一个写到终端
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zap.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}
	lg := zap.New(core, zap.AddCaller())
	//替换全局logger对象
	zap.ReplaceGlobals(lg)
	return
}

// 设置编码器配置
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()             //创建新的编码器配置
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder         //设置时间格式为ISO 08601
	encoderConfig.TimeKey = timeKey                               //时间输出字段名为time
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder       //设置日志级别以大写形式输出“ERROR”
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder //设置持续时间以秒为单位输出
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder       //设置调用者信息以简短形式输出（文件名和行号）
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 设置日志文件
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{ //用于分割日志
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
