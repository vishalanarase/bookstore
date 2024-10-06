package rating

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vishalanarase/bookstore/internal/datastore"
)

// IRatingController is a controller interface for RatingController
type IRatingController interface {
	Create(ctx *gin.Context)
}

// RatingController represents a RatingController structure
type RatingController struct {
	store *datastore.Store
}

// NewRatingController initialise a new RatingController and returns
func NewRatingController(dbm *datastore.Store) IRatingController {
	return &RatingController{
		store: dbm,
	}
}

// Rate a book
func (ctrl *RatingController) Create(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	rate := datastore.Rating{}
	if err := ctx.ShouldBindJSON(&rate); err != nil {
		log.WithError(err).Error("Failed to bind")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad rating data"})
	}

	rate.UserID = userID.(string)
	err := ctrl.store.Rating.Create(ctx, rate)
	if err != nil {
		log.WithError(err).Error("Failed to rate")
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "Accepted"})
}
