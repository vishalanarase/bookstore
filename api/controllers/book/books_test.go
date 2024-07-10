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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vishalanarase/bookstore/internal/datastore"
	"github.com/vishalanarase/bookstore/internal/mockdatastore"
)

// BookResponse represents a book response
type BookResponse struct {
	code int
	body datastore.Book
}

// TestMain sets the test mode for testing
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

// TestCreate creates a new book
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
				body: datastore.Book{
					ID:     "11",
					Title:  "Test",
					Rating: 4,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mm := mockdatastore.NewMockStore()
			mockBook := new(mockdatastore.MockBook)
			mm.Book = mockBook

			mockBook.On("Create", mock.AnythingOfType("*gin.Context"), test.resp.body).Return(test.resp.body, nil)

			rr := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(test.payload))
			g.Expect(err).To(BeNil())

			ctrl := NewBookController(mm)
			router := gin.Default()
			router.POST("/v1/books", ctrl.Create)
			router.ServeHTTP(rr, request)

			g.Expect(test.resp.code).To(Equal(rr.Code))

			respBody := datastore.Book{}
			err = json.Unmarshal(rr.Body.Bytes(), &respBody)
			g.Expect(err).To(BeNil())

			g.Expect(test.resp.body).To(Equal(respBody))
		})
	}
}

// TestGet gets the book
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
				body: datastore.Book{
					ID:     "1",
					Title:  "Test",
					Rating: 4,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mm := mockdatastore.NewMockStore()
			mockBook := new(mockdatastore.MockBook)
			mm.Book = mockBook

			mockBook.On("Get", mock.AnythingOfType("*gin.Context"), test.payload).Return(test.resp.body, nil)

			rr := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, "/v1/books/"+test.payload, nil)
			g.Expect(err).To(BeNil())

			ctrl := NewBookController(mm)
			router := gin.Default()
			router.GET("/v1/books/:id", ctrl.Get)
			router.ServeHTTP(rr, request)

			g.Expect(test.resp.code).To(Equal(rr.Code))

			respBody := datastore.Book{}
			err = json.Unmarshal(rr.Body.Bytes(), &respBody)
			g.Expect(err).To(BeNil())

			g.Expect(test.resp.body).To(Equal(respBody))

		})
	}
}

// TestBookModel test the book model
func TestBookModel(t *testing.T) {
	mockDB := new(mockdatastore.MockBook)
	// Optionally, set expectations on the mock methods

	t.Run("List", func(t *testing.T) {
		// Define expected return values and errors
		expectedBooks := []datastore.Book{
			{ID: "1", Title: "Book 1", Rating: 5, Author: "Author 1"},
			{ID: "2", Title: "Book 2", Rating: 4, Author: "Author 2"},
		}
		mockDB.On("List", mock.Anything).Return(expectedBooks, nil)

		// Call the List method of the mock
		books, err := mockDB.List(&gin.Context{})

		// Assert the result
		assert.NoError(t, err)
		assert.Equal(t, expectedBooks, books)

		// Optionally, verify that the expected method was called
		mockDB.AssertExpectations(t)
	})

	t.Run("Get", func(t *testing.T) {
		// Define expected return values and errors
		expectedBook := datastore.Book{ID: "1", Title: "Book 1", Rating: 5, Author: "Author 1"}
		mockDB.On("Get", mock.Anything, "1").Return(expectedBook, nil)

		// Call the Get method of the mock
		book, err := mockDB.Get(&gin.Context{}, "1")

		// Assert the result
		assert.NoError(t, err)
		assert.Equal(t, expectedBook, book)

		// Optionally, verify that the expected method was called
		mockDB.AssertExpectations(t)
	})

	t.Run("Create", func(t *testing.T) {
		// Define expected return values and errors
		newBook := datastore.Book{Title: "New Book", Rating: 4, Author: "Author 3"}
		mockDB.On("Create", mock.Anything, newBook).Return(newBook, nil)

		// Call the Create method of the mock
		createdBook, err := mockDB.Create(&gin.Context{}, newBook)

		// Assert the result
		assert.NoError(t, err)
		assert.Equal(t, newBook, createdBook)

		// Optionally, verify that the expected method was called
		mockDB.AssertExpectations(t)
	})
}
