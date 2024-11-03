package tests

import (
	"life/pkg/middlewares/ratelimit"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func TestIPRateLimiter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	url := "/test/ip"
	limiter := ratelimit.NewIPRateLimiter(rate.Every(5*time.Second), 1, time.Minute)
	r.Use(ratelimit.IPRLMiddleware(limiter))
	r.POST(url, func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	//第一次请求
	// 测试成功请求
	t.Run("Successful request", func(t *testing.T) {
		w := performRequest(r, "POST", url, "127.0.0.1")
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "test", w.Body.String())
	})

	//第二次请求
	// 测试频率限制
	t.Run("Rate limited", func(t *testing.T) {
		w := performRequest(r, "POST", url, "127.0.0.1")
		assert.Equal(t, http.StatusTooManyRequests, w.Code)
	})

	// 测试不同IP
	t.Run("Different IP not limited", func(t *testing.T) {
		w := performRequest(r, "POST", url, "127.0.0.2")
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// 测试限制重置
	t.Run("Limit reset after duration", func(t *testing.T) {
		time.Sleep(5 * time.Second)
		w := performRequest(r, "POST", url, "127.0.0.1")
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func performRequest(r *gin.Engine, method, url, ip string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Set("X-Forwarded-For", ip)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
