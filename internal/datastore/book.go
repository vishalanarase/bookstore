package datastore

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishalanarase/bookstore/pkg/errors"
	"gorm.io/gorm"
)

// BookInterface represents a Book interface
type BookInterface interface {
	GetDatabaseObject() (*gorm.DB, *errors.APIError)
	List(ctx *gin.Context) ([]Book, *errors.APIError)
	Get(ctx *gin.Context, uuid string) (Book, *errors.APIError)
	Create(ctx *gin.Context, book Book) (Book, *errors.APIError)
	Delete(ctx *gin.Context, uuid string) *errors.APIError
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
func (b *BookRepo) GetDatabaseObject() (*gorm.DB, *errors.APIError) {
	if b.DB != nil {
		return b.DB, nil
	}
	return nil, &errors.APIError{Status: http.StatusNotFound, Message: "Database object not found"}
}

// List return books from database or return error
func (b *BookRepo) List(ctx *gin.Context) ([]Book, *errors.APIError) {
	books := []Book{}
	err := &errors.APIError{}
	result := b.DB.Find(&books)
	if result.Error != nil {
		err.Status = http.StatusInternalServerError
		err.Message = result.Error.Error()
		return books, err
	}

	return books, nil
}

// Get returns a book from database or returns an error
func (b *BookRepo) Get(ctx *gin.Context, uuid string) (Book, *errors.APIError) {
	var book Book
	err := &errors.APIError{}
	result := b.DB.Where("id = ?", uuid).Where("deleted_at IS NULL").First(&book)
	if result.Error != nil {
		err.Status = http.StatusInternalServerError
		err.Message = result.Error.Error()
		return book, err
	}

	return book, nil
}

// Create creates a new Book into the database and returns a new Book or an error
func (b *BookRepo) Create(ctx *gin.Context, book Book) (Book, *errors.APIError) {
	result := b.DB.Create(&book)
	err := &errors.APIError{}
	if result.Error != nil {
		err.Status = http.StatusInternalServerError
		err.Message = result.Error.Error()
		return book, err
	}

	return book, nil
}

func (b *BookRepo) Delete(ctx *gin.Context, uuid string) *errors.APIError {
	return nil
}
