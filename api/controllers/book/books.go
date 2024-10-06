package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vishalanarase/bookstore/internal/datastore"
)

// IBookController is a controller interface for BookController
type IBookController interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

// BookController represents a BookController structure
type BookController struct {
	store *datastore.Store
}

// NewBookController initialise a new BookController and returns
func NewBookController(dbm *datastore.Store) IBookController {
	return &BookController{
		store: dbm,
	}
}

// List returns a list of books
func (ctrl *BookController) List(ctx *gin.Context) {
	books, err := ctrl.store.Book.List(ctx)
	if err != nil {
		log.WithError(err).Error("Failed to list books")

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in Listing Books"})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

// Get returns the book
func (ctrl *BookController) Get(ctx *gin.Context) {
	uuid := ctx.Param("id")

	book, err := ctrl.store.Book.Get(ctx, uuid)
	if err != nil {
		log.WithError(err).Error("Failed to get book")

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in getting Books"})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// Create creates a new book
func (ctrl *BookController) Create(ctx *gin.Context) {
	book := datastore.Book{}

	if err := ctx.ShouldBindJSON(&book); err != nil {
		log.WithError(err).Error("Failed to bind")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad book data"})
		return
	}

	book, err := ctrl.store.Book.Create(ctx, book)
	if err != nil {
		log.WithError(err).Error("Failed to create")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in inserting Books"})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

// Delete a book
func (ctrl *BookController) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, nil)
}
