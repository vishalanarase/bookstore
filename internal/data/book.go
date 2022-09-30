package data

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type BookInterface interface {
	List(ctx *gin.Context) ([]Book, error)
	GetBook(ctx *gin.Context, uuid string) (Book, error)
}

type Book struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BookModel struct{}

func NewBookModel() *BookModel {
	return &BookModel{}
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
