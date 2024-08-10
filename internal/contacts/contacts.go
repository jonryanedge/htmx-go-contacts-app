package contacts

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
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

const contactsFile string = "internal/contacts/contacts.json"

func GetContactCount() int {
	file, err := os.Open(contactsFile)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteData, _ := io.ReadAll(file)
	var contacts Contacts
	json.Unmarshal(byteData, &contacts.Contacts)
	count := len(contacts.Contacts)
	return count
}

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
	for _, data := range contacts.Contacts {
		if data.ID == id {
			contact = data
			return contact, nil
		}
	}
	err = errors.New("user not found")

	return contact, err
}

func GetContacts() Contacts {
	file, err := os.Open(contactsFile)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteData, _ := io.ReadAll(file)
	var contacts Contacts
	json.Unmarshal(byteData, &contacts.Contacts)

	return contacts
}

func AddContact(n Contact) error {
	file, err := os.Open(contactsFile)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteData, _ := io.ReadAll(file)
	var contacts Contacts
	json.Unmarshal(byteData, &contacts.Contacts)

	lastId := contacts.Contacts[len(contacts.Contacts)-1].ID
	nextId := lastId + 1

	n.ID = nextId
	contacts.Contacts = append(contacts.Contacts, n)

	out, err := json.MarshalIndent(contacts.Contacts, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile(contactsFile, out, 0644)
	if err != nil {
		return err
	}
	return nil
}

func UpdateContact(id int, n Contact) error {
	file, err := os.Open(contactsFile)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	var contacts Contacts
	byteData, _ := io.ReadAll(file)
	json.Unmarshal(byteData, &contacts.Contacts)

	var keepers []Contact
	for _, data := range contacts.Contacts {
		if id == data.ID {
			data.First = n.First
			data.Last = n.Last
			data.Email = n.Email
			data.Phone = n.Phone

			keepers = append(keepers, data)
		}
		if id != data.ID {
			keepers = append(keepers, data)
		}
	}

	out, err := json.MarshalIndent(keepers, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(contactsFile, out, 0644)
	if err != nil {
		return err
	}
	return nil
}

func DeleteContact(id int) error {
	file, err := os.Open(contactsFile)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteData, _ := io.ReadAll(file)
	var contacts Contacts
	json.Unmarshal(byteData, &contacts.Contacts)

	var keepers []Contact
	for _, data := range contacts.Contacts {
		if id != data.ID {
			keepers = append(keepers, data)
		}
	}
	out, err := json.MarshalIndent(keepers, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile(contactsFile, out, 0644)
	if err != nil {
		return err
	}
	return nil
}

func SearchContacts(q string) Contacts {
	file, err := os.Open(contactsFile)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteData, _ := io.ReadAll(file)

	var contacts Contacts
	json.Unmarshal(byteData, &contacts.Contacts)

	var matches []Contact
	for _, data := range contacts.Contacts {
		match_first := strings.Contains(data.First, q)
		match_last := strings.Contains(data.Last, q)
		match_email := strings.Contains(data.Email, q)
		match_phone := strings.Contains(data.Phone, q)
		if match_first || match_last || match_email || match_phone {
			matches = append(matches, data)
		}
	}
	results := Contacts{
		Contacts: matches,
	}
	return results
}
