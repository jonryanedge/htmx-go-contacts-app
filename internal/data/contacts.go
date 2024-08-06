package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Contact struct {
	ID     int               `json:"id"`
	First  string            `json:"first"`
	Last   string            `json:"last"`
	Phone  string            `json:"phone"`
	Email  string            `json:"email"`
	Errors map[string]string `json:"errors"`
}

type Contacts struct {
	Contacts []Contact
}

const contactsFile string = "internal/data/contacts.json"

func GetContact(id int) (Contact, error) {
	file, err := os.Open(contactsFile)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteData, _ := io.ReadAll(file)

	var contacts Contacts
	json.Unmarshal(byteData, &contacts.Contacts)

	var contact Contact
	for i := 0; i < len(contacts.Contacts); i++ {
		if contacts.Contacts[i].ID == id {
			contact = contacts.Contacts[i]
			return contact, nil
		} else {
			err := errors.New("user not found")
			return contact, err
		}
	}

	return contact, nil
}

func GetContacts() Contacts {
	file, err := os.Open(contactsFile)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteData, _ := io.ReadAll(file)

	// fmt.Println("ByteData:" + string(byteData))
	var contacts Contacts

	json.Unmarshal(byteData, &contacts.Contacts)

	// for i := 0; i < len(contacts.Contacts); i++ {
	// 	fmt.Println("Name: " + contacts.Contacts[i].First + " " + contacts.Contacts[i].Last)
	// }

	return contacts

}
