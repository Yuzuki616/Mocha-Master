package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func (m *Middleware) Logger(c *gin.Context) {
	// Start timer
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	start := time.Now()
	m.logger.Info("Request Handling.",
		zap.String("Method", c.Request.Method),
		zap.Int("Status", c.Writer.Status()),
		zap.String("Source", c.ClientIP()),
		zap.String("Path", path))

	// Process request
	c.Next()
	latency := time.Now().Sub(start)
	if latency > time.Minute {
		latency = latency - latency%time.Second
	}
	if err := c.Err(); err != nil {
		m.logger.Error("Request Handling Error.",
			zap.String("Method", c.Request.Method),
			zap.Int("Status", c.Writer.Status()),
			zap.String("Source", c.ClientIP()),
			zap.String("Path", path),
			zap.Error(err))
		return
	}
	if raw != "" {
		path = path + "?" + raw
	}
	m.logger.Info("Request Handled.",
		zap.String("Method", c.Request.Method),
		zap.Int("Status", c.Writer.Status()),
		zap.String("Source", c.ClientIP()),
		zap.String("Path", path),
		zap.String("Latency", latency.String()))
}
