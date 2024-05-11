package mockdatastore

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/vishalanarase/bookstore/internal/datastore"
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
func (m *MockBook) GetDatabaseObject() (*gorm.DB, error) {
	args := m.Called()
	return args.Get(0).(*gorm.DB), args.Error(1)
}

// List returns the mock books from the mock store
func (m *MockBook) List(ctx *gin.Context) ([]datastore.Book, error) {
	args := m.Called(ctx)
	return args.Get(0).([]datastore.Book), args.Error(1)
}

// Get returns a book from the mock strore
func (m *MockBook) Get(ctx *gin.Context, uuid string) (datastore.Book, error) {
	args := m.Called(ctx, uuid)
	return args.Get(0).(datastore.Book), args.Error(1)
}

// Create creates a new book into the mock store
func (m *MockBook) Create(ctx *gin.Context, book datastore.Book) (datastore.Book, error) {
	args := m.Called(ctx, book)
	return args.Get(0).(datastore.Book), args.Error(1)
}
