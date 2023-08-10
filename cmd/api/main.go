package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"os"
)

var echoLambda *echoadapter.EchoLambda
var e *echo.Echo

func init() {
	e = echo.New()

	// Hide the banner and port in the logs, when starting the local server,
	// This is not necessary, but it's nice to have
	// NOTE: The banner and port will not be shown when running on AWS Lambda
	e.HideBanner = true
	e.HidePort = true

	// Set up the handlers

	// Middleware
	e.Use(middleware.Recover())

	// Set up the lambda adapter
	echoLambda = echoadapter.New(e)
}

func main() {
	log := logrus.New()
	log.Level = logrus.DebugLevel

	// Check if we are running on AWS Lambda
	// AWS_LAMBDA_FUNCTION_NAME is set by AWS Lambda, and is empty otherwise
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		log.Debug("Running on AWS Lambda")

		// Make the handler available for AWS Lambda
		lambda.Start(lambdaHandler)
	} else {
		log.Info("Running on localhost:8080")
		// Make the handler run locally as a web server
		e.Logger.Fatal(e.Start(":8080"))
	}
}

func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}
