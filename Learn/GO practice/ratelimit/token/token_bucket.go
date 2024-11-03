package token

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// RLMiddleware token令牌桶速率限制
func RLMiddleware(fillInternal time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInternal, cap)
	return func(c *gin.Context) {
		//如果取不到令牌就返回响应
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusTooManyRequests, "rate limit exceeded")
			c.Abort()
			return
		}
		//取到令牌放行
		c.Next()
	}
}
