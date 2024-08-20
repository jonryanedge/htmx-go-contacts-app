package main

import (
	"html/template"
	"net/http"

	"go.igmp.app/ui"

	"github.com/labstack/echo/v4"
)

func (app *app) routes() http.Handler {
	e := echo.New()

	templates := make(map[string]*template.Template)
	templates["archive"] = template.Must(template.ParseFS(ui.Files, "html/archive.html"))
	templates["rows"] = template.Must(template.ParseFS(ui.Files, "html/rows.html"))
	templates["new"] = template.Must(template.ParseFS(ui.Files, "html/new.html", "html/base.html"))
	templates["view"] = template.Must(template.ParseFS(ui.Files, "html/view.html", "html/base.html"))
	templates["edit"] = template.Must(template.ParseFS(ui.Files, "html/edit.html", "html/base.html"))
	templates["layout"] = template.Must(template.ParseFS(ui.Files, "html/archive.html", "html/rows.html", "html/controls.html", "html/layout.html", "html/base.html"))
	templates["home"] = template.Must(template.ParseFS(ui.Files, "html/home.html", "html/base.html"))
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
	e.GET("/contacts/archive/file", app.getContactsArchiveFile)
	e.GET("/contacts/count", app.getContactsCount)
	e.Debug = app.debug

	if app.debug {
		// static file handler for running development
		e.Static("/static", "ui/static")
	} else {
		// static file handler for server binary
		e.GET("/static/*", echo.WrapHandler(http.FileServer(http.FS(ui.Files))))
	}

	return e
}
