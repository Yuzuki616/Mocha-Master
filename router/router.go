package router

import (
	"fmt"
	"github.com/Yuzuki616/Mocha-Master/handle"
	"github.com/Yuzuki616/Mocha-Master/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	e *gin.Engine
	h *handle.Handle
	m *middleware.Middleware
}

func NewRouter(h *handle.Handle, m *middleware.Middleware) *Router {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(m.Logger, gin.Recovery())
	return &Router{
		e: engine,
		h: h,
		m: m,
	}
}

func (r *Router) Start(addr string) error {
	err := r.loadRoute()
	if err != nil {
		return fmt.Errorf("load route err: %v", err)
	}
	return r.e.Run(addr)
}
