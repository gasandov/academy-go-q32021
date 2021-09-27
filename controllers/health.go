package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type health struct {}

type HealthController interface {
	GetHealthCheck(c echo.Context) error
}

// Verify if api is running by sending back a server message
func (h *health) GetHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Server up and running!")
}

func NewHealthController() HealthController {
	return &health{}
}
