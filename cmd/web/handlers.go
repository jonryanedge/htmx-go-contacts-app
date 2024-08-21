package main

import (
	"fmt"
	"net/http"
	"strconv"

	"go.igmp.app/internal/contacts"
)

func (app *app) getIndex(w http.ResponseWriter, r *http.Request) {
	app.Render(w, r, http.StatusOK, "home", nil)
}

func (app *app) redirectIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusMovedPermanently)
}

func (app *app) getContacts(w http.ResponseWriter, r *http.Request) {
	trigger := GetHeaders(r, "HX-Trigger")
	if trigger == "search" {
		search := r.URL.Query().Get("q")
		if search != "" {
			contacts := contacts.SearchContacts(search)
			data := map[string]interface{}{
				"Contacts": contacts.Contacts,
				"Archiver": &app.archive,
				"Query":    search,
			}
			app.Render(w, r, http.StatusOK, "partial.rows", data)
			return
		} else {
			contacts := contacts.GetContacts()
			data := map[string]interface{}{
				"Contacts": contacts.Contacts,
				"Archiver": &app.archive,
				"Query":    search,
			}
			app.Render(w, r, http.StatusOK, "partial.rows", data)
			return
		}
	}
	contacts := contacts.GetContacts()
	data := map[string]interface{}{
		"Contacts": contacts.Contacts,
		"Archiver": &app.archive,
	}
	// data := fmt.Sprintf("contacts: %s\n", contacts)
	// app.JSON(http.StatusOK, contacts)
	app.Render(w, r, http.StatusOK, "layout", data)
}

func (app *app) getContactsNew(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Contact": map[string]string{
			"First": "",
			"Last":  "",
			"Email": "",
			"Phone": "",
		},
	}
	app.Render(w, r, http.StatusOK, "new", data)
}

func (app *app) postContactsNew(w http.ResponseWriter, r *http.Request) {
	input := GetContactData(r)
	err := contacts.AddContact(input)
	if err != nil {
		app.SendString(w, r, http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func (app *app) getContact(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	contact, err := contacts.GetContact(id)
	if err != nil {
		app.SendString(w, r, http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	data := map[string]interface{}{
		"Contact": contact,
	}
	app.Render(w, r, http.StatusOK, "view", data)
}

func (app *app) getContactEmail(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	contact, err := contacts.GetContact(id)
	if err != nil {
		app.SendString(w, r, http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	app.SendString(w, r, http.StatusOK, fmt.Sprintf("%s", contact.Email))
}

func (app *app) getContactEdit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	contact, err := contacts.GetContact(id)
	if err != nil {
		app.SendString(w, r, http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	data := map[string]interface{}{
		"Contact": contact,
	}
	app.Render(w, r, http.StatusOK, "edit", data)
}

func (app *app) postContactEdit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	input := GetContactData(r)
	err := contacts.UpdateContact(id, input)
	if err != nil {
		app.SendString(w, r, http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	http.Redirect(w, r, fmt.Sprintf("/contacts/%d", id), http.StatusSeeOther)
}

func (app *app) deleteContact(w http.ResponseWriter, r *http.Request) {
	trigger := GetHeaders(r, "HX-Trigger")
	id, _ := strconv.Atoi(r.PathValue("id"))
	err := contacts.DeleteContact(id)
	if err != nil {
		app.SendString(w, r, http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	if trigger != "delete-btn" {
		app.SendString(w, r, http.StatusOK, "")
	}
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func (app *app) deleteContacts(w http.ResponseWriter, r *http.Request) {
	selections := GetSelectedContacts(r, "selected_contact_ids")
	for _, sel := range selections {
		id, _ := strconv.Atoi(sel)
		contacts.DeleteContact(id)
	}
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

// ARCHIVE handlers
func (app *app) getContactsArchive(w http.ResponseWriter, r *http.Request) {
	archiver := &app.archive
	data := map[string]interface{}{
		"Archiver": archiver,
	}
	app.Render(w, r, http.StatusOK, "partial.archive", data)
}
func (app *app) postContactsArchive(w http.ResponseWriter, r *http.Request) {
	archiver := &app.archive
	archiver.Run()
	data := map[string]interface{}{
		"Archiver": archiver,
	}
	app.Render(w, r, http.StatusOK, "partial.archive", data)
}
func (app *app) deleteContactsArchive(w http.ResponseWriter, r *http.Request) {
	archiver := &app.archive
	archiver.Reset()
	data := map[string]interface{}{
		"Archiver": archiver,
	}
	app.Render(w, r, http.StatusOK, "partial.archive", data)
}
func (app *app) getContactsArchiveFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=contacts.json")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	archiver := &app.archive
	http.ServeFile(w, r, archiver.ArchiveFile())
}

// Feature handlers
func (app *app) getContactsCount(w http.ResponseWriter, r *http.Request) {
	count := contacts.GetContactCount()
	app.SendString(w, r, http.StatusOK, "("+strconv.Itoa(count)+" total contacts)")
}
