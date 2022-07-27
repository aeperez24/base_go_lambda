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

func (ph PinHandlerImpl) HandleHello(c *gin.Context) {
	receivedRequest := make(map[string]string)
	c.ShouldBindJSON(&receivedRequest)
	c.JSON(200, gin.H{
		"message": ph.PingService.SendHello(receivedRequest["name"]),
	})
}
