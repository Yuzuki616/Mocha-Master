package handle

import (
	"Mocha-Master/log"
	"go.uber.org/zap"
)

type Handle struct {
	logger *zap.Logger
}

func NewHandle() *Handle {
	return &Handle{
		logger: log.SubLogger("Http/Handler"),
	}
}
