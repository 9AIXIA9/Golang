package middlewares

import (
	"net/http"
	"time"

	"github.com/juju/ratelimit"

	"github.com/gin-gonic/gin"
)

// RateLimitMiddleware 令牌桶
func RateLimitMiddleware(fillInternal time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInternal, cap)
	return func(c *gin.Context) {
		//如果取不到令牌就返回响应
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		//取到令牌放行
		c.Next()
	}
}