package mockdatastore

import "github.com/vishalanarase/bookstore/internal/datastore"

// NewMockStore returns a new mockstore
func NewMockStore() *datastore.Store {
	return &datastore.Store{
		Book: NewBookMockStore(),
	}
}
