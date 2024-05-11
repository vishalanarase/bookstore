package datastore

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookInterface interface {
	List(ctx *gin.Context) ([]Book, error)
	Get(ctx *gin.Context, uuid string) (Book, error)
	Create(ctx *gin.Context, book Book) (Book, error)
	GetDatabaseObject() (*gorm.DB, error)
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

func NewBookStore(db *gorm.DB) BookInterface {
	return &BookModel{
		DB: db,
	}
}

func (b *BookModel) GetDatabaseObject() (*gorm.DB, error) {
	if b.DB != nil {
		return b.DB, nil
	}
	return nil, errors.New("databse not initialized")
}

func (b *BookModel) List(ctx *gin.Context) ([]Book, error) {
	books := []Book{}

	result := b.DB.Find(&books)
	if result.Error != nil {
		err := fmt.Errorf("failed: %v", result.Error)
		return books, err
	}

	return books, nil
}

func (b *BookModel) Get(ctx *gin.Context, uuid string) (Book, error) {
	var book Book

	result := b.DB.Where("id = ?", uuid).Where("deleted_at IS NULL").First(&book)
	if result.Error != nil {
		err := fmt.Errorf("failed: %v", result.Error)
		return book, err
	}

	return book, nil
}

func (b *BookModel) Create(ctx *gin.Context, book Book) (Book, error) {
	result := b.DB.Create(&book)
	if result.Error != nil {
		err := fmt.Errorf("failed: %v", result.Error)
		return book, err
	}

	return book, nil
}
