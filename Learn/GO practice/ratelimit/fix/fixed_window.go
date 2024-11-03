package fix

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	presentNum int64
	lastTime   time.Time
)

func Window(interval time.Duration, maxNum int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		//如果还没开始计数
		if lastTime.IsZero() {
			lastTime = time.Now()
		}
		//现在的时间与启动时间的间距
		gap := time.Now().Sub(lastTime)
		//判断时间是否超过间距
		if gap < interval && presentNum >= maxNum {
			c.String(http.StatusTooManyRequests, "rate limit...")
			c.Abort()
			return
		}
		if gap >= interval {
			lastTime = time.Now()
			presentNum = 0
		}
		presentNum++
		c.Next()
	}
}
