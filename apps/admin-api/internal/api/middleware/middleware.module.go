package middleware

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAuthMiddleware),
	fx.Provide(NewJWTAuthMiddleware),
	fx.Provide(NewPermissionMiddleware),
	fx.Provide(NewCORSMiddleware),
	fx.Provide(NewRateLimitMiddleware),
	fx.Provide(NewSanitizeMiddleware),
)
