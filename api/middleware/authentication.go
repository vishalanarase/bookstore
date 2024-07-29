package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vishalanarase/bookstore/pkg/token"
)

// AuthenticationMiddleware is a middleware for authenticating
func AuthenticationMiddleware(c *gin.Context) {

	if !strings.Contains(c.Request.URL.Path, "login") {
		// Extract JWT token from request headers
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(c.Writer, "Unauthorized: No token found", http.StatusUnauthorized)
			c.Abort()
			return
		}

		if strings.Contains(tokenString, "Bearer") {
			tokenString = tokenString[len("Bearer "):]
		}

		// Verify the token
		_, err := token.VerifyToken(tokenString)
		if err != nil {
			http.Error(c.Writer, "Unauthorized: Token is not valid", http.StatusUnauthorized)
			c.Abort()
			return
		}
	}
	// Continue with the next middleware or route handler
	c.Next()
}
