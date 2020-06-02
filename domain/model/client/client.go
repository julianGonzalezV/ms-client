package client

import "ms-client/domain/model/contact"

// New function is used to create a client, the string at the end states that all input variable are string
// type, if you have other type for a field mthen try ti leave it at the end, eg (ID, name, image string, age int)
func New(ID, IDType, firstName string) *Client {
	return &Client{
		ID:        ID,
		IDType:    IDType,
		FirstName: firstName,
	}
}

//Client ... struct that contains the client attributes
// Note además las anotaciones encerradas e m ` esto ayudará a realizar automaticamente la conversion a json
// o viceversa
type Client struct {
	ID             string          `json:"Id"`
	IDType         string          `json:"IdType"`
	FirstName      string          `json:"firstName,omitempty"`
	SecondName     string          `json:"secondName,omitempty"`
	FirstLastName  string          `json:"firstLastName,omitempty"`
	SecondLastName string          `json:"secondLastName,omitempty"`
	Age            int             `json:"age,omitempty"`
	Gender         string          `json:"gender,omitempty"`
	Contact        contact.Contact `json:"contact,omitempty"`
}
