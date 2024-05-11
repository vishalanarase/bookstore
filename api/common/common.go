package common

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/vishalanarase/bookstore/internal/data"
)

func GetModelsFromContext(ctx *gin.Context) (*data.Models, error) {
	models, ok := ctx.Get("Models")
	if !ok {
		return nil, errors.New("failed to get models handler")
	}

	db, ok := models.(data.Models)
	if !ok {
		return nil, errors.New("failed to cast data.Models")
	}

	return &db, nil
}
