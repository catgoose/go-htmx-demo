// setup:feature:demo

package middleware

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"catgoose/dothog/internal/routes/hypermedia"

	"github.com/labstack/echo/v4"
)

// VisitRecorder records page visits for frecency tracking.
type VisitRecorder interface {
	RecordVisit(ctx context.Context, sessionID, path, title string) error
}

// FrecencyMiddleware records page visits and sets the session ID on the
// request context so the FrecencySource can read it downstream.
// Only records full page loads (not HTMX partial requests) for tracked routes.
func FrecencyMiddleware(recorder VisitRecorder, validRoutes map[string]bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL.Path

			// Set session ID on context for the FrecencySource to read.
			sessionID := deriveSessionID(c)
			if sessionID != "" {
				ctx := hypermedia.WithSessionID(c.Request().Context(), sessionID)
				c.SetRequest(c.Request().WithContext(ctx))
			}

			// Only record full page loads for tracked routes.
			isHTMX := c.Request().Header.Get("HX-Request") == "true"
			if !isHTMX && sessionID != "" && c.Request().Method == http.MethodGet && validRoutes[path] {
				// Record in a goroutine — don't block the response.
				// Use a detached context since the request context will be
				// cancelled after the response is sent.
				sid := sessionID
				p := path
				title := hypermedia.TitleFromPath(path)
				go func() {
					_ = recorder.RecordVisit(context.Background(), sid, p, title)
				}()
			}

			return next(c)
		}
	}
}

// deriveSessionID builds a stable, anonymous session identifier from the
// request. For the demo (no auth), this hashes IP + User-Agent.
func deriveSessionID(c echo.Context) string {
	ip := c.RealIP()
	if ip == "" {
		return ""
	}
	ua := c.Request().UserAgent()
	h := sha256.Sum256([]byte(ip + "|" + ua))
	return hex.EncodeToString(h[:8])
}
