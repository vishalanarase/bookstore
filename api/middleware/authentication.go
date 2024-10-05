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

		if token.IsTokenBlacklisted(tokenString) {
			http.Error(c.Writer, "Unauthorized: Token is blacklisted", http.StatusUnauthorized)
			c.Abort()
			return
		}

		// Verify the token
		_, err := token.VerifyToken(c, tokenString)
		if err != nil {
			http.Error(c.Writer, "Unauthorized: Token is not valid", http.StatusUnauthorized)
			c.Abort()
			return
		}
	}
	// Continue with the next middleware or route handler

	c.Next()
}

// AdminMiddleware is a middleware for authorizing admins
func AdminMiddleware(c *gin.Context) {
	// Check role is set to admin
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admins only"})
		c.Abort()
		return
	}
	c.Next()
}
