package middleware

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"auzy-api/internal/api/constant"
	"auzy-api/internal/api/dto"
	"auzy-api/internal/api/service"
	"auzy-api/lib"
)

type JWTAuthMiddleware struct {
	jwtHandler  *lib.JWTHandler
	authService service.AuthService
}

func NewJWTAuthMiddleware(
	jwtHandler *lib.JWTHandler,
	authService service.AuthService,
) JWTAuthMiddleware {
	return JWTAuthMiddleware{
		jwtHandler:  jwtHandler,
		authService: authService,
	}
}

func (mw JWTAuthMiddleware) JWTAuth() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.RegisteredClaims)
		},
		SigningKey:     mw.jwtHandler.Options.JWTSecret,
		ParseTokenFunc: mw.ParseToken,
	})
}

func (mw JWTAuthMiddleware) ParseToken(c echo.Context, auth string) (interface{}, error) {
	claims := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(auth, &claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != echojwt.AlgorithmHS256 {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return mw.jwtHandler.Options.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	session, err := mw.GetSession(claims)
	if err != nil {
		return nil, err
	}
	c.Set(constant.KeySession, session)
	return token, nil
}

func (mw JWTAuthMiddleware) GetSession(claims jwt.RegisteredClaims) (dto.Session, error) {
	id, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return dto.Session{}, err
	}
	session, err := mw.authService.GetSession(uint32(id))
	if err != nil {
		return dto.Session{}, errors.New("session not found")
	}
	if claims.ID != session.ID {
		return dto.Session{}, errors.New("invalid session")
	}
	return session, nil
}
