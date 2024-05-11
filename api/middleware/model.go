package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vishalanarase/bookstore/internal/data"
)

// Models is a middleware provide Models handler
func Models(models data.Models) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Models", models)
		c.Next()
	}
}
