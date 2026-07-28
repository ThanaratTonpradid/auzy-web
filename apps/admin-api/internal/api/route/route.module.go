package route

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewSwaggerRoute),
	fx.Provide(NewAuthRoute),
	fx.Provide(NewStaffRoute),
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	swagger SwaggerRoute,
	auth AuthRoute,
	staff StaffRoute,
) Routes {
	return Routes{
		swagger,
		auth,
		staff,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
