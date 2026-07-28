package route

import (
	"github.com/dollarsignteam/go-logger"

	"auzy-api/internal/api/controller"
	"auzy-api/internal/api/middleware"
	"auzy-api/lib"
)

type AuthRoute struct {
	logger            *logger.Logger
	handler           *lib.HttpHandler
	ctrl              controller.AuthController
	authMiddleware    middleware.AuthMiddleware
	jwtAuthMiddleware middleware.JWTAuthMiddleware
}

func NewAuthRoute(
	logger *logger.Logger,
	handler *lib.HttpHandler,
	ctrl controller.AuthController,
	authMiddleware middleware.AuthMiddleware,
	jwtAuthMiddleware middleware.JWTAuthMiddleware,
) AuthRoute {
	return AuthRoute{
		logger:            logger,
		handler:           handler,
		ctrl:              ctrl,
		authMiddleware:    authMiddleware,
		jwtAuthMiddleware: jwtAuthMiddleware,
	}
}

func (r AuthRoute) Setup() {
	r.logger.Info("Setting up auth route")

	api := r.handler.Engine.Group("/api")
	api.POST("/auth/login", r.ctrl.Login)
	api.POST("/auth/refresh", r.ctrl.RefreshToken)
	api.POST("/auth/logout", r.ctrl.Logout, r.jwtAuthMiddleware.JWTAuth())
}
