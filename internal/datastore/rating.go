package datastore

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vishalanarase/bookstore/pkg/errors"
	"gorm.io/gorm"
)

// RatingInterface represents a Rating interface
type RatingInterface interface {
	Create(ctx *gin.Context, rate Rating) *errors.APIError
	//Get(ctx *gin.Context, rate Rating) *errors.APIError
	List(ctx *gin.Context, rate Rating) ([]Rating, *errors.APIError)
}

// RatingRepo represents a Rating
type RatingRepo struct {
	DB *gorm.DB
}

// NewRatingStore creates a new RatingStore instance
func NewRatingStore(db *gorm.DB) RatingInterface {
	return &RatingRepo{
		DB: db,
	}
}

func (r *RatingRepo) Create(ctx *gin.Context, rate Rating) *errors.APIError {
	rate.ID = uuid.New().String()
	result := r.DB.Create(&rate)
	err := &errors.APIError{}
	if result.Error != nil {
		err.Status = http.StatusInternalServerError
		err.Message = result.Error.Error()
		return err
	}

	return nil
}

func (r *RatingRepo) List(ctx *gin.Context, rate Rating) ([]Rating, *errors.APIError) {
	return []Rating{}, &errors.APIError{}
}
