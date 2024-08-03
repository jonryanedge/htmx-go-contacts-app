package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.scuttlebutt.app/internal/data"
)

func getContacts(c echo.Context) error {
	contacts := data.GetContacts()
	fmt.Printf("contacts: %s\n", contacts)
	return nil
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
