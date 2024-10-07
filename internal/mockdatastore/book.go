package mockdatastore

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/vishalanarase/bookstore/internal/datastore"
	"github.com/vishalanarase/bookstore/pkg/errors"
	"gorm.io/gorm"
)

// MockBook represents a mock book
type MockBook struct {
	mock.Mock
}

// NewBookMockStore initializes a new MockBook and reruns
func NewBookMockStore() datastore.BookInterface {
	return &MockBook{
		Mock: mock.Mock{},
	}
}

// GetDatabase returns the database object or an error
func (m *MockBook) GetDatabaseObject() (*gorm.DB, *errors.APIError) {
	args := m.Called()
	return args.Get(0).(*gorm.DB), &errors.APIError{Message: args.Error(1).Error()}
}

// List returns the mock books from the mock store
func (m *MockBook) List(ctx *gin.Context) ([]datastore.Book, *errors.APIError) {
	args := m.Called(ctx)
	return args.Get(0).([]datastore.Book), &errors.APIError{Message: args.Error(1).Error()}
}

// Get returns a book from the mock strore
func (m *MockBook) Get(ctx *gin.Context, uuid string) (datastore.Book, *errors.APIError) {
	args := m.Called(ctx, uuid)
	return args.Get(0).(datastore.Book), &errors.APIError{Message: args.Error(1).Error()}
}

// Create creates a new book into the mock store
func (m *MockBook) Create(ctx *gin.Context, book datastore.Book) (datastore.Book, *errors.APIError) {
	args := m.Called(ctx, book)

	if args.Error(1) != nil {
		return args.Get(0).(datastore.Book), &errors.APIError{Message: args.Error(1).Error()}
	}

	return args.Get(0).(datastore.Book), nil
}

func (m *MockBook) Delete(ctx *gin.Context, uuid string) *errors.APIError {
	return nil
}

func (m *MockBook) Rate(ctx *gin.Context, rate datastore.Rating) *errors.APIError {
	return nil
}
