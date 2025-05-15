package router

func (r *Router) loadRoute() error {
	// Guest
	guest := r.e.Group("/api/v1/guest")
	r.h.InitGuestRoute(guest)
	// Admin
	admin := r.e.Group("/api/v1/user", r.m.AdminAuth)
	r.h.InitAdminRoute(admin)
	// Server
	sn := r.e.Group("/api/v1/server", r.m.ServerAuth)
	r.h.InitServerRoute(sn)
	return nil
}
