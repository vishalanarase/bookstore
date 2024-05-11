package mockdatastore

import "github.com/vishalanarase/bookstore/internal/datastore"

func NewMockModels() *datastore.Models {
	return &datastore.Models{
		Book: NewBookMockModel(),
	}
}
