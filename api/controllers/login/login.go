package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vishalanarase/bookstore/internal/datastore"
)

// ILoginController is a controller interface for LoginController
type ILoginController interface {
	Login(ctx *gin.Context)
}

// LoginController represents a LoginController structure
type LoginController struct {
	store *datastore.Store
}

// NewBLoginController initialise a new LoginController and returns
func NewLoginController(dbm *datastore.Store) ILoginController {
	return &LoginController{
		store: dbm,
	}
}

// Login
func (ctrl *LoginController) Login(ctx *gin.Context) {
	login, err := ctrl.store.Login.Login(ctx)
	if err != nil {
		log.WithError(err).Error("Failed to list books")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in Listing Books"})
		return
	}

	ctx.JSON(http.StatusOK, login.Key)
}
