package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IPingController interface {
	Ping(ctx *gin.Context)
}

type PingController struct {
}

func NewPingController() IPingController {
	return &PingController{}
}

func (ctrl *PingController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
