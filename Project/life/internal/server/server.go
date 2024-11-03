package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	addr   string
	engine *gin.Engine
	srv    *http.Server
}

// New 新建服务器对象
func New(addr string, engine *gin.Engine) *Server {
	return &Server{
		addr:   addr,
		engine: engine,
	}
}

// Start 启动服务器
func (s *Server) Start() error {
	s.srv = &http.Server{
		Addr:    s.addr,
		Handler: s.engine,
	}
	return s.srv.ListenAndServe()
}

// Shutdown 优雅关闭服务器
func (s *Server) Shutdown(ctx context.Context) error {
	if s.srv != nil {
		return s.srv.Shutdown(ctx)
	}
	return nil
}
