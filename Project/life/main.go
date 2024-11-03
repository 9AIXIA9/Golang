package main

import (
	"context"
	"errors"
	"fmt"
	"life/api"
	"life/internal/controllers"
	"life/internal/dao/mysql"
	"life/internal/dao/redis"
	"life/internal/server"
	"life/internal/settings"
	"life/pkg/utils/logger"
	"life/pkg/utils/snowflake"
	"life/pkg/utils/sync"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

const (
	timeOutDur = 5 * time.Second
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// 初始化所有组件
	if err := initialize(); err != nil {
		return fmt.Errorf("failed to initialize components: %v", err)
	}

	// 确保资源最终被释放
	defer func() {
		if err := mysql.Close(); err != nil {
			zap.L().Error("failed to close mysql connection", zap.Error(err))
		}
		if err := redis.Close(); err != nil {
			zap.L().Error("failed to close redis connection", zap.Error(err))
		}
		sync.Close()
	}()

	// 设置路由并创建服务器
	router := api.SetupRoutes()
	srv := server.New(settings.Config.App.Port, router)

	// 创建用于通知的通道
	quit := make(chan os.Signal, 1)
	// 监听中断信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 在独立的 goroutine 中启动服务器
	go func() {

		zap.L().Info("starting server...")
		if err := srv.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Fatal("server start failed", zap.Error(err))
		}
	}()

	// 等待中断信号
	<-quit
	zap.L().Info("shutting down server...")

	// 创建一个 5 秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), timeOutDur)
	//给服务器一个最长等待时间（5秒）
	//如果在5秒内没有完成关闭，就强制关闭
	//防止服务器关闭时间过长
	defer cancel()

	// 优雅关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server: %v", err)
	}
	zap.L().Info("server exited")
	return nil
}

func initialize() error {
	// 按照依赖顺序初始化组件
	initializers := []struct {
		name string
		fn   func() error
	}{
		{"config", settings.Init},
		{"logger", logger.Init},
		{"mysql", mysql.Init},
		{"redis", redis.Init},
		{"sync ticker", sync.InitSync},
		{"snowflake", func() error {
			return snowflake.Init(settings.Config.App.StartTime, settings.Config.App.MachineID)
		}},
		{"validator", func() error {
			return controllers.InitTrans(settings.Config.App.Language)
		}},
	}

	for _, init := range initializers {
		if err := init.fn(); err != nil {
			return fmt.Errorf("%s initialization failed: %v", init.name, err)
		}
		zap.L().Info(fmt.Sprintf("%s initialized successfully", init.name))
	}

	return nil
}
