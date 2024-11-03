package main

import (
	"fmt"
	"net/http"
	"ratelimit/fix"
	"ratelimit/leak"
	"ratelimit/token"

	"ratelimit/slide"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/fixed", fix.Window(30*time.Second, 3), func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	Limiter := fix.NewWindowLimiter(30*time.Second, 3)
	r.GET("/ai/fixed", Limiter.Window(), func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	r.GET("/slide", slide.Window(30*time.Second, 3), func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	Limiter2 := slide.NewSlidingWindowLimiter(30*time.Second, 3)
	r.GET("/ai/slide", Limiter2.Window(), func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	r.GET("/leak", leak.RateLimit(), func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	r.GET("/token", token.RLMiddleware(30*time.Second, 3), func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	if err := r.Run(); err != nil {
		fmt.Println(err)
	}
}
