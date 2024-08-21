package main

import (
	"net/http"

	"go.igmp.app/ui"

	"github.com/justinas/alice"
)

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	dynamic := alice.New()

	mux.Handle("GET /", dynamic.ThenFunc(app.redirectIndex))
	mux.Handle("GET /home", dynamic.ThenFunc(app.getIndex))
	mux.Handle("GET /contacts", dynamic.ThenFunc(app.getContacts))
	mux.Handle("GET /contacts/new", dynamic.ThenFunc(app.getContactsNew))
	mux.Handle("POST /contacts/new", dynamic.ThenFunc(app.postContactsNew))
	mux.Handle("GET /contacts/:id", dynamic.ThenFunc(app.getContact))
	mux.Handle("GET /contacts/:id/email", dynamic.ThenFunc(app.getContactEmail))
	mux.Handle("GET /contacts/:id/edit", dynamic.ThenFunc(app.getContactEdit))
	mux.Handle("POST /contacts/:id/edit", dynamic.ThenFunc(app.postContactEdit))
	mux.Handle("DELETE /contacts/:id", dynamic.ThenFunc(app.deleteContact))
	mux.Handle("POST /contacts/delete", dynamic.ThenFunc(app.deleteContacts))
	mux.Handle("GET /contacts/archive", dynamic.ThenFunc(app.getContactsArchive))
	mux.Handle("POST /contacts/archive", dynamic.ThenFunc(app.postContactsArchive))
	mux.Handle("DELETE /contacts/archive", dynamic.ThenFunc(app.deleteContactsArchive))
	mux.Handle("GET /contacts/archive/file", dynamic.ThenFunc(app.getContactsArchiveFile))
	mux.Handle("GET /contacts/count", dynamic.ThenFunc(app.getContactsCount))

	standard := alice.New()

	return standard.Then(mux)
}
