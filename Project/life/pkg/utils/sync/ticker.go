package sync

import (
	"sync"
	"time"

	"go.uber.org/zap"
)

const (
	updateTime = 2 * time.Hour
)

// TickerCap 用于定时处理能力更新
type TickerCap struct {
	sync.Mutex
	ticker *time.Ticker
	stop   chan bool
}

// NewTickerCap 初始化并返回一个 TickerCap 对象
func NewTickerCap() *TickerCap {
	return &TickerCap{
		ticker: time.NewTicker(updateTime),
		stop:   make(chan bool),
	}
}

// Start 开始定时更新
func (t *TickerCap) Start() {
	go func() {
		for {
			select {
			case <-t.ticker.C:
				t.UpdateCapability()
			case <-t.stop:
				t.ticker.Stop()
				return
			}
		}
	}()
}

// Stop 停止定时更新
func (t *TickerCap) Stop() {
	t.stop <- true
}

// UpdateCapability 模拟执行的能力更新操作
func (t *TickerCap) UpdateCapability() {
	t.Lock()
	defer t.Unlock()
	// 在这里执行你的能力更新逻辑
	println("Updating capability...")
	if err := UpdateCapability(); err != nil {
		zap.L().Error("update capability failed", zap.Error(err))
	}
}
