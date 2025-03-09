package router

func (r *Router) loadRoute() error {
	// User
	user := r.e.Group("/api/v1/user", r.m.UserAuth)
	un := user.Group("/node")
	un.Handle("GET", "/list", r.h.Node.List)
	un.Handle("POST", "/create", r.h.Node.Create)
	un.Handle("POST", "/update", r.h.Node.Update)
	un.Handle("POST", "/delete", r.h.Node.Delete)
	un.Handle("POST", "/get", r.h.Node.Get)
	server := r.e.Group("/server", r.m.ServerAuth)

	// Server
	sn := server.Group("/api/v1/server", r.m.ServerAuth)
	sn.Handle("POST", "/node/get", r.h.Node.Get)
	return nil
}
