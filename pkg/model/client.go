package model

//Client ... struct that contains the client attributes
// Note además las anotaciones encerradas e m ` esto ayudará a realizar automaticamente la conversion a json
// o viceversa
type Client struct {
	ID             string  `json:"Id"`
	IDType         string  `json:"IdType"`
	FirstName      string  `json:"firstName,omitempty"`
	SecondName     string  `json:"secondName,omitempty"`
	FirstLastName  string  `json:"firstLastName,omitempty"`
	SecondLastName string  `json:"secondLastName,omitempty"`
	Age            int     `json:"age,omitempty"`
	Gender         string  `json:"gender,omitempty"`
	Contact        Contact `json:"contact,omitempty"`
}
