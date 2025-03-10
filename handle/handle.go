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
	Node NodeHandler
}

func NewHandle(d *data.Data) *Handle {
	h := Handle{
		logger: log.SubLogger("Handle"),
		d:      d,
	}
	h.Handlers = &Handlers{
		Node: NodeHandler{
			Handle: h,
		},
	}
	return &h
}
