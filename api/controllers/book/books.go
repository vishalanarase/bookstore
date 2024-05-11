package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/vishalanarase/bookstore/api/common"
	"github.com/vishalanarase/bookstore/internal/datastore"
)

func List(ctx *gin.Context) {

	models, err := common.GetModelsFromContext(ctx)
	if err != nil {
		log.Error().
			Str("controller", "book").
			Str("method", "List").
			Err(err).
			Send()

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in getting Models"})
		return
	}

	books, err := models.Book.List(ctx)
	if err != nil {
		log.Error().
			Str("controller", "book").
			Str("method", "List").
			Err(err).
			Send()

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in Listing Books"})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func Get(ctx *gin.Context) {

	models, err := common.GetModelsFromContext(ctx)
	if err != nil {
		log.Error().
			Str("controller", "book").
			Str("method", "Get").
			Err(err).
			Send()

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in getting Models"})
		return
	}

	uuid := ctx.Param("id")

	book, err := models.Book.Get(ctx, uuid)
	if err != nil {
		log.Error().
			Str("controller", "book").
			Str("method", "Get").
			Err(err).
			Send()

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in getting Books"})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func Create(ctx *gin.Context) {
	loger := log.Error().
		Str("controller", "book").
		Str("method", "Create")
	models, err := common.GetModelsFromContext(ctx)
	if err != nil {
		loger.Err(err).Send()
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in getting Models"})
		return
	}

	book := datastore.Book{}

	if err := ctx.ShouldBindJSON(&book); err != nil {
		log.Error().
			Err(err).
			Send()
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad book data"})
		return
	}

	book, err = models.Book.Create(ctx, book)
	if err != nil {
		loger.Err(err).Send()
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in inserting Books"})
		return
	}

	ctx.JSON(http.StatusOK, book)
}
