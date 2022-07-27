package handlers

import (
	"aeperez24/goLambda/service"

	"github.com/gin-gonic/gin"
)

type PinHandlerImpl struct {
	PingService service.PingService
}

func (ph PinHandlerImpl) HandlePing(c *gin.Context) {
	writeResponse(c, 200, ph.PingService.SendPing())

}

func (ph PinHandlerImpl) HandleHello(c *gin.Context) {
	receivedRequest := make(map[string]string)
	c.ShouldBindJSON(&receivedRequest)
	writeResponse(c, 200, ph.PingService.SendHello(receivedRequest["name"]))
}

func writeResponse(context *gin.Context, status int, data interface{}) {
	context.JSON(status, gin.H{
		"message": data,
	})
}
