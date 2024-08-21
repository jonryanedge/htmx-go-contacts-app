package main

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"net/http"
	"strings"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c *http.Request) error {
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

func (app *app) SendString(w http.ResponseWriter, r *http.Request, status int, send string) {
	buf := new(bytes.Buffer)
	buf.WriteString(send)
	w.WriteHeader(status)
	buf.WriteTo(w)
}

func (app *app) Render(
	w http.ResponseWriter, r *http.Request, status int, name string, data interface{}) {
	base := "base"

	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		name = parts[1]
		base = parts[0]
	}

	page := name + ".html"
	ts, ok := app.templates[page]
	if !ok {
		err := errors.New("Template not found ->" + page)
		app.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buf, base, data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}
