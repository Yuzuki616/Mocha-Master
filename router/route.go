package router

func (r *Router) loadRoute() error {
	// User
	user := r.e.Group("/api/v1/user", r.m.UserAuth)
	un := user.Group("/server")
	un.Handle("GET", "/list", r.h.Server.List)
	un.Handle("POST", "/create", r.h.Server.Create)
	un.Handle("POST", "/update", r.h.Server.Update)
	un.Handle("POST", "/delete", r.h.Server.Delete)
	un.Handle("POST", "/get", r.h.Server.Get)

	// Server
	sn := r.e.Group("/api/v1/server", r.m.ServerAuth)
	sn.Handle("POST", "/rule/list", r.h.Rule.List)
	return nil
}
