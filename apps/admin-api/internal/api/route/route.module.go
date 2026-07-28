package route

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewSwaggerRoute),
	fx.Provide(NewAuthRoute),
	fx.Provide(NewStaffRoute),
	fx.Provide(NewRoleRoute),
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	swagger SwaggerRoute,
	auth AuthRoute,
	staff StaffRoute,
	role RoleRoute,
) Routes {
	return Routes{
		swagger,
		auth,
		staff,
		role,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
