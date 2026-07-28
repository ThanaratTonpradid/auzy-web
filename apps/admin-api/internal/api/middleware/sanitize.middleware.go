package middleware

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"

	"mini-api/helper"
)

type SanitizeMiddleware struct{}

func NewSanitizeMiddleware() SanitizeMiddleware {
	return SanitizeMiddleware{}
}

// SanitizeInput sanitizes request body input
func (m SanitizeMiddleware) SanitizeInput() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Only process POST, PUT, PATCH requests with JSON content type
			if c.Request().Method == echo.POST || 
			   c.Request().Method == echo.PUT || 
			   c.Request().Method == echo.PATCH {
				contentType := c.Request().Header.Get(echo.HeaderContentType)
				if contentType == echo.MIMEApplicationJSON || 
				   contentType == "application/json" {
					// Read body
					bodyBytes, err := io.ReadAll(c.Request().Body)
					if err != nil {
						return err
					}
					
					// Restore body for downstream handlers
					c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
					
					// Parse JSON
					var data map[string]interface{}
					if err := json.Unmarshal(bodyBytes, &data); err == nil {
						// Sanitize string fields
						sanitized := sanitizeMap(data)
						
						// Marshal back to JSON
						sanitizedBytes, err := json.Marshal(sanitized)
						if err == nil {
							c.Request().Body = io.NopCloser(bytes.NewBuffer(sanitizedBytes))
						}
					}
				}
			}
			
			return next(c)
		}
	}
}

// sanitizeMap recursively sanitizes all string values in a map
func sanitizeMap(data map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	
	for key, value := range data {
		switch v := value.(type) {
		case string:
			// Don't sanitize password field
			if key == "password" || key == "refreshToken" {
				result[key] = v
			} else {
				result[key] = helper.SanitizeString(v)
			}
		case map[string]interface{}:
			result[key] = sanitizeMap(v)
		case []interface{}:
			result[key] = sanitizeSlice(v)
		default:
			result[key] = v
		}
	}
	
	return result
}

// sanitizeSlice recursively sanitizes all string values in a slice
func sanitizeSlice(data []interface{}) []interface{} {
	result := make([]interface{}, len(data))
	
	for i, value := range data {
		switch v := value.(type) {
		case string:
			result[i] = helper.SanitizeString(v)
		case map[string]interface{}:
			result[i] = sanitizeMap(v)
		case []interface{}:
			result[i] = sanitizeSlice(v)
		default:
			result[i] = v
		}
	}
	
	return result
}

