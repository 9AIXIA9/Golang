package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ratelimit2 "github.com/juju/ratelimit"
	ratelimit1 "go.uber.org/ratelimit"
)

func pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func heiHandler(c *gin.Context) {
	c.String(http.StatusOK, "ha")
}

// 基于漏桶的中间件
func rateLimit1() func(ctx *gin.Context) {
	//生成一个限流器
	rl := ratelimit1.New(1)
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

// 基于令牌桶的中间件
func rateLimit2() func(ctx *gin.Context) {
	rl := ratelimit2.NewBucket(2*time.Second, 1)
	return func(ctx *gin.Context) {
		//rl.Take()//此次可以欠账
		//rl.TakeAvailable() //有令牌才能取
		if rl.TakeAvailable(1) != 1 { //此次未取到令牌
			ctx.String(http.StatusOK, "rate limit...")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", rateLimit1(), pingHandler)
	r.GET("/hei", rateLimit2(), heiHandler)

	_ = r.Run()
}
