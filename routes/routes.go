package routes

import (
	h "login-pro/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		return c.String(http.StatusOK, "Request ID: "+reqID)
	})
	e.POST("/login", h.LogIn)
	e.POST("/logout", h.LogOut)
	e.POST("/todo", h.CreateTodo)
}
