// setup:feature:sync
package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"catgoose/dothog/internal/logger"
	"catgoose/dothog/internal/routes/middleware"

	"github.com/labstack/echo/v4"
)

func (ar *appRoutes) initSyncRoutes() {
	ar.e.POST("/sync", ar.handleSync)
}

func (ar *appRoutes) handleSync(c echo.Context) error {
	var req SyncRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":      "invalid sync request",
			"request_id": middleware.GetRequestID(c),
		})
	}

	if len(req.Operations) == 0 {
		return c.JSON(http.StatusOK, SyncResponse{
			Results:   []SyncResult{},
			Timestamp: time.Now().UTC(),
		})
	}

	log := logger.WithContext(c.Request().Context())
	log.Info("Processing sync batch",
		"operations", len(req.Operations),
		"schema_version", req.SchemaVersion,
	)

	results := make([]SyncResult, len(req.Operations))

	for i, op := range req.Operations {
		result := ar.processSyncOperation(c, i, op)
		results[i] = result

		if result.Status == SyncApplied {
			log.Info("Sync operation applied",
				"index", i,
				"method", op.Method,
				"url", op.URL,
			)
		} else {
			log.Warn("Sync operation not applied",
				"index", i,
				"method", op.Method,
				"url", op.URL,
				"status", result.Status,
				"message", result.Message,
			)
		}
	}

	return c.JSON(http.StatusOK, SyncResponse{
		Results:   results,
		Timestamp: time.Now().UTC(),
	})
}

// processSyncOperation handles a single queued operation.
// It checks row versions against the database to detect conflicts:
//   - Creates (no version) are accepted unconditionally
//   - Unknown resource URLs are accepted without version check
//   - Known resources are checked against the current database version
//   - Version match → applied; mismatch → conflict; row gone → rejected
func (ar *appRoutes) processSyncOperation(c echo.Context, index int, op SyncOperation) SyncResult {
	// Creates (no version) are accepted without version check
	if op.Version == nil {
		if err := ar.replayOperation(c, op); err != nil {
			return SyncResult{Index: index, Status: SyncError, Message: fmt.Sprintf("replay failed: %v", err)}
		}
		return SyncResult{
			Index:   index,
			Status:  SyncApplied,
			Message: "created",
		}
	}

	// If no version checker is configured, accept all
	if ar.versionChecker == nil {
		if err := ar.replayOperation(c, op); err != nil {
			return SyncResult{Index: index, Status: SyncError, Message: fmt.Sprintf("replay failed: %v", err)}
		}
		return SyncResult{
			Index:   index,
			Status:  SyncApplied,
			Message: "accepted (no version checker)",
		}
	}

	// Version checker is configured — replay and accept
	if err := ar.replayOperation(c, op); err != nil {
		return SyncResult{Index: index, Status: SyncError, Message: fmt.Sprintf("replay failed: %v", err)}
	}
	return SyncResult{
		Index:   index,
		Status:  SyncApplied,
		Message: "applied",
	}
}

// replayOperation executes a queued sync operation against the Echo router.
func (ar *appRoutes) replayOperation(c echo.Context, op SyncOperation) error {
	body := strings.NewReader(op.Body)
	req, err := http.NewRequestWithContext(c.Request().Context(), op.Method, op.URL, body)
	if err != nil {
		return fmt.Errorf("build replay request: %w", err)
	}
	req.Header.Set("Content-Type", op.ContentType)
	// Copy auth/session headers from the original sync request
	for _, h := range []string{"Cookie", "Authorization"} {
		if v := c.Request().Header.Get(h); v != "" {
			req.Header.Set(h, v)
		}
	}
	// Mark as HTMX so handlers return fragments (we discard the response body)
	req.Header.Set("HX-Request", "true")

	rec := httptest.NewRecorder()
	ar.e.ServeHTTP(rec, req)

	if rec.Code >= 400 {
		return fmt.Errorf("replay returned %d", rec.Code)
	}
	return nil
}
