package router

import (
	"fmt"
	_ "github.com/Yuzuki616/Mocha-Master/docs" // swagger docs
	"github.com/Yuzuki616/Mocha-Master/handle"
	"github.com/Yuzuki616/Mocha-Master/middleware"
	"github.com/gin-contrib/cors"
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
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))
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
