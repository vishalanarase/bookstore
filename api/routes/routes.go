package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishalanarase/bookstore/api/controllers/book"
	"github.com/vishalanarase/bookstore/api/controllers/ping"
	"github.com/vishalanarase/bookstore/internal/datastore"
)

// AddRoutes will add all the routes to the router
func AddRoutes(router *gin.Engine, dbm *datastore.Store) {
	v1Routes := router.Group("/v1")
	pctrl := ping.NewPingController()
	bctrl := book.NewBookController(dbm)

	v1Routes.Use()
	{
		v1Routes.GET("/ping", pctrl.Ping)
		v1Routes.GET("/books", bctrl.List)
		v1Routes.GET("/books/:id", bctrl.Get)
		v1Routes.POST("/books", bctrl.Create)
	}
}
