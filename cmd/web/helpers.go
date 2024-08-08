package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func GetHeadersFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		headers := req.Header
		trigger := headers.Get("HX-Trigger")

		fmt.Println(trigger)

		return next(c)
	}
}

func GetHeaders(c echo.Context, key string) string {
	req := c.Request()
	headers := req.Header
	trigger := headers.Get(key)

	// fmt.Printf("%s: %s\n", key, trigger)

	return trigger
}
