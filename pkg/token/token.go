package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Add a new global variable for the secret key
var jwtKey = []byte("TopSecretKey")

func GenerateToken(username string, roles []string) (string, error) {
	mapClaims := &jwt.MapClaims{
		"sub":  username,
		"iss":  "bookstore",
		"role": roles,
		"exp":  time.Now().Add(time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return claims.SignedString(jwtKey)
}

// Function to verify JWT tokens
func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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

	// Return the verified token
	return token, nil
}
