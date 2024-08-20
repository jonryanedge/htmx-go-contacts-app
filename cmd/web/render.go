package main

import (
	"errors"
	"html/template"
	"io"
	"strings"

	"github.com/labstack/echo/v4"
)

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
