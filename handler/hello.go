package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func helloHandler(c echo.Context) error {
	log := getLogger(c)

	name := c.QueryParam("name")
	if name == "" {
		log.Info("invalid request",
			zap.String("message", "name query parameter is empty"),
			zap.String("name", name),
		)
		return c.JSON(400, messageResponse{Message: "name is empty"})
	}

	message := fmt.Sprintf("Hello %s!", name)

	return c.JSON(200, messageResponse{Message: message})
}
