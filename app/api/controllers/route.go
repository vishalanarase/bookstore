package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vishalanarase/bookstore/app/api/controllers/ping"
)

// AddRoutes will add all the routes to the router
func AddRoutes(router *gin.Engine) {
	v1Routes := router.Group("/v1")
	v1Routes.Use()
	{
		v1Routes.GET("/ping", ping.Ping)
	}
}