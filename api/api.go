package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vishalanarase/bookstore/api/middleware"
	"github.com/vishalanarase/bookstore/api/routes"
	"github.com/vishalanarase/bookstore/internal/configs"
	"github.com/vishalanarase/bookstore/internal/datastore"
)

type Application struct {
	Server *gin.Engine
}

// NewApplication returns a new Application
func NewApplication() *Application {
	return &Application{
		// Create new engine instance
		Server: gin.New(),
	}
}

// Start starts the application
func (app *Application) Start(envConfig configs.GlobalConfig) {
	log.Info("Starting bookstore app")

	// Get database connection
	db, err := configs.DatabaseConnection(envConfig)
	if err != nil {
		log.Fatal(err, "Failed to connect to database")
	}

	// Set the mode
	//gin.SetMode(gin.ReleaseMode)

	// Use the logging
	app.Server.Use(gin.Logger())

	// Recover from panic
	app.Server.Use(gin.Recovery())

	// Rate limit api
	app.Server.Use(middleware.RateLimitHandler)

	// Log the request
	app.Server.Use(middleware.LogHandler)

	// Register the routes
	routes.AddRoutes(app.Server, datastore.NewStore(db))

	// Start the api
	err = app.Server.Run(":8080")
	if err != nil {
		log.Fatal(err, "Failed to start the gin server")
	}
}
