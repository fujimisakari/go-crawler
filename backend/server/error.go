package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func JSONHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := "Internal Server Error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = http.StatusText(code)
	}
	if !c.Response().Committed {
		c.JSON(code, map[string]interface{}{
			"statusCode": code,
			"message":    msg,
		})
	}
}
