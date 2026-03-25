package routes

import (
	"net/http"
	"strconv"
	"time"

	"catgoose/dothog/internal/routes/middleware"

	"github.com/catgoose/promolog"
	"github.com/labstack/echo/v4"
)

// APITraceSummary is the JSON representation of a trace for the internal API.
type APITraceSummary struct {
	RequestID  string    `json:"request_id"`
	CreatedAt  time.Time `json:"created_at"`
	StatusCode int       `json:"status_code"`
	Method     string    `json:"method"`
	Route      string    `json:"route"`
	ErrorChain string    `json:"error_chain"`
	RemoteIP   string    `json:"remote_ip"`
	UserID     string    `json:"user_id,omitempty"`
}

// APITraceListResponse is the response envelope for the trace list endpoint.
type APITraceListResponse struct {
	Traces []APITraceSummary `json:"traces"`
	Total  int               `json:"total"`
}

func (ar *appRoutes) initAPIRoutes(apiKey string) {
	if ar.reqLogStore == nil {
		return
	}

	api := ar.e.Group("/api", middleware.InternalAPIAuth(apiKey))
	api.GET("/error-traces", ar.handleAPIErrorTraces)
}

func (ar *appRoutes) handleAPIErrorTraces(c echo.Context) error {
	limit := 50
	if l, err := strconv.Atoi(c.QueryParam("limit")); err == nil && l > 0 && l <= 500 {
		limit = l
	}

	f := promolog.TraceFilter{
		Q:       c.QueryParam("q"),
		Status:  c.QueryParam("status"),
		Method:  c.QueryParam("method"),
		Sort:    "CreatedAt",
		Dir:     "desc",
		Page:    1,
		PerPage: limit,
	}

	// Optional "since" filter — ISO 8601 timestamp
	if since := c.QueryParam("since"); since != "" {
		// Pass through as search query — the store's search covers created_at.
		// For proper date filtering, the store would need a dedicated field.
		// For now, this endpoint returns the most recent traces up to limit.
		_ = since // TODO: add date range filtering to promolog.TraceFilter
	}

	traces, total, err := ar.reqLogStore.ListTraces(c.Request().Context(), f)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to query traces",
		})
	}

	apiTraces := make([]APITraceSummary, len(traces))
	for i, t := range traces {
		apiTraces[i] = APITraceSummary{
			RequestID:  t.RequestID,
			CreatedAt:  t.CreatedAt,
			StatusCode: t.StatusCode,
			Method:     t.Method,
			Route:      t.Route,
			ErrorChain: t.ErrorChain,
			RemoteIP:   t.RemoteIP,
			UserID:     t.UserID,
		}
	}

	return c.JSON(http.StatusOK, APITraceListResponse{
		Traces: apiTraces,
		Total:  total,
	})
}
