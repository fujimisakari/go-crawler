package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func routing(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
