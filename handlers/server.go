package handlers

import (
	"aeperez24/goLambda/service"
	"net/http"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

func BuildGinServer(portNumber string) http.Server {
	return http.Server{Addr: portNumber, Handler: buildRouter()}
}

func BuildGinLambdaServer() *ginadapter.GinLambda {
	return ginadapter.New(buildRouter())
}

func buildRouter() *gin.Engine {

	router := gin.Default()
	pingService := service.NewPingService()
	pinHander := PinHandlerImpl{PingService: pingService}
	router.GET("/ping/", pinHander.HandlePing)
	router.POST("/hello/", pinHander.HandleHello)
	return router
}
