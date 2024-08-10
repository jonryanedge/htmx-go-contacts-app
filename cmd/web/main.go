package main

import (
	"errors"
	"html/template"
	"io"
	"strings"

	"go.igmp.app/internal/archiver"
	// "go.igmp.app/internal/contacts"

	"github.com/labstack/echo/v4"
)

type app struct {
	archive archiver.Archiver
	debug   bool
}

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	base := "base"
	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		name = parts[1]
		base = parts[0]
	}
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found ->" + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, base, data)
}

func main() {
	app := app{
		archive: *archiver.NewArchiver(),
		debug:   false,
	}

	e := echo.New()

	templates := make(map[string]*template.Template)
	templates["archive"] = template.Must(template.ParseFiles("ui/html/archive.html"))
	templates["rows"] = template.Must(template.ParseFiles("ui/html/rows.html"))
	templates["new"] = template.Must(template.ParseFiles("ui/html/new.html", "ui/html/base.html"))
	templates["view"] = template.Must(template.ParseFiles("ui/html/view.html", "ui/html/base.html"))
	templates["edit"] = template.Must(template.ParseFiles("ui/html/edit.html", "ui/html/base.html"))
	templates["layout"] = template.Must(template.ParseFiles("ui/html/archive.html", "ui/html/rows.html", "ui/html/controls.html", "ui/html/layout.html", "ui/html/base.html"))
	templates["home"] = template.Must(template.ParseFiles("ui/html/home.html", "ui/html/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	e.GET("/", app.redirectIndex)
	e.GET("/home", app.getIndex)
	e.GET("/contacts", app.getContacts)
	e.GET("/contacts/new", app.getContactsNew)
	e.POST("/contacts/new", app.postContactsNew)
	e.GET("/contacts/:id", app.getContact)
	e.GET("/contacts/:id/email", app.getContactEmail)
	e.GET("/contacts/:id/edit", app.getContactEdit)
	e.POST("/contacts/:id/edit", app.postContactEdit)
	e.DELETE("/contacts/:id", app.deleteContact)
	e.POST("/contacts/delete", app.deleteContacts)
	e.GET("/contacts/archive", app.getContactsArchive)
	e.POST("/contacts/archive", app.postContactsArchive)
	e.DELETE("/contacts/archive", app.deleteContactsArchive)
	e.GET("contacts/archive/file", app.getContactsArchiveFile)
	e.GET("contacts/count", app.getContactsCount)
	e.Debug = true
	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":3333"))
}
