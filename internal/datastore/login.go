package datastore

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vishalanarase/bookstore/pkg/errors"
	"github.com/vishalanarase/bookstore/pkg/token"

	"gorm.io/gorm"
)

// LoginInterface represents a Login interface
type LoginInterface interface {
	Login(ctx *gin.Context, user User) (Login, error)
	Logout(ctx *gin.Context, token string) error
}

// Login represents a Login
type Login struct {
	Key string `json:"key,omitempty"`
}

// LoginRepo represents a Login
type LoginRepo struct {
	DB *gorm.DB
}

// NewLoginStore creates a new LoginStore instance
func NewLoginStore(db *gorm.DB) LoginInterface {
	return &LoginRepo{
		DB: db,
	}
}

func (lb *LoginRepo) Login(ctx *gin.Context, u User) (Login, error) {
	var user User
	result := lb.DB.Where("username = ?", u.Username).First(&user)
	if result.Error != nil {
		log.WithError(result.Error).Error("Unable to fect record")
		errors.NewAPIError(http.StatusInternalServerError, "Failed to login")
	}

	if user.Username == u.Username && user.Password == u.Password {
		key, err := token.GenerateToken(u.Username, user.Role)
		if err != nil {
			return Login{}, err
		}

		return Login{Key: key}, nil
	}

	return Login{}, errors.NewAPIError(http.StatusUnauthorized, "Invalid username or password")
}

func (lb *LoginRepo) Logout(ctx *gin.Context, tokenString string) error {
	token.BlacklistToken(tokenString)
	return nil
}
