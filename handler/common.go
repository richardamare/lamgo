package handler

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"os"
)

func SetupHandlers(e *echo.Echo) {
	e.GET("/health", healthHandler)

	e.GET("/hello", helloHandler)
}

type messageResponse struct {
	Message string `json:"message"`
}

func getLogger(c echo.Context) *zap.Logger {
	requestId, ok := c.Get("requestId").(string)
	if !ok {
		requestId = "unknown"
	}

	isLambda := os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != ""

	var log *zap.Logger
	if isLambda {
		log = zap.Must(zap.NewProduction())
	} else {
		log = zap.Must(zap.NewDevelopment())
	}

	return log.With(zap.String("requestId", requestId))
}
