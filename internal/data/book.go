package data

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookInterface interface {
	List(ctx *gin.Context) ([]Book, error)
	GetBook(ctx *gin.Context, uuid string) (Book, error)
}

type Book struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Rating     int    `json:"rating"`
	Authorname string `json:"authorname"`
}

type BookModel struct {
	DB *gorm.DB
}

func NewBookModel(db *gorm.DB) *BookModel {
	return &BookModel{
		DB: db,
	}
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
	var book Book

	result := b.DB.Where("id = ?", uuid).Where("deleted_at IS NULL").First(&book)
	if result.Error != nil {
		err := fmt.Errorf("failed: %v", result.Error)
		return book, err
	}

	return book, nil
}
