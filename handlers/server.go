package handlers

import (
	"net/http"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

func BuildGinServer(portNumber string) http.Server {
	return http.Server{Addr: portNumber, Handler: InitRouter()}
}

func BuildGinLambdaServer() *ginadapter.GinLambda {
	return ginadapter.New(InitRouter())
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping/", Ping)
	return router
}
