package entity

import value "ms-client/domain/valueobj"

// New function is used to create a client, the string at the end states that all input variable are string
// type, if you have other type for a field mthen try ti leave it at the end, eg (ID, name, image string, age int)
func NewClient(idType, idNumber, gender, firstName, secondName, firstLastName,
	secondLastName, birthdate, email, cellphone, address, city, state, country string) *Client {
	return &Client{
		IDNumber:       idNumber,
		IDType:         idType,
		FirstName:      firstName,
		SecondName:     secondName,
		FirstLastName:  firstLastName,
		SecondLastName: secondLastName,
		Birthdate:      birthdate,
		Gender:         gender,
		Contact:        value.Contact{Email: email, Cellphone: cellphone, Address: address, City: city, Country: country},
	}
}

//Client ... struct that contains the client attributes
// Note además las anotaciones encerradas e m ` esto ayudará a realizar automaticamente la conversion a json
// o viceversa
type Client struct {
	IDNumber       string        `json:"idNumber"`
	IDType         string        `json:"idType"`
	FirstName      string        `json:"firstName,omitempty"`
	SecondName     string        `json:"secondName,omitempty"`
	FirstLastName  string        `json:"firstLastName,omitempty"`
	SecondLastName string        `json:"secondLastName,omitempty"`
	Birthdate      string        `json:"birthdate,omitempty"`
	Gender         string        `json:"gender,omitempty"`
	Contact        value.Contact `json:"contact,omitempty"`
}
