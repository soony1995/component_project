package routes

import (
	h "soon-pro/handlers"
	m "soon-pro/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", h.HelloWorld, m.MyMiddleware)
}
