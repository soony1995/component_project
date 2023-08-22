package routes

import (
	h "login-pro/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/login", h.LogIn)
	e.POST("/logout", h.LogOut)
	e.POST("/todo", h.CreateTodo)
}
