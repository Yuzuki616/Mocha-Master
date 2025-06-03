package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (r *Router) loadRoute() error {
	// Swagger documentation
	r.e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Guest
	guest := r.e.Group("guest")
	r.h.InitGuestRoute(guest)
	// Admin
	admin := r.e.Group("admin", r.m.AdminAuth)
	r.h.InitAdminRoute(admin)
	// Server
	sn := r.e.Group("server", r.m.ServerAuth)
	r.h.InitServerRoute(sn)
	return nil
}
