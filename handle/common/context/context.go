package context

import (
	"github.com/Yuzuki616/Mocha-Master/conf"
	"github.com/Yuzuki616/Mocha-Master/data"
	"github.com/eko/gocache/lib/v4/cache"
	"go.uber.org/zap"
)

type Context struct {
	L     *zap.Logger
	Data  *data.Data
	Cache *cache.Cache[any]
	Conf  *conf.Conf
}

func NewContext(l *zap.Logger, data *data.Data, cache *cache.Cache[any], c *conf.Conf) *Context {
	return &Context{
		L:     l,
		Data:  data,
		Cache: cache,
		Conf:  c,
	}
}
