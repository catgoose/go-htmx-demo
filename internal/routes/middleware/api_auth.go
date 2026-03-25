package middleware

import (
	"crypto/subtle"
	"net/http"

	"github.com/labstack/echo/v4"
)

// InternalAPIAuth returns middleware that validates requests against a
// pre-shared API key. The key is checked in the X-API-Key header.
// If apiKey is empty, the middleware rejects all requests (API disabled).
func InternalAPIAuth(apiKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if apiKey == "" {
				return c.JSON(http.StatusServiceUnavailable, map[string]string{
					"error": "internal API not configured",
				})
			}

			provided := c.Request().Header.Get("X-API-Key")
			if subtle.ConstantTimeCompare([]byte(provided), []byte(apiKey)) != 1 {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid or missing API key",
				})
			}

			return next(c)
		}
	}
}
