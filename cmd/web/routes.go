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
	if len(c.QueryParams()) != 0 {
		search := c.QueryParam("q")
		if search != "" {
			contacts := data.SearchContacts(search)
			data := map[string]interface{}{
				"Contacts": contacts.Contacts,
				"Archive":  contacts,
				"Query":    search,
			}
			return c.Render(http.StatusOK, "partial.rows", data)
		} else {
			contacts := data.GetContacts()
			data := map[string]interface{}{
				"Contacts": contacts.Contacts,
				"Archive":  contacts,
				"Query":    search,
			}
			return c.Render(http.StatusOK, "partial.rows", data)
		}
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

func getContactsNew(c echo.Context) error {
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

func postContactsNew(c echo.Context) error { return nil }
func getContact(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	contact, err := data.GetContact(id)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	data := map[string]interface{}{
		"Contact": contact,
	}
	return c.Render(http.StatusOK, "view", data)
}

func getContactEmail(c echo.Context) error { return nil }
func getContactEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	contact, err := data.GetContact(id)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	data := map[string]interface{}{
		"Contact": contact,
	}
	return c.Render(http.StatusOK, "edit", data)
}

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
