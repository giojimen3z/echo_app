package main

import (
	"os"

	"github.com/giojimen3z/echo_app/cmd/api/app/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.SetupRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Puerto por defecto
	}

	e.Logger.Fatal(e.Start(":" + port))
}
