package main

import (
	"aeperez24/goLambda/config"
	"aeperez24/goLambda/handlers"
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	ginLambda = handlers.BuildGinLambdaServer()
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	if os.Getenv("LAMBDA_TASK_ROOT") != "" {
		lambda.Start(Handler)
	} else {
		startLocal()
	}
}

func startLocal() {
	config.LoadViperConfig("envs", "local")
	server := handlers.BuildGinServer(":8089")
	err := server.ListenAndServe()
	if err != nil {
		println(err)
		panic(err)
	}
}
