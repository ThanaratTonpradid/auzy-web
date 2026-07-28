package lib

import (
	"time"

	"github.com/golang-jwt/jwt"

	"mini-api/helper"
)

type JWTToken struct {
	ID        string
	Token     string
	IssuedAt  int64
	ExpiresAt int64
}

type JWTOptions struct {
	JWTSecret            []byte
	JWTExpiresTTL        time.Duration
	RefreshTokenExpiresTTL time.Duration
}

type JWTHandler struct {
	Options JWTOptions
}

func NewJWTHandler(opts JWTOptions) *JWTHandler {
	return &JWTHandler{
		Options: opts,
	}
}

func (h JWTHandler) CreateToken(subject string) (JWTToken, error) {
	id := helper.UUID()
	issuedAt := time.Now().Unix()
	expiresAt := time.Now().Add(h.Options.JWTExpiresTTL).Unix()
	claims := &jwt.StandardClaims{
		Id:        id,
		ExpiresAt: expiresAt,
		IssuedAt:  issuedAt,
		Subject:   subject,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.Options.JWTSecret)
	return JWTToken{
		ID:        id,
		Token:     tokenString,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}, err
}

func (h JWTHandler) CreateRefreshToken(subject string) (JWTToken, error) {
	id := helper.UUID()
	issuedAt := time.Now().Unix()
	expiresAt := time.Now().Add(h.Options.RefreshTokenExpiresTTL).Unix()
	claims := &jwt.StandardClaims{
		Id:        id,
		ExpiresAt: expiresAt,
		IssuedAt:  issuedAt,
		Subject:   subject,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.Options.JWTSecret)
	return JWTToken{
		ID:        id,
		Token:     tokenString,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}, err
}

func (h JWTHandler) VerifyToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return h.Options.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
