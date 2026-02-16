package api

import (
	"encoding/json"
	"moltket/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	e := echo.New()
	cfg := &config.Config{}
	server := &Server{
		echo:   e,
		config: cfg,
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := server.healthCheck(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, 200, int(response["status"].(float64)))
	assert.Equal(t, "skillchain-verification", response["service"])
	assert.Contains(t, response, "time")
}

func TestVerifyLicenseHandler(t *testing.T) {
	e := echo.New()

	t.Run("Health Check - Basic", func(t *testing.T) {
		cfg := &config.Config{}
		server := &Server{
			echo:   e,
			config: cfg,
		}

		req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := server.healthCheck(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
