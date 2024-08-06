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
	e.Static("/static", "static")
	e.GET("/contacts", getContacts)
	e.GET("/contacts/new", getContactsNew)
	e.POST("/contacts/new", postContactsNew)
	e.GET("/contacts/:id", getContact)
	e.GET("/contacts/:id/email", getContactEmail)
	e.GET("/contacts/:id/edit", getContactEdit)
	e.POST("/contacts/:id/edit", postContactEdit)
	e.DELETE("/contacts/:id", deleteContact)
	e.DELETE("/contacts/", deleteContacts)
	e.GET("/contacts/archive", getContactsArchive)
	e.POST("/contacts/archive", postContactsArchive)
	e.DELETE("/contacts/archive", deleteContactsArchive)
	e.GET("contacts/archive/file", getContactsArchiveFile)
	e.GET("contacts/count", getContactsCount)
	e.Logger.Fatal(e.Start(":3333"))
}
