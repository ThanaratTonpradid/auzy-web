package route

import (
	"github.com/dollarsignteam/go-logger"

	"mini-api/internal/api/controller"
	"mini-api/internal/api/middleware"
	"mini-api/lib"
)

type StaffRoute struct {
	logger            *logger.Logger
	handler           *lib.HttpHandler
	ctrl              controller.StaffController
	authMiddleware    middleware.AuthMiddleware
	jwtAuthMiddleware middleware.JWTAuthMiddleware
}

func NewStaffRoute(
	logger *logger.Logger,
	handler *lib.HttpHandler,
	ctrl controller.StaffController,
	authMiddleware middleware.AuthMiddleware,
	jwtAuthMiddleware middleware.JWTAuthMiddleware,
) StaffRoute {
	return StaffRoute{
		logger:            logger,
		handler:           handler,
		ctrl:              ctrl,
		authMiddleware:    authMiddleware,
		jwtAuthMiddleware: jwtAuthMiddleware,
	}
}

func (r StaffRoute) Setup() {
	r.logger.Info("Setting up auth route")

	api := r.handler.Engine.Group("/api")
	api.GET("/staff/profile", r.ctrl.GetStaffProfile, r.jwtAuthMiddleware.JWTAuth())
}
