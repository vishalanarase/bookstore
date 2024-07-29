package datastore

import (
	"github.com/gin-gonic/gin"
	"github.com/vishalanarase/bookstore/pkg/token"

	"gorm.io/gorm"
)

// LoginInterface represents a Login interface
type LoginInterface interface {
	Login(ctx *gin.Context) (Login, error)
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

func (lb *LoginRepo) Login(ctx *gin.Context) (Login, error) {
	key, err := token.GenerateToken("admin", []string{"admin"})
	if err != nil {
		return Login{}, err
	}

	return Login{Key: key}, nil
}
