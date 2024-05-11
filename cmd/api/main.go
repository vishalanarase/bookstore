package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/vishalanarase/bookstore/api/middleware"
	"github.com/vishalanarase/bookstore/api/routes"
	"github.com/vishalanarase/bookstore/internal/data"
	"github.com/vishalanarase/bookstore/internal/setup"
)

func main() {
	fmt.Println("It's API")
	engine := gin.New()

	config := setup.Config("../../")
	db, err := setup.DatabaseConnection(config)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to connect to database")
	}

	// Recover from panic
	engine.Use(gin.Recovery())

	// Rate limit api
	engine.Use(middleware.RateLimitHandler)

	engine.Use(middleware.Models(*data.NewModels(db)))

	routes.AddRoutes(engine)

	err = engine.Run(":8080")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start the gin server")
	}
}
