package route

import (
	"github.com/dollarsignteam/go-logger"

	"auzy-api/internal/api/constant"
	"auzy-api/internal/api/controller"
	"auzy-api/internal/api/middleware"
	"auzy-api/lib"
)

type RoleRoute struct {
	logger               *logger.Logger
	handler              *lib.HttpHandler
	ctrl                 controller.RoleController
	jwtAuthMiddleware    middleware.JWTAuthMiddleware
	permissionMiddleware middleware.PermissionMiddleware
}

func NewRoleRoute(
	logger *logger.Logger,
	handler *lib.HttpHandler,
	ctrl controller.RoleController,
	jwtAuthMiddleware middleware.JWTAuthMiddleware,
	permissionMiddleware middleware.PermissionMiddleware,
) RoleRoute {
	return RoleRoute{
		logger:               logger,
		handler:              handler,
		ctrl:                 ctrl,
		jwtAuthMiddleware:    jwtAuthMiddleware,
		permissionMiddleware: permissionMiddleware,
	}
}

func (r RoleRoute) Setup() {
	r.logger.Info("Setting up role route")

	api := r.handler.Engine.Group("/api")
	auth := r.jwtAuthMiddleware.JWTAuth()

	api.GET("/roles", r.ctrl.ListRoles, auth, r.permissionMiddleware.Require(constant.RolesRead))
	api.GET("/roles/:id", r.ctrl.GetRole, auth, r.permissionMiddleware.Require(constant.RolesRead))
	api.PUT("/roles/:id/permissions", r.ctrl.UpdateRolePermissions, auth, r.permissionMiddleware.Require(constant.RolesUpdate))
	api.GET("/permissions", r.ctrl.ListPermissions, auth, r.permissionMiddleware.Require(constant.RolesRead))
}
