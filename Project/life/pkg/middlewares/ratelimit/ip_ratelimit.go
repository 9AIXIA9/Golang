package ratelimit

//ip限流
import (
	"life/internal/controllers"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"

	"golang.org/x/time/rate"
)

type IPRateLimiter struct {
	ips      sync.Map //并发map,防止并发时被出错
	duration rate.Limit
	num      int
	cleanup  time.Duration
}

type rateLimiterEntry struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// IPRLMiddleware 限制同一IP访问速率
func IPRLMiddleware(limiter *IPRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		l := limiter.GetLimiter(ip)
		if !l.Allow() {
			controllers.ResponseRateLimit(c)
			zap.L().Info("user visit too frequent")
			c.Abort()
			return
		}
		c.Next()
	}
}

// NewIPRateLimiter 为IP创建新的速率限制器
func NewIPRateLimiter(duration rate.Limit, num int, cleanup time.Duration) *IPRateLimiter {
	limiter := &IPRateLimiter{
		duration: duration,
		num:      num,
		cleanup:  cleanup,
	}
	go limiter.cleanupLoop()
	return limiter
}

// GetLimiter 获取IP的限制器，如果不存在则创建
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	v, exists := i.ips.Load(ip)
	if !exists {
		limiter := rate.NewLimiter(i.duration, i.num)
		v, _ = i.ips.LoadOrStore(ip, &rateLimiterEntry{
			limiter:  limiter,
			lastSeen: time.Now(),
		})
	}
	entry := v.(*rateLimiterEntry)
	entry.lastSeen = time.Now()
	return entry.limiter
}

// cleanupLoop 定期清理不活跃的IP
func (i *IPRateLimiter) cleanupLoop() {
	ticker := time.NewTicker(i.cleanup) //创建定时器，每过一段时间触发一次
	for range ticker.C {
		now := time.Now()
		i.ips.Range(func(key, value interface{}) bool {
			entry := value.(*rateLimiterEntry)
			if now.Sub(entry.lastSeen) > i.cleanup {
				i.ips.Delete(key)
			}
			return true
		})
	}
}
