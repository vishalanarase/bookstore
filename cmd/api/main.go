package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/vishalanarase/bookstore/api/controllers"
	"github.com/vishalanarase/bookstore/api/middleware"
	"github.com/vishalanarase/bookstore/internal/data"
)

func main() {
	fmt.Println("It's API")
	engine := gin.New()

	engine.Use(middleware.Models(*data.NewModels()))

	controllers.AddRoutes(engine)

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start the gin server")
	}
}
