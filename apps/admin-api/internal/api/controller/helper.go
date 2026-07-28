package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"auzy-api/internal/api/constant"
	"auzy-api/internal/api/dto"
	"auzy-api/lib"
)

func GetSession(c echo.Context) dto.Session {
	if session, ok := c.Get(constant.KeySession).(dto.Session); ok {
		return session
	}
	return dto.Session{}
}

func NewCommonErrorBadRequest(err error) error {
	return lib.CommonError{
		StatusCode:    http.StatusBadRequest,
		ErrorCode:     constant.ErrCodeBadRequest,
		ErrorInstance: err,
	}
}
