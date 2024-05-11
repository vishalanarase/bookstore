package mockdatastore

import "github.com/vishalanarase/bookstore/internal/datastore"

func NewMockStore() *datastore.Store {
	return &datastore.Store{
		Book: NewBookMockStore(),
	}
}
