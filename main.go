package main

import (
	"login-pro/handlers"
	"login-pro/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	handlers.Init()

	e.Use(middleware.Logger())

	// e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{ // custom header
	// 	Skipper:      middleware.DefaultSkipper,
	// 	Generator:    func() string { return uuid.NewV4().String() }, // You can use any generator you prefer
	// 	TargetHeader: echo.HeaderXRequestID,
	// }))

	corsConfig := middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Replace with the allowed origins (e.g., "https://example.com")
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}

	e.Use(middleware.CORSWithConfig(corsConfig))

	routes.RegisterRoutes(e)

	if err := e.Start("localhost:2222"); err != nil {
		e.Logger.Fatal(err)
	}
}
