package slide

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	timeline []time.Time
)

func Window(interval time.Duration, maxNum int) gin.HandlerFunc {
	return func(c *gin.Context) {
		//初始化
		if len(timeline) == 0 {
			timeline = make([]time.Time, maxNum)
			timeline[0] = time.Now()
			c.Next()
			return
		}
		now := time.Now()
		// 消去所有差值大于interval的时间点
		end := -1
		for i, v := range timeline {
			if v.IsZero() {
				end = i
				break
			}
			if now.Sub(v) > interval {
				end = i
				break
			}
		}

		//都未超过interval即超载
		if end == -1 {
			c.String(http.StatusTooManyRequests, "rate limit...")
			c.Abort()
			return
		}
		copy(timeline[1:], timeline[:end])
		timeline[0] = now
		c.Next()
	}
}
