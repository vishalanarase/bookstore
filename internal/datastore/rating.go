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
	List(ctx *gin.Context) ([]Rating, *errors.APIError)
	Get(ctx *gin.Context, uuid string) (Rating, *errors.APIError)
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

// Create creates a new Rating
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

// List lists all Ratings
func (r *RatingRepo) List(ctx *gin.Context) ([]Rating, *errors.APIError) {
	ratings := []Rating{}
	userID, ok := ctx.Get("user_id")
	if !ok {
		return ratings, &errors.APIError{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		}
	}

	result := r.DB.Where("user_id = ?", userID).Find(&ratings)
	err := &errors.APIError{}
	if result.Error != nil {
		err.Status = http.StatusInternalServerError
		err.Message = result.Error.Error()
		return ratings, err
	}

	return ratings, nil
}

// Get a rating
func (r *RatingRepo) Get(ctx *gin.Context, uuid string) (Rating, *errors.APIError) {
	rating := Rating{}
	result := r.DB.Where("id = ?", uuid).First(&rating)
	err := &errors.APIError{}
	if result.Error != nil {
		err.Status = http.StatusInternalServerError
		err.Message = result.Error.Error()
		return rating, err
	}
	return rating, nil
}
