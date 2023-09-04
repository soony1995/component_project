package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// func CorsMiddleware() echo.MiddlewareFunc {
// 	corsConfig := middleware.CORSConfig{
// 		AllowOrigins: []string{"*"},
// 		AllowMethods: []string{http.MethodGet, http.MethodPost},
// 		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
// 	}
// 	return middleware.CORSWithConfig(corsConfig)
// }

func CorsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		corsConfig := middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{http.MethodGet, http.MethodPost},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			AllowCredentials: true, 

		}
		middleware.CORSWithConfig(corsConfig)(next)(c)
		return nil
	}
}

func EchoMiddleware(e *echo.Echo) {
	e.Use(
		middleware.Logger(),
		CorsMiddleware)
}
