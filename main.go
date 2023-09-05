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

	// if err := e.Start("localhost:2222"); err != nil {
	// 	e.Logger.Fatal(err)
	// }

	if err := e.StartTLS("localhost:2222", "./certs/_wildcard.example.dev+3.pem", "./certs/_wildcard.example.dev+3-key.pem"); err != nil {
		e.Logger.Fatal(err)
	}
}
