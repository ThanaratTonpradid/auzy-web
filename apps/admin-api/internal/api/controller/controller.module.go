package controller

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAuthController),
	fx.Provide(NewStaffController),
	fx.Provide(NewRoleController),
	fx.Provide(NewVisitorController),
)
