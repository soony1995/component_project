package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	routes.RegisterRoutes(e)

	if err := e.Start(":2222"); err != nil {
		e.Logger.Fatal(err)
	}
}
