package leak

import (
	"log"
	"net/http"
	"time"

	"go.uber.org/ratelimit"

	"github.com/gin-gonic/gin"
)

// RateLimit 基于漏桶的中间件
func RateLimit() func(ctx *gin.Context) {
	//生成一个限流器
	rl := ratelimit.New(3, ratelimit.Per(30*time.Second))
	return func(c *gin.Context) {
		//取水滴
		now := time.Now()
		next := rl.Take()
		if next.After(now) {
			//time.Sleep(rl.Take().Sub(time.Now())) //需要等这么长时间下一滴水才会掉下来
			log.Printf("cant used")
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		log.Printf("u can use")

		c.Next()
	}
}
