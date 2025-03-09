package middleware

import (
	"Mocha-Master/conf"
	"Mocha-Master/log"
	"go.uber.org/zap"
)

type Middleware struct {
	logger *zap.Logger
	c      *conf.Conf
}

func New(c *conf.Conf) *Middleware {
	return &Middleware{
		logger: log.SubLogger("Http/Engine"),
		c:      c,
	}
}
