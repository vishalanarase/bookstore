package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishalanarase/bookstore/api/controllers/book"
	"github.com/vishalanarase/bookstore/api/controllers/login"
	"github.com/vishalanarase/bookstore/api/controllers/ping"
	"github.com/vishalanarase/bookstore/api/middleware"
	"github.com/vishalanarase/bookstore/internal/datastore"
)

// AddRoutes will add all the routes to the router
func AddRoutes(router *gin.Engine, dbm *datastore.Store) {
	pctrl := ping.NewPingController()
	bctrl := book.NewBookController(dbm)
	lctrl := login.NewLoginController(dbm)

	// Public routes
	router.GET("/ping", pctrl.Ping)
	router.POST("/v1/login", lctrl.Login)
	//router.POST("/logout", logout)
	router.GET("/v1/books", bctrl.List) // Everyone can view books

	// Admin routes
	admin := router.Group("/v1/admin")
	admin.Use(middleware.AuthenticationMiddleware, middleware.AdminMiddleware)
	admin.POST("/books", bctrl.Create)

	// User routes
	user := router.Group("/v1/books/:id")
	user.Use(middleware.AuthenticationMiddleware)
	//user.POST("/rate", rateBook)
}
