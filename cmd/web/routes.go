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
	trigger := GetHeaders(c, "HX-Trigger")
	if trigger == "search" {
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

func postContactsNew(c echo.Context) error {
	input := GetContactData(c)
	err := data.AddContact(input)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	return c.Redirect(http.StatusSeeOther, "/contacts")
}

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

func getContactEmail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	contact, err := data.GetContact(id)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	return c.String(http.StatusOK, fmt.Sprintf("%s", contact.Email))
}

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

func postContactEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	input := GetContactData(c)
	err := data.UpdateContact(id, input)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/contacts/%d", id))
}

func deleteContact(c echo.Context) error {
	trigger := GetHeaders(c, "HX-Trigger")
	id, _ := strconv.Atoi(c.Param("id"))
	err := data.DeleteContact(id)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	if trigger != "delete-btn" {
		return c.String(http.StatusOK, "")
	}
	return c.Redirect(http.StatusSeeOther, "/contacts")
}

func deleteContacts(c echo.Context) error {
	selections := GetSelectedContacts(c, "selected_contact_ids")
	for _, sel := range selections {
		id, _ := strconv.Atoi(sel)
		data.DeleteContact(id)
	}
	return c.Redirect(http.StatusSeeOther, "/contacts")
}

// ARCHIVE handlers
func getContactsArchive(c echo.Context) error     { return nil }
func postContactsArchive(c echo.Context) error    { return nil }
func deleteContactsArchive(c echo.Context) error  { return nil }
func getContactsArchiveFile(c echo.Context) error { return nil }

// Feature handlers
func getContactsCount(c echo.Context) error {
	count := data.GetContactCount()
	return c.String(http.StatusOK, "("+strconv.Itoa(count)+" total contacts)")
}
