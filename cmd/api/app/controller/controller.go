package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandlePing maneja el POST request en /myapp/ping
func HandlePing(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
