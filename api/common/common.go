package common

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/vishalanarase/bookstore/internal/datastore"
)

func GetModelsFromContext(ctx *gin.Context) (*datastore.Models, error) {
	models, ok := ctx.Get("Models")
	if !ok {
		return nil, errors.New("failed to get models handler")
	}

	db, ok := models.(datastore.Models)
	if !ok {
		return nil, errors.New("failed to cast data.Models")
	}

	return &db, nil
}
