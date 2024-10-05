package login

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vishalanarase/bookstore/internal/datastore"
)

// ILoginController is a controller interface for LoginController
type ILoginController interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
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
	var user datastore.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	login, err := ctrl.store.Login.Login(ctx, user)
	if err != nil {
		log.WithError(err).Error("Failed to login")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in login"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": login.Key})
}

// Logout
func (ctrl *LoginController) Logout(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if strings.Contains(tokenString, "Bearer") {
		tokenString = tokenString[len("Bearer "):]
	}
	err := ctrl.store.Login.Logout(ctx, tokenString)
	if err != nil {
		log.WithError(err).Error("Failed to logout")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in logout"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
