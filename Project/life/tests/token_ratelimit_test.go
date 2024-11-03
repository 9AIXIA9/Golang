package tests

import (
	"life/pkg/middlewares/ratelimit"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter(middleware gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middleware)
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})
	return router
}

func performRequest(router *gin.Engine, method, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, nil)
	router.ServeHTTP(w, req)
	return w
}

func TestTokenRLMiddleware(t *testing.T) {
	type args struct {
		fillInterval time.Duration
		cap          int64
	}
	tests := []struct {
		name          string
		args          args
		requestCount  int           // 发送请求的次数
		requestGap    time.Duration // 请求之间的时间间隔
		expectedCodes []int         // 期望的HTTP状态码
	}{
		{
			name: "normal_rate",
			args: args{
				fillInterval: time.Second,
				cap:          5,
			},
			requestCount:  3,
			requestGap:    100 * time.Millisecond,
			expectedCodes: []int{200, 200, 200},
		},
		{
			name: "exceed_rate",
			args: args{
				fillInterval: time.Second,
				cap:          2,
			},
			requestCount:  4,
			requestGap:    100 * time.Millisecond,
			expectedCodes: []int{200, 200, 429, 429},
		},
		{
			name: "recover_after_wait",
			args: args{
				fillInterval: 100 * time.Millisecond,
				cap:          1,
			},
			requestCount:  3,
			requestGap:    150 * time.Millisecond, // 等待时间大于填充间隔
			expectedCodes: []int{200, 200, 200},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			middleware := ratelimit.TokenRLMiddleware(tt.args.fillInterval, tt.args.cap)
			router := setupRouter(middleware)

			for i := 0; i < tt.requestCount; i++ {
				w := performRequest(router, "GET", "/test")
				if w.Code != tt.expectedCodes[i] {
					t.Errorf("request %d: got status code %d, want %d",
						i, w.Code, tt.expectedCodes[i])
				}
				if i < tt.requestCount-1 {
					time.Sleep(tt.requestGap)
				}
			}
		})
	}
}
