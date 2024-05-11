package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vishalanarase/bookstore/api/common"
	"github.com/vishalanarase/bookstore/internal/datastore"
)

func List(ctx *gin.Context) {

	models, err := common.GetModelsFromContext(ctx)
	if err != nil {
		log.WithError(err).Error("Failed to get models from context")

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in getting Models"})
		return
	}

	books, err := models.Book.List(ctx)
	if err != nil {
		log.WithError(err).Error("Failed to list books")

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in Listing Books"})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func Get(ctx *gin.Context) {

	models, err := common.GetModelsFromContext(ctx)
	if err != nil {
		log.WithError(err).Error("Failed to get models from context")

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in getting Models"})
		return
	}

	uuid := ctx.Param("id")

	book, err := models.Book.Get(ctx, uuid)
	if err != nil {
		log.WithError(err).Error("Failed to get book")

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in getting Books"})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func Create(ctx *gin.Context) {
	models, err := common.GetModelsFromContext(ctx)
	if err != nil {
		log.WithError(err).Error("Failed to get models from context")

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in getting Models"})
		return
	}

	book := datastore.Book{}

	if err := ctx.ShouldBindJSON(&book); err != nil {
		log.WithError(err).Error("Failed to bind")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad book data"})
		return
	}

	book, err = models.Book.Create(ctx, book)
	if err != nil {
		log.WithError(err).Error("Failed to create")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in inserting Books"})
		return
	}

	ctx.JSON(http.StatusOK, book)
}
