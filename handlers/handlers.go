package handlers

import (
	"aeperez24/goLambda/service"

	"github.com/gin-gonic/gin"
)

type PinHandlerImpl struct {
	PingService service.PingService
}

func (ph PinHandlerImpl) HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": ph.PingService.SendPing(),
	})
}
