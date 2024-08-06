package main

import (
	//  "encoding/json"
	// "fmt"
	"net/http"

	"go.scuttlebutt.app/internal/data"

	"github.com/labstack/echo/v4"
)

func getContacts(c echo.Context) error {
	contacts := data.GetContacts()
	// data := fmt.Sprintf("contacts: %s\n", contacts)
	return c.JSON(http.StatusOK, contacts)
}
func getContactsNew(c echo.Context) error         { return nil }
func postContactsNew(c echo.Context) error        { return nil }
func getContact(c echo.Context) error             { return nil }
func getContactEmail(c echo.Context) error        { return nil }
func getContactEdit(c echo.Context) error         { return nil }
func postContactEdit(c echo.Context) error        { return nil }
func deleteContact(c echo.Context) error          { return nil }
func deleteContacts(c echo.Context) error         { return nil }
func getContactsArchive(c echo.Context) error     { return nil }
func postContactsArchive(c echo.Context) error    { return nil }
func deleteContactsArchive(c echo.Context) error  { return nil }
func getContactsArchiveFile(c echo.Context) error { return nil }
func getContactsCount(c echo.Context) error       { return nil }
