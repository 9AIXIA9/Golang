package main

import (
	"bluebell/controllers"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/myerrors"
	"bluebell/pkg/snowflake"
	"bluebell/routes"
	"bluebell/settings"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

//go web开发较通用的脚手架

func main() {
	//1.加载配置文件
	if err := settings.Init(); err != nil {
		log.Fatal(errors.Join(myerrors.InitSettings, err))
	}

	//2.初始化日志
	if err := logger.Init(settings.Config.LogConfig, settings.Config.App.Mode); err != nil {
		log.Fatal(errors.Join(myerrors.InitLogger, err))
	}
	defer func() {
		if err := zap.L().Sync(); err != nil {
			zap.L().Fatal(errors.Join(myerrors.ShutDownServer, err).Error())
		}
	}()
	zap.L().Debug("logger init success...")

	//3.初始化Mysql连接
	if err := mysql.Init(settings.Config.MysqlConfig); err != nil {
		zap.L().Fatal(errors.Join(myerrors.InitMysql, err).Error())
	}
	defer mysql.Close()
	//4.初始化Redis连接
	if err := redis.Init(settings.Config.RedisConfig); err != nil {
		zap.L().Fatal(errors.Join(myerrors.InitRedis, err).Error())
	}
	defer redis.Close()

	//雪花算法生成ID的初始化
	if err := snowflake.Init(settings.Config.App.StartTime, settings.Config.App.MachineID); err != nil {
		zap.L().Fatal(errors.Join(myerrors.InitSnowflake, err).Error())
	}

	//5.注册路由
	//初始化翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		zap.L().Fatal(errors.Join(myerrors.InitValidator, err).Error())
	}

	r := routes.SetUp(settings.Config.App.Mode)

	//6.启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", settings.Config.App.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Fatal(errors.Join(myerrors.GoRoutine, err).Error())
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Fatal(myerrors.ShutDownServer.Error())
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal(errors.Join(myerrors.ShutDownServer, err).Error())
	}

	zap.L().Fatal(myerrors.ShutDownServer.Error())
}
