package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("It's API")
	engine := gin.New()

	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start the gin server")
	}
}
