package handle

import (
	"github.com/Yuzuki616/Mocha-Master/data"
	"github.com/Yuzuki616/Mocha-Master/log"
	"go.uber.org/zap"
)

type Handle struct {
	logger *zap.Logger
	d      *data.Data
	*Handlers
}

type Handlers struct {
	Rule   ServerHandler
	Server ServerHandler
	User   *UserHandler
}

func NewHandle(d *data.Data) *Handle {
	h := Handle{
		logger: log.SubLogger("Handle"),
		d:      d,
	}
	h.Handlers = &Handlers{
		Rule: ServerHandler{
			Handle: h,
		},
		Server: ServerHandler{
			Handle: h,
		},
		User: &UserHandler{
			Handle: &h,
		},
	}
	return &h
}
