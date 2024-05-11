package datastore

import "gorm.io/gorm"

type Store struct {
	Book BookInterface
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		Book: NewBookStore(db),
	}
}
