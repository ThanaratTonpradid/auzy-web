package route

import (
	"github.com/dollarsignteam/go-logger"

	"auzy-api/internal/api/constant"
	"auzy-api/internal/api/controller"
	"auzy-api/internal/api/middleware"
	"auzy-api/lib"
)

type StaffRoute struct {
	logger                *logger.Logger
	handler               *lib.HttpHandler
	ctrl                  controller.StaffController
	jwtAuthMiddleware     middleware.JWTAuthMiddleware
	permissionMiddleware  middleware.PermissionMiddleware
}

func NewStaffRoute(
	logger *logger.Logger,
	handler *lib.HttpHandler,
	ctrl controller.StaffController,
	jwtAuthMiddleware middleware.JWTAuthMiddleware,
	permissionMiddleware middleware.PermissionMiddleware,
) StaffRoute {
	return StaffRoute{
		logger:               logger,
		handler:              handler,
		ctrl:                 ctrl,
		jwtAuthMiddleware:    jwtAuthMiddleware,
		permissionMiddleware: permissionMiddleware,
	}
}

func (r StaffRoute) Setup() {
	r.logger.Info("Setting up staff route")

	api := r.handler.Engine.Group("/api")
	auth := r.jwtAuthMiddleware.JWTAuth()

	api.GET("/staff/profile", r.ctrl.GetStaffProfile, auth)
	api.GET("/staff", r.ctrl.ListStaff, auth, r.permissionMiddleware.Require(constant.StaffsRead))
	api.GET("/staff/:id", r.ctrl.GetStaff, auth, r.permissionMiddleware.Require(constant.StaffsRead))
	api.POST("/staff", r.ctrl.CreateStaff, auth, r.permissionMiddleware.Require(constant.StaffsCreate))
	api.PUT("/staff/:id", r.ctrl.UpdateStaff, auth, r.permissionMiddleware.Require(constant.StaffsUpdate))
	api.DELETE("/staff/:id", r.ctrl.DeleteStaff, auth, r.permissionMiddleware.Require(constant.StaffsDelete))
}
