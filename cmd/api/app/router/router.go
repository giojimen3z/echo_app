package router

import (
	"github.com/giojimen3z/echo_app/cmd/api/app/controller"
	"github.com/labstack/echo/v4"
)

// SetupRoutes configura todas las rutas para la aplicación
func SetupRoutes(e *echo.Echo) {
	e.GET("/echo-app/ping", controller.HandlePing)
}
