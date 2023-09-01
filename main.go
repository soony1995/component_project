package main

import (
	"login-pro/handlers"
	"login-pro/middlewares"
	"login-pro/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	handlers.Init()

	middlewares.EchoMiddleware(e)

	routes.RegisterRoutes(e)

	if err := e.Start("localhost:2222"); err != nil {
		e.Logger.Fatal(err)
	}
}
