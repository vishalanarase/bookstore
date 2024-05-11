package ping

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

func TestPing(t *testing.T) {
	router := gin.Default()
	router.GET("/v1/ping", Ping)

	rr := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/v1/ping", nil)
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
}
