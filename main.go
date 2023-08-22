package main

import (
	"login-pro/handlers"
	"login-pro/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	handlers.Init()

	e.Use(middleware.Logger())

	routes.RegisterRoutes(e)

	if err := e.Start(":2222"); err != nil {
		e.Logger.Fatal(err)
	}
}
