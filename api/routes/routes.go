package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishalanarase/bookstore/api/controllers/book"
	"github.com/vishalanarase/bookstore/api/controllers/ping"
)

// AddRoutes will add all the routes to the router
func AddRoutes(router *gin.Engine) {
	v1Routes := router.Group("/v1")
	v1Routes.Use()
	{
		v1Routes.GET("/ping", ping.Ping)
		v1Routes.GET("/books", book.List)
		v1Routes.GET("/books/:id", book.Get)
		v1Routes.POST("/books", book.Create)
	}
}
