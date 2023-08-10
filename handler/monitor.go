package handler

import "github.com/labstack/echo/v4"

func healthHandler(c echo.Context) error {
	return c.JSON(200, messageResponse{Message: "OK"})
}
