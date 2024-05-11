package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/vishalanarase/bookstore/internal/data"
)

type MockBook struct {
	mock.Mock
}

func NewBookMockModel() *MockBook {
	return &MockBook{
		Mock: mock.Mock{},
	}
}

func (m *MockBook) List(ctx *gin.Context) ([]data.Book, error) {
	args := m.Called(ctx)
	return args.Get(0).([]data.Book), args.Error(1)
}

func (m *MockBook) Get(ctx *gin.Context, uuid string) (data.Book, error) {
	args := m.Called(ctx, uuid)
	return args.Get(0).(data.Book), args.Error(1)
}

func (m *MockBook) Create(ctx *gin.Context, book data.Book) (data.Book, error) {
	args := m.Called(ctx, book)
	return args.Get(0).(data.Book), args.Error(1)
}
