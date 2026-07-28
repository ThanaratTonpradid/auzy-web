package middleware

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"auzy-api/internal/api/constant"
	"auzy-api/internal/api/dto"
	"auzy-api/lib"
)

type PermissionMiddleware struct{}

func NewPermissionMiddleware() PermissionMiddleware {
	return PermissionMiddleware{}
}

func (mw PermissionMiddleware) Require(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, ok := c.Get(constant.KeySession).(dto.Session)
			if !ok || session.StaffID == 0 {
				return lib.CommonError{
					StatusCode:    http.StatusUnauthorized,
					ErrorCode:     constant.ErrCodeUnauthorized,
					ErrorInstance: errors.New("session not found"),
				}
			}
			if session.IsAdmin {
				return next(c)
			}
			for _, code := range session.Permissions {
				if code == permission {
					return next(c)
				}
			}
			return lib.CommonError{
				StatusCode:    http.StatusForbidden,
				ErrorCode:     constant.ErrCodeForbidden,
				ErrorInstance: errors.New("insufficient permissions"),
			}
		}
	}
}
