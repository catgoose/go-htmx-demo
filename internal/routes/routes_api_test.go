package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	dialect "github.com/catgoose/fraggle"
	"github.com/catgoose/promolog"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testAPIKey = "test-secret-key"

func newTestAPIRoutes(t *testing.T) (*echo.Echo, *appRoutes) {
	t.Helper()

	db, _, err := dialect.OpenSQLite(context.Background(), ":memory:")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	store := promolog.NewStore(db)
	require.NoError(t, store.InitSchema())

	e := echo.New()
	ar := &appRoutes{
		e:           e,
		ctx:         context.Background(),
		reqLogStore: store,
	}
	ar.initAPIRoutes(testAPIKey)
	return e, ar
}

func apiRequest(t *testing.T, e *echo.Echo, apiKey string, target string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, target, http.NoBody)
	if apiKey != "" {
		req.Header.Set("X-API-Key", apiKey)
	}
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	return rec
}

func TestAPIErrorTraces_ValidKey(t *testing.T) {
	e, _ := newTestAPIRoutes(t)
	rec := apiRequest(t, e, testAPIKey, "/api/error-traces")

	assert.Equal(t, http.StatusOK, rec.Code)

	var resp APITraceListResponse
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.NotNil(t, resp.Traces)
	assert.Equal(t, 0, resp.Total)
}

func TestAPIErrorTraces_MissingKey(t *testing.T) {
	e, _ := newTestAPIRoutes(t)
	rec := apiRequest(t, e, "", "/api/error-traces")

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestAPIErrorTraces_WrongKey(t *testing.T) {
	e, _ := newTestAPIRoutes(t)
	rec := apiRequest(t, e, "wrong-key", "/api/error-traces")

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestAPIErrorTraces_DisabledWhenNoKey(t *testing.T) {
	db, _, err := dialect.OpenSQLite(context.Background(), ":memory:")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	store := promolog.NewStore(db)
	require.NoError(t, store.InitSchema())

	e := echo.New()
	ar := &appRoutes{
		e:           e,
		ctx:         context.Background(),
		reqLogStore: store,
	}
	ar.initAPIRoutes("") // empty key = disabled

	rec := apiRequest(t, e, "anything", "/api/error-traces")
	assert.Equal(t, http.StatusServiceUnavailable, rec.Code)
}

func TestAPIErrorTraces_LimitParam(t *testing.T) {
	e, ar := newTestAPIRoutes(t)

	// Seed a few traces
	ctx := context.Background()
	for i := 0; i < 5; i++ {
		ar.reqLogStore.Promote(ctx, promolog.ErrorTrace{
			RequestID:  "test-" + string(rune('a'+i)),
			StatusCode: 500,
			Method:     "GET",
			Route:      "/test",
			ErrorChain: "test error",
		})
	}

	rec := apiRequest(t, e, testAPIKey, "/api/error-traces?limit=2")
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp APITraceListResponse
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.Len(t, resp.Traces, 2)
	assert.Equal(t, 5, resp.Total)
}
