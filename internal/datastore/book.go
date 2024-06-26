package datastore

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// BookInterface represents a Book interface
type BookInterface interface {
	List(ctx *gin.Context) ([]Book, error)
	Get(ctx *gin.Context, uuid string) (Book, error)
	Create(ctx *gin.Context, book Book) (Book, error)
	GetDatabaseObject() (*gorm.DB, error)
}

// Book represents a Book
type Book struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Rating     int    `json:"rating"`
	Authorname string `json:"authorname"`
}

// BookRepo represents a Book
type BookRepo struct {
	DB *gorm.DB
}

// NewBookStore creates a new BookStore instance
func NewBookStore(db *gorm.DB) BookInterface {
	return &BookRepo{
		DB: db,
	}
}

// GetDatabaseObject return the databaseobject is set or return an error
func (b *BookRepo) GetDatabaseObject() (*gorm.DB, error) {
	if b.DB != nil {
		return b.DB, nil
	}
	return nil, errors.New("databse not initialized")
}

// List return books from database or return error
func (b *BookRepo) List(ctx *gin.Context) ([]Book, error) {
	books := []Book{}

	result := b.DB.Find(&books)
	if result.Error != nil {
		err := fmt.Errorf("failed: %v", result.Error)
		return books, err
	}

	return books, nil
}

// Get returns a book from database or returns an error
func (b *BookRepo) Get(ctx *gin.Context, uuid string) (Book, error) {
	var book Book

	result := b.DB.Where("id = ?", uuid).Where("deleted_at IS NULL").First(&book)
	if result.Error != nil {
		err := fmt.Errorf("failed: %v", result.Error)
		return book, err
	}

	return book, nil
}

// Create creates a new Book into the database and returns a new Book or an error
func (b *BookRepo) Create(ctx *gin.Context, book Book) (Book, error) {
	result := b.DB.Create(&book)
	if result.Error != nil {
		err := fmt.Errorf("failed: %v", result.Error)
		return book, err
	}

	return book, nil
}
