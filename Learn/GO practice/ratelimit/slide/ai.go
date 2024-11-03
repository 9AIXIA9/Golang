package slide

import (
	"container/list"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type SlidingWindowLimiter struct {
	mu           sync.Mutex
	window       *list.List
	interval     time.Duration
	maxNum       int
	errorHandler gin.HandlerFunc
}

func NewSlidingWindowLimiter(interval time.Duration, maxNum int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		window:   list.New(),
		interval: interval,
		maxNum:   maxNum,
		errorHandler: func(c *gin.Context) {
			c.String(http.StatusTooManyRequests, "rate limit exceeded")
			c.Abort()
		},
	}
}

func (swl *SlidingWindowLimiter) Window() gin.HandlerFunc {
	return func(c *gin.Context) {
		swl.mu.Lock()
		defer swl.mu.Unlock()

		now := time.Now()

		// Remove outdated timestamps
		for swl.window.Len() > 0 {
			if now.Sub(swl.window.Back().Value.(time.Time)) <= swl.interval {
				break
			}
			swl.window.Remove(swl.window.Back())
		}

		// Check if limit is exceeded
		if swl.window.Len() >= swl.maxNum {
			swl.errorHandler(c)
			return
		}

		// Add current timestamp
		swl.window.PushFront(now)

		c.Next()
	}
}
