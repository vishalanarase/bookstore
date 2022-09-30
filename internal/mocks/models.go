package mocks

import "github.com/vishalanarase/bookstore/internal/data"

func NewMockModels() *data.Models {
	return &data.Models{
		Book: NewBookMockModel(),
	}
}
