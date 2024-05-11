package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vishalanarase/bookstore/internal/datastore"
)

// Models is a middleware provide Models handler
func Models(models datastore.Models) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Models", models)
		c.Next()
	}
}
