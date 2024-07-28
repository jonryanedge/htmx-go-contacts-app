package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type app struct {
	debug bool
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/contacts")
	})
	e.Logger.Fatal(e.Start(":3333"))
}
