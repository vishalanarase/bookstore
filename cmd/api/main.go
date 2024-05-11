package main

import (
	"path"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vishalanarase/bookstore/api/middleware"
	"github.com/vishalanarase/bookstore/api/routes"
	"github.com/vishalanarase/bookstore/internal/config"
	"github.com/vishalanarase/bookstore/internal/datastore"
)

func init() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
			return "", fileName
		},
	})
}

func main() {
	log.Info("It's API")

	engine := gin.New()

	envConfig := config.Config("../../")
	db, err := config.DatabaseConnection(envConfig)
	if err != nil {
		log.Fatal(err, "Failed to connect to database")
	}

	// Recover from panic
	engine.Use(gin.Recovery())

	// Rate limit api
	engine.Use(middleware.RateLimitHandler)

	engine.Use(middleware.Models(*datastore.NewModels(db)))

	routes.AddRoutes(engine)

	err = engine.Run(":8080")
	if err != nil {
		log.Fatal(err, "Failed to start the gin server")
	}
}
