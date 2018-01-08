package server

import (
	"github.com/labstack/echo"
)

func Serve() {
	e := echo.New()
	setRoute(e)
	e.Logger.Fatal(e.Start(":8000"))
}
