package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vishalanarase/bookstore/api/middleware"
	"github.com/vishalanarase/bookstore/api/routes"
	"github.com/vishalanarase/bookstore/internal/configs"
	"github.com/vishalanarase/bookstore/internal/datastore"
	"github.com/vishalanarase/bookstore/pkg/metrics"
)

type Application struct {
	router  *gin.Engine
	Metrics *metrics.Metrics
}

// NewApplication returns a new Application
func NewApplication() *Application {
	return &Application{
		// Create new engine instance
		router: gin.New(),
		// Set up metrics
		Metrics: metrics.NewMetrics(),
	}
}

// Start starts the application
func (app *Application) Start(envConfig configs.GlobalConfig) error {
	log.Info("Starting bookstore app")

	// Get database connection
	db, err := configs.DatabaseConnection(envConfig)
	if err != nil {
		log.Fatal(err, "Failed to connect to database")
	}

	// Register metrics
	app.Metrics.Register()

	// Set the mode
	//gin.SetMode(gin.ReleaseMode)

	// Use the logging
	app.router.Use(gin.Logger())

	// Recover from panic
	app.router.Use(gin.Recovery())

	// Rate limit api
	app.router.Use(middleware.RateLimitHandler)

	// Log the request
	app.router.Use(middleware.LogHandler)

	// Metrics
	app.router.Use(middleware.MetricsHandlerMiddleware(app.Metrics))

	// Register the routes
	routes.AddRoutes(app.router, app.Metrics, datastore.NewStore(db))

	// Start the api
	return app.router.Run(":8080")
}
