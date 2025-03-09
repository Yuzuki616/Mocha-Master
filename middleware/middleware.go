package middleware

import (
	"Mocha-Master/log"
	"go.uber.org/zap"
)

type Middleware struct {
	logger *zap.Logger
}

func New() *Middleware {
	return &Middleware{
		logger: log.SubLogger("Http/Engine"),
	}
}
