package data

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookInterface interface {
	List(ctx *gin.Context) ([]Book, error)
	GetBook(ctx *gin.Context, uuid string) (Book, error)
}

type Book struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BookModel struct {
	DB *gorm.DB
}

func NewBookModel(db *gorm.DB) *BookModel {
	return &BookModel{
		DB: db,
	}
}

var books = []Book{
	{
		ID:   "1",
		Name: "One",
	},
	{
		ID:   "2",
		Name: "Two",
	},
}

func (b *BookModel) List(ctx *gin.Context) ([]Book, error) {
	var books []Book

	result := b.DB.Find(&books)
	if result.Error != nil {
		err := fmt.Errorf("failed: %v", result.Error)
		return books, err
	}

	return books, nil
}

func (b *BookModel) GetBook(ctx *gin.Context, uuid string) (Book, error) {
	for _, b := range books {
		if b.ID == uuid {
			return b, nil
		}
	}

	return Book{}, errors.New("not Found")
}
