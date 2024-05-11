package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishalanarase/bookstore/api/controllers/book"
	"github.com/vishalanarase/bookstore/api/controllers/ping"
	"github.com/vishalanarase/bookstore/internal/datastore"
	"github.com/vishalanarase/bookstore/pkg/metrics"
)

// AddRoutes will add all the routes to the router
func AddRoutes(router *gin.Engine, m *metrics.Metrics, dbm *datastore.Store) {
	v1Routes := router.Group("/v1")
	pctrl := ping.NewPingController()
	bctrl := book.NewBookController(dbm)

	v1Routes.Use()
	{
		v1Routes.GET("/metrics", wrapHandler(m.Handler()))
		v1Routes.GET("/ping", pctrl.Ping)
		v1Routes.GET("/books", bctrl.List)
		v1Routes.GET("/books/:id", bctrl.Get)
		v1Routes.POST("/books", bctrl.Create)
	}
}

// wrapHandler wraps an http.Handler into a gin.HandlerFunc
func wrapHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
