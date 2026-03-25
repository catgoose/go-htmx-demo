package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInternalAPIAuth_ValidKey(t *testing.T) {
	e := echo.New()
	mw := InternalAPIAuth("test-secret-key")
	handler := mw(func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/api/test", http.NoBody)
	req.Header.Set("X-API-Key", "test-secret-key")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestInternalAPIAuth_MissingKey(t *testing.T) {
	e := echo.New()
	mw := InternalAPIAuth("test-secret-key")
	handler := mw(func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/api/test", http.NoBody)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestInternalAPIAuth_WrongKey(t *testing.T) {
	e := echo.New()
	mw := InternalAPIAuth("test-secret-key")
	handler := mw(func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/api/test", http.NoBody)
	req.Header.Set("X-API-Key", "wrong-key")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestInternalAPIAuth_EmptyConfigKey(t *testing.T) {
	e := echo.New()
	mw := InternalAPIAuth("")
	handler := mw(func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/api/test", http.NoBody)
	req.Header.Set("X-API-Key", "anything")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusServiceUnavailable, rec.Code)
}
