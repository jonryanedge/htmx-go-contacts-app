package main

import (
	"fmt"
	"net/http"
	"strconv"

	"go.igmp.app/internal/contacts"

	"github.com/labstack/echo/v4"
)

func (app *app) getIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}

func (app *app) redirectIndex(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/contacts")
}

func (app *app) getContacts(c echo.Context) error {
	trigger := GetHeaders(c, "HX-Trigger")
	if trigger == "search" {
		search := c.QueryParam("q")
		if search != "" {
			contacts := contacts.SearchContacts(search)
			data := map[string]interface{}{
				"Contacts": contacts.Contacts,
				"Archiver": &app.archive,
				"Query":    search,
			}
			return c.Render(http.StatusOK, "partial.rows", data)
		} else {
			contacts := contacts.GetContacts()
			data := map[string]interface{}{
				"Contacts": contacts.Contacts,
				"Archiver": &app.archive,
				"Query":    search,
			}
			return c.Render(http.StatusOK, "partial.rows", data)
		}
	}
	contacts := contacts.GetContacts()
	data := map[string]interface{}{
		"Contacts": contacts.Contacts,
		"Archiver": &app.archive,
	}
	// data := fmt.Sprintf("contacts: %s\n", contacts)
	// return c.JSON(http.StatusOK, contacts)
	return c.Render(http.StatusOK, "layout", data)
}

func (app *app) getContactsNew(c echo.Context) error {
	data := map[string]interface{}{
		"Contact": map[string]string{
			"First": "",
			"Last":  "",
			"Email": "",
			"Phone": "",
		},
	}
	return c.Render(http.StatusOK, "new", data)
}

func (app *app) postContactsNew(c echo.Context) error {
	input := GetContactData(c)
	err := contacts.AddContact(input)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	return c.Redirect(http.StatusSeeOther, "/contacts")
}

func (app *app) getContact(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	contact, err := contacts.GetContact(id)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	data := map[string]interface{}{
		"Contact": contact,
	}
	return c.Render(http.StatusOK, "view", data)
}

func (app *app) getContactEmail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	contact, err := contacts.GetContact(id)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	return c.String(http.StatusOK, fmt.Sprintf("%s", contact.Email))
}

func (app *app) getContactEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	contact, err := contacts.GetContact(id)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	data := map[string]interface{}{
		"Contact": contact,
	}
	return c.Render(http.StatusOK, "edit", data)
}

func (app *app) postContactEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	input := GetContactData(c)
	err := contacts.UpdateContact(id, input)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/contacts/%d", id))
}

func (app *app) deleteContact(c echo.Context) error {
	trigger := GetHeaders(c, "HX-Trigger")
	id, _ := strconv.Atoi(c.Param("id"))
	err := contacts.DeleteContact(id)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	if trigger != "delete-btn" {
		return c.String(http.StatusOK, "")
	}
	return c.Redirect(http.StatusSeeOther, "/contacts")
}

func (app *app) deleteContacts(c echo.Context) error {
	selections := GetSelectedContacts(c, "selected_contact_ids")
	for _, sel := range selections {
		id, _ := strconv.Atoi(sel)
		contacts.DeleteContact(id)
	}
	return c.Redirect(http.StatusSeeOther, "/contacts")
}

// ARCHIVE handlers
func (app *app) getContactsArchive(c echo.Context) error {
	archiver := &app.archive
	data := map[string]interface{}{
		"Archiver": archiver,
	}
	return c.Render(http.StatusOK, "partial.archive", data)
}
func (app *app) postContactsArchive(c echo.Context) error {
	archiver := &app.archive
	archiver.Run()
	data := map[string]interface{}{
		"Archiver": archiver,
	}
	return c.Render(http.StatusOK, "partial.archive", data)
}
func (app *app) deleteContactsArchive(c echo.Context) error {
	archiver := &app.archive
	archiver.Reset()
	data := map[string]interface{}{
		"Archiver": archiver,
	}
	return c.Render(http.StatusOK, "partial.archive", data)
}
func (app *app) getContactsArchiveFile(c echo.Context) error {
	c.Response().Header().Set("Content-Disposition", "attachment; filename=contacts.json")
	c.Response().Header().Set("Content-Type", c.Request().Header.Get("Content-Type"))
	archiver := &app.archive
	return c.File(archiver.ArchiveFile())
}

// Feature handlers
func (app *app) getContactsCount(c echo.Context) error {
	count := contacts.GetContactCount()
	return c.String(http.StatusOK, "("+strconv.Itoa(count)+" total contacts)")
}
