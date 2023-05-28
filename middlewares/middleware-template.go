package middlewares

import (
	"github.com/labstack/echo/v4"
)

func MyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// your code here

		return next(c)
	}
}
