package main

import (
	"net/http"

	"go.igmp.app/internal/contacts"
)

func GetHeaders(r *http.Request, key string) string {
	headers := r.Header
	trigger := headers.Get(key)

	// fmt.Printf("%s: %s\n", key, trigger)

	return trigger
}

func GetSelectedContacts(r *http.Request, key string) []string {
	req := r
	req.ParseForm()

	for k, v := range req.Form {
		if k == key {
			return v
		}
	}
	var list []string
	return list
}

func GetContactData(r *http.Request) contacts.Contact {
	req := r
	req.ParseForm()

	var first, last, email, phone string
	for k, v := range req.Form {
		if k == "first_name" {
			first = v[0]
		}
		if k == "last_name" {
			last = v[0]
		}
		if k == "email" {
			email = v[0]
		}
		if k == "phone" {
			phone = v[0]
		}
	}
	contact := contacts.Contact{
		First: first,
		Last:  last,
		Email: email,
		Phone: phone,
	}
	return contact
}
