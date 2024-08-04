package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Contact struct {
	ID    string `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	// Errors map[string]string `json:"errors"`
}

type Contacts struct {
	Contacts []Contact
}

const file string = "internal/data/contacts.json"

func GetContacts() Contacts {
	file, err := os.Open(file)
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
