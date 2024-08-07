package main

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type app struct {
	debug bool
}

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found ->" + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base", data)
}

func main() {
	e := echo.New()

	templates := make(map[string]*template.Template)
	templates["layout"] = template.Must(template.ParseFiles("ui/html/controls.html", "ui/html/layout.html", "ui/html/base.html"))
	templates["home"] = template.Must(template.ParseFiles("ui/html/home.html", "ui/html/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	e.GET("/", getIndex)
	e.GET("/home", getIndex)
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

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":3333"))
}
