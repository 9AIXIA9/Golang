package fix

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type WindowLimiter struct {
	mu           sync.Mutex
	lastTime     time.Time
	presentNum   int64
	interval     time.Duration
	maxNum       int64
	errorHandler gin.HandlerFunc
}

func NewWindowLimiter(interval time.Duration, maxNum int64) *WindowLimiter {
	return &WindowLimiter{
		interval: interval,
		maxNum:   maxNum,
		errorHandler: func(c *gin.Context) {
			c.String(http.StatusTooManyRequests, "rate limit exceeded")
			c.Abort()
		},
	}
}

func (wl *WindowLimiter) SetErrorHandler(handler gin.HandlerFunc) {
	wl.errorHandler = handler
}

func (wl *WindowLimiter) Window() gin.HandlerFunc {
	return func(c *gin.Context) {
		wl.mu.Lock()
		defer wl.mu.Unlock()

		now := time.Now()
		if wl.lastTime.IsZero() {
			wl.lastTime = now
		}

		if now.Sub(wl.lastTime) >= wl.interval {
			wl.lastTime = now
			wl.presentNum = 0
		}

		if wl.presentNum >= wl.maxNum {
			wl.errorHandler(c)
			return
		}

		wl.presentNum++
		c.Next()
	}
}
