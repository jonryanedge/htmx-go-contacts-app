package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type app struct {
	debug bool
}

func main() {
	msg := "go+htmx in-book contacts app"
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, msg)
	})
	e.Logger.Fatal(e.Start(":3333"))
}
