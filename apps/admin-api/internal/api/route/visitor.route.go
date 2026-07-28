package route

import (
	"github.com/dollarsignteam/go-logger"

	"auzy-api/internal/api/constant"
	"auzy-api/internal/api/controller"
	"auzy-api/internal/api/middleware"
	"auzy-api/lib"
)

type VisitorRoute struct {
	logger               *logger.Logger
	handler              *lib.HttpHandler
	ctrl                 controller.VisitorController
	jwtAuthMiddleware    middleware.JWTAuthMiddleware
	permissionMiddleware middleware.PermissionMiddleware
	rateLimitMiddleware  middleware.RateLimitMiddleware
}

func NewVisitorRoute(
	logger *logger.Logger,
	handler *lib.HttpHandler,
	ctrl controller.VisitorController,
	jwtAuthMiddleware middleware.JWTAuthMiddleware,
	permissionMiddleware middleware.PermissionMiddleware,
	rateLimitMiddleware middleware.RateLimitMiddleware,
) VisitorRoute {
	return VisitorRoute{
		logger:               logger,
		handler:              handler,
		ctrl:                 ctrl,
		jwtAuthMiddleware:    jwtAuthMiddleware,
		permissionMiddleware: permissionMiddleware,
		rateLimitMiddleware:  rateLimitMiddleware,
	}
}

func (r VisitorRoute) Setup() {
	r.logger.Info("Setting up visitor route")

	api := r.handler.Engine.Group("/api")
	api.POST(
		"/public/visit",
		r.ctrl.RecordVisit,
		r.rateLimitMiddleware.RateLimitByIP(),
	)
	api.GET(
		"/visitor-logs",
		r.ctrl.ListVisitorLogs,
		r.jwtAuthMiddleware.JWTAuth(),
		r.permissionMiddleware.Require(constant.VisitorsRead),
	)
}
