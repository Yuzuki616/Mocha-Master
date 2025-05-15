package handle

import (
	"github.com/Yuzuki616/Mocha-Master/conf"
	"github.com/Yuzuki616/Mocha-Master/data"
	"github.com/Yuzuki616/Mocha-Master/handle/admin"
	"github.com/Yuzuki616/Mocha-Master/handle/common/context"
	"github.com/Yuzuki616/Mocha-Master/handle/guest"
	"github.com/Yuzuki616/Mocha-Master/handle/server"
	"github.com/Yuzuki616/Mocha-Master/log"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/gin-gonic/gin"
)

type Handle struct {
	Admin  admin.Handler
	Guest  guest.Handler
	Server server.Handler
}

func NewHandle(d *data.Data, cache *cache.Cache[any], cf *conf.Conf) *Handle {
	c := context.NewContext(log.SubLogger("Handler"), d, cache, cf)
	return &Handle{
		Admin: admin.Handler{
			Context: c,
		},
		Guest: guest.Handler{
			Context: c,
		},
		Server: server.Handler{
			Context: c,
		},
	}
}

func (h *Handle) InitGuestRoute(r *gin.RouterGroup) {
	r.Handle("POST", "tokenCheck", h.Guest.TokenCheck)
}

func (h *Handle) InitAdminRoute(r *gin.RouterGroup) {
	un := r.Group("/server")
	un.Handle("GET", "/list", h.Admin.ListServerHandle)
	un.Handle("POST", "/create", h.Admin.CreateServerHandle)
	un.Handle("POST", "/update", h.Admin.UpdateServerHandle)
	un.Handle("POST", "/delete", h.Admin.DeleteServerHandle)
	un.Handle("POST", "/get", h.Admin.GetServerHandle)
	ru := r.Group("/rule")
	ru.Handle("POST", "/list", h.Admin.ListRuleHandle)
	ru.Handle("POST", "/create", h.Admin.CreateRuleHandle)
	ru.Handle("POST", "/update", h.Admin.UpdateRuleHandle)
	ru.Handle("POST", "/delete", h.Admin.DeleteRuleHandle)
}

func (h *Handle) InitServerRoute(r *gin.RouterGroup) {
	r.Handle("GET", "/getConfig", h.Server.GetConfigHandle)
	r.Handle("POST", "/reportStatus", h.Server.ReportStatusHandle)
}
