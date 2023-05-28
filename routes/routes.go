package routes

import (
	h "github.com/isaiorellana-dev/radio-api/handlers"
	m "github.com/isaiorellana-dev/radio-api/middlewares"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", h.HelloWorld, m.MyMiddleware)
}
