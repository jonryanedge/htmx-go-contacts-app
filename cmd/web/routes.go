package main

import (
	//  "encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go.scuttlebutt.app/internal/data"

	"github.com/labstack/echo/v4"
)

func getIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}

func redirectIndex(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/contacts")
}

func getContacts(c echo.Context) error {
	search := c.QueryParam("q")
	if search != "" {
		contacts := data.GetContacts()
		data := map[string]interface{}{
			"Contacts": contacts.Contacts,
			"Archive":  contacts,
			"Query":    search,
		}
		return c.Render(http.StatusOK, "rows", data)
	}
	contacts := data.GetContacts()
	data := map[string]interface{}{
		"Contacts": contacts.Contacts,
		"Archive":  contacts,
	}
	// data := fmt.Sprintf("contacts: %s\n", contacts)
	// return c.JSON(http.StatusOK, contacts)
	return c.Render(http.StatusOK, "layout", data)
}

func getContactsNew(c echo.Context) error  { return nil }
func postContactsNew(c echo.Context) error { return nil }
func getContact(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	contact, err := data.GetContact(id)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	return c.JSON(http.StatusOK, contact)
}

func getContactEmail(c echo.Context) error        { return nil }
func getContactEdit(c echo.Context) error         { return nil }
func postContactEdit(c echo.Context) error        { return nil }
func deleteContact(c echo.Context) error          { return nil }
func deleteContacts(c echo.Context) error         { return nil }
func getContactsArchive(c echo.Context) error     { return nil }
func postContactsArchive(c echo.Context) error    { return nil }
func deleteContactsArchive(c echo.Context) error  { return nil }
func getContactsArchiveFile(c echo.Context) error { return nil }

func getContactsCount(c echo.Context) error {
	count := data.GetContactCount()
	return c.String(http.StatusOK, "("+strconv.Itoa(count)+" total contacts)")
}
