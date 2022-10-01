package book

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vishalanarase/bookstore/api/middleware"
	"github.com/vishalanarase/bookstore/internal/data"
	"github.com/vishalanarase/bookstore/internal/mocks"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

// func TestList(t *testing.T) {
// 	router := gin.Default()
// 	router.GET("/v1/books", List)

// 	rr := httptest.NewRecorder()
// 	request, err := http.NewRequest(http.MethodGet, "/v1/books", nil)
// 	assert.NoError(t, err)

// 	router.ServeHTTP(rr, request)

// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rr.Code)
// }

func TestGet(t *testing.T) {
	mockBookResp := &data.Book{
		ID:   "1",
		Name: "Test",
	}

	mm := mocks.NewMockModels()
	mockBook := new(mocks.MockBook)
	mockBook.On("GetBook", mock.AnythingOfType("*gin.Context"), "1").Return(mockBookResp, nil)
	mm.Book = mockBook

	rr := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/v1/books/1", nil)
	assert.NoError(t, err)

	router := gin.Default()
	router.Use(middleware.Models(*mm))
	router.GET("/v1/books/:id", Get)

	router.ServeHTTP(rr, request)

	assert.Equal(t, http.StatusOK, rr.Code)

	respBody := data.Book{}
	err = json.Unmarshal(rr.Body.Bytes(), &respBody)
	assert.NoError(t, err)

	assert.Equal(t, mockBookResp, &respBody)
	mockBook.AssertExpectations(t)
}
