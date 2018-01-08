package server

import (
	"github.com/labstack/echo"
)

func Serve() {
	e := echo.New()
	routing(e)
	e.Logger.Fatal(e.Start(":8000"))
}
