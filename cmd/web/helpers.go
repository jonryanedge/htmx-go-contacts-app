package main

import (
	"fmt"

	"go.scuttlebutt.app/internal/data"

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

func GetSelectedContacts(c echo.Context, key string) []string {
	req := c.Request()
	req.ParseForm()

	for k, v := range req.Form {
		if k == key {
			return v
		}
	}
	var list []string
	return list
}

func GetContactData(c echo.Context) data.Contact {
	req := c.Request()
	req.ParseForm()

	var first, last, email, phone string
	for k, v := range req.Form {
		if k == "first_name" {
			first = v[0]
		}
		if k == "last_name" {
			last = v[0]
		}
		if k == "email" {
			email = v[0]
		}
		if k == "phone" {
			phone = v[0]
		}
	}
	contact := data.Contact{
		First: first,
		Last:  last,
		Email: email,
		Phone: phone,
	}
	return contact
}
