package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/richardamare/lamgo/handler"
	"go.uber.org/zap"
	"os"
)

var echoLambda *echoadapter.EchoLambda
var e *echo.Echo
var isLambda bool

func init() {
	e = echo.New()

	// Hide the banner and port in the logs, when starting the local server,
	// This is not necessary, but it's nice to have
	// NOTE: The banner and port will not be shown when running on AWS Lambda
	e.HideBanner = true
	e.HidePort = true

	// AWS_LAMBDA_FUNCTION_NAME is set by AWS Lambda, and is empty otherwise
	isLambda = os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != ""

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		requestId := uuid.NewString()

		var log *zap.Logger
		if isLambda {
			log = zap.Must(zap.NewProduction())
		} else {
			log = zap.Must(zap.NewDevelopment())
		}

		return func(c echo.Context) error {
			c.Set("requestId", requestId)
			log.Info("received request",
				zap.String("requestId", requestId),
				zap.String("path", c.Path()),
				zap.String("method", c.Request().Method),
			)
			return next(c)
		}
	})

	// Set up the handlers
	handler.SetupHandlers(e)

	// Middleware
	e.Use(middleware.Recover())

	// Set up the lambda adapter
	echoLambda = echoadapter.New(e)
}

func main() {

	// Check if we are running on AWS Lambda
	if isLambda {
		log := zap.Must(zap.NewProduction())
		log.Debug("Running on AWS Lambda")

		// Make the handler available for AWS Lambda
		lambda.Start(lambdaHandler)
	} else {
		log := zap.Must(zap.NewDevelopment())
		log.Info("Running on localhost:8080")
		// Make the handler run locally as a web server
		log.Fatal(e.Start(":8080").Error())
	}
}

func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}
