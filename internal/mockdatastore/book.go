package mockdatastore

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/vishalanarase/bookstore/internal/datastore"
	"gorm.io/gorm"
)

type MockBook struct {
	mock.Mock
}

func NewBookMockStore() datastore.BookInterface {
	return &MockBook{
		Mock: mock.Mock{},
	}
}

func (m *MockBook) GetDatabaseObject() (*gorm.DB, error) {
	args := m.Called()
	return args.Get(0).(*gorm.DB), args.Error(1)
}

func (m *MockBook) List(ctx *gin.Context) ([]datastore.Book, error) {
	args := m.Called(ctx)
	return args.Get(0).([]datastore.Book), args.Error(1)
}

func (m *MockBook) Get(ctx *gin.Context, uuid string) (datastore.Book, error) {
	args := m.Called(ctx, uuid)
	return args.Get(0).(datastore.Book), args.Error(1)
}

func (m *MockBook) Create(ctx *gin.Context, book datastore.Book) (datastore.Book, error) {
	args := m.Called(ctx, book)
	return args.Get(0).(datastore.Book), args.Error(1)
}
