package book

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"github.com/vishalanarase/bookstore/api/middleware"
	"github.com/vishalanarase/bookstore/internal/data"
	"github.com/vishalanarase/bookstore/internal/mocks"
)

type BookResponse struct {
	code int
	body data.Book
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
func TestCreate(t *testing.T) {
	g := NewWithT(t)

	tests := map[string]struct {
		payload string
		resp    BookResponse
	}{
		"should return 200": {
			payload: `{"id":"11","name":"Test","rating":4}`,
			resp: BookResponse{
				code: http.StatusOK,
				body: data.Book{
					ID:     "11",
					Name:   "Test",
					Rating: 4,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mm := mocks.NewMockModels()
			mockBook := new(mocks.MockBook)
			mm.Book = mockBook

			mockBook.On("Create", mock.AnythingOfType("*gin.Context"), test.resp.body).Return(test.resp.body, nil)

			rr := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(test.payload))
			g.Expect(err).To(BeNil())

			router := gin.Default()
			router.Use(middleware.Models(*mm))
			router.POST("/v1/books", Create)
			router.ServeHTTP(rr, request)

			g.Expect(test.resp.code).To(Equal(rr.Code))

			respBody := data.Book{}
			err = json.Unmarshal(rr.Body.Bytes(), &respBody)
			g.Expect(err).To(BeNil())

			g.Expect(test.resp.body).To(Equal(respBody))
		})
	}
}

func TestGet(t *testing.T) {
	g := NewWithT(t)

	tests := map[string]struct {
		payload string
		resp    BookResponse
	}{
		"should return 200": {
			payload: "1",
			resp: BookResponse{
				code: http.StatusOK,
				body: data.Book{
					ID:     "1",
					Name:   "Test",
					Rating: 4,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mm := mocks.NewMockModels()
			mockBook := new(mocks.MockBook)
			mm.Book = mockBook

			mockBook.On("Get", mock.AnythingOfType("*gin.Context"), test.payload).Return(test.resp.body, nil)

			rr := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, "/v1/books/"+test.payload, nil)
			g.Expect(err).To(BeNil())

			router := gin.Default()
			router.Use(middleware.Models(*mm))
			router.GET("/v1/books/:id", Get)
			router.ServeHTTP(rr, request)

			g.Expect(test.resp.code).To(Equal(rr.Code))

			respBody := data.Book{}
			err = json.Unmarshal(rr.Body.Bytes(), &respBody)
			g.Expect(err).To(BeNil())

			g.Expect(test.resp.body).To(Equal(respBody))

		})
	}
}
