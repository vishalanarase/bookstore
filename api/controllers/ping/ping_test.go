package ping

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestMain sets the test mode for testing
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

// TestPing test the ping
func TestPing(t *testing.T) {
	ctrl := NewPingController()
	router := gin.Default()
	router.GET("/v1/ping", ctrl.Ping)

	rr := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/v1/ping", nil)
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
}
