package routes

import (
	h "login-pro/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	e.GET("/IP", func(c echo.Context) error {
		clientIP := echo.ExtractIPDirect()(c.Request())
		return c.String(http.StatusOK, "Client IP: "+clientIP)
	})
	e.POST("/login", h.LogIn)
	e.POST("/logout", h.LogOut)
	e.POST("/todo", h.CreateTodo)
}
