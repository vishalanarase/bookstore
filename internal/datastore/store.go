package datastore

import "gorm.io/gorm"

// Store represents a datastore for interacting with the database
type Store struct {
	Book   BookInterface
	Login  LoginInterface
	Rating RatingInterface
}

// NewStore creates a new Store instance
func NewStore(db *gorm.DB) *Store {
	return &Store{
		Book:   NewBookStore(db),
		Login:  NewLoginStore(db),
		Rating: NewRatingStore(db),
	}
}
