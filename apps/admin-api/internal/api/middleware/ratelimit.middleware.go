package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

type RateLimitMiddleware struct{}

func NewRateLimitMiddleware() RateLimitMiddleware {
	return RateLimitMiddleware{}
}

// RateLimit applies rate limiting to requests
func (m RateLimitMiddleware) RateLimit() echo.MiddlewareFunc {
	// Create a rate limiter: 20 requests per second with burst of 30
	limiter := rate.NewLimiter(rate.Every(time.Second/20), 30)
	
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !limiter.Allow() {
				return echo.NewHTTPError(
					http.StatusTooManyRequests,
					"Rate limit exceeded. Please try again later.",
				)
			}
			return next(c)
		}
	}
}

// RateLimitByIP applies rate limiting per IP address
func (m RateLimitMiddleware) RateLimitByIP() echo.MiddlewareFunc {
	// Use echo's built-in rate limiter with per-IP tracking
	return middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20))
}

// RateLimitStrict applies stricter rate limiting (e.g., for login endpoints)
func (m RateLimitMiddleware) RateLimitStrict() echo.MiddlewareFunc {
	// 5 requests per minute
	limiter := rate.NewLimiter(rate.Every(time.Minute/5), 10)
	
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !limiter.Allow() {
				return echo.NewHTTPError(
					http.StatusTooManyRequests,
					"Too many attempts. Please try again later.",
				)
			}
			return next(c)
		}
	}
}

