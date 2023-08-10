package handler

import "github.com/labstack/echo/v4"

func SetupHandlers(e *echo.Echo) {
	e.GET("/health", healthHandler)
}

type messageResponse struct {
	Message string `json:"message"`
}
