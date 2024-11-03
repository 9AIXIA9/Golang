package ratelimit

//令牌桶限流
import (
	"life/internal/controllers"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// TokenRLMiddleware token令牌桶速率限制
func TokenRLMiddleware(fillInternal time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInternal, cap)
	return func(c *gin.Context) {
		//如果取不到令牌就返回响应
		if bucket.TakeAvailable(1) == 0 {
			zap.L().Info("user visit too frequent")
			controllers.ResponseRateLimit(c)
			c.Abort()
			return
		}
		//取到令牌放行
		c.Next()
	}
}
