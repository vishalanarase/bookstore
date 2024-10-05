package token

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// Add a new global variable for the secret key
var jwtKey = []byte("TopSecretKey")

// tokenBlacklist is a map to store blacklisted tokens
var tokenBlacklist = make(map[string]bool)

// JWT Claims
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(username string, role string) (string, error) {
	claims := &Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			Issuer:    "bookstore",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Function to verify JWT tokens
func VerifyToken(c *gin.Context, tokenString string) (*jwt.Token, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Save username and role in context for later use
	c.Set("username", claims.Username)
	c.Set("role", claims.Role)

	// Return the verified token
	return token, nil
}

// BlacklistToken blacklist token
func BlacklistToken(token string) {
	fmt.Println("Blacklisting", token)
	tokenBlacklist[token] = true
}

// IsTokenBlacklisted returns true if token is blacklisted
func IsTokenBlacklisted(token string) bool {
	fmt.Println("Blacklisting: %v", tokenBlacklist)
	return tokenBlacklist[token]
}
