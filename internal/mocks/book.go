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
	return []data.Book{}, nil
}

func (m *MockBook) Get(ctx *gin.Context, uuid string) (data.Book, error) {
	ret := m.Called(ctx, uuid)

	// first value passed to "Return"
	var r0 *data.Book
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).(*data.Book)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return *r0, r1
}

func (m *MockBook) Create(ctx *gin.Context, book data.Book) (data.Book, error) {
	return data.Book{}, nil
}
