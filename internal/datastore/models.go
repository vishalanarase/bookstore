package datastore

import "gorm.io/gorm"

type Models struct {
	Book BookInterface
}

func NewModels(db *gorm.DB) *Models {
	return &Models{
		Book: NewBookModel(db),
	}
}
