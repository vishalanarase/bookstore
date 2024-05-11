package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IPingController is a controller interface for PingController
type IPingController interface {
	Ping(ctx *gin.Context)
}

// PingController represents a PingController structure
type PingController struct {
}

// NewPingController initialise a new PingController and returns
func NewPingController() IPingController {
	return &PingController{}
}

// Ping returns a ping response
func (ctrl *PingController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
