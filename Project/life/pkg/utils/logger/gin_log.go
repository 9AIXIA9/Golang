package logger

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinLogger 自定义接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery 用于恢复可能出现的panic
func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var ne *net.OpError
				if errors.As(err.(error), &ne) {
					var se *os.SyscallError
					if errors.As(ne.Err, &se) {
						if isConnectionError(se) {
							//检查是否为断开的连接，这种情况不值得记录完整的堆栈跟踪
							//这种错误不值得详细记录
							logBrokenPipe(c, err)
							c.Abort()
							return
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false) //将 HTTP 请求转储（dump）为字节切片
				//当发生 panic 时，这个函数帮助捕获当时的 HTTP 请求信息
				//这对于后续的错误分析非常有价值，因为它提供了触发 panic 的请求的完整上下文
				headers := make(map[string][]string)
				for k, v := range c.Request.Header {
					headers[k] = v
				}

				zap.L().Error("[Recovery from panic]",
					zap.Time("time", time.Now()),
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.Any("headers", headers),
					zap.String("stack", string(debug.Stack())),
				)

				c.AbortWithStatus(http.StatusInternalServerError)
				//在客户端发出请求后，服务器在处理请求过程中发生了未知的问题，导致服务器无法完成请求
			}
		}()
		c.Next()
	}

}

// 发生 "broken pipe" 或 "connection reset by peer"
// 这类网络错误时，logBrokenPipe 会被调用。
// 这些错误通常发生在客户端突然断开连接的情况下
func logBrokenPipe(c *gin.Context, err interface{}) {
	httpRequest, _ := httputil.DumpRequest(c.Request, false)
	zap.L().Error(c.Request.URL.Path,
		zap.Any("error", err),
		zap.String("request", string(httpRequest)),
	)
}

// 识别特定类型的网络连接错误
// "broken pipe"（管道破裂）
// "connection reset by peer"（对方重置连接）
func isConnectionError(se *os.SyscallError) bool {
	errStr := se.Error()
	//TOLower将字母转化为小写
	return strings.Contains(strings.ToLower(errStr), "broken pipe") ||
		strings.Contains(strings.ToLower(errStr), "connection reset by peer")
}
