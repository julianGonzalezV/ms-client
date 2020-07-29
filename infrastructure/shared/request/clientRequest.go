package request

type Contact struct {
	Email, Cellphone, Address, City, State, Country string
}

type ClientRequest struct {
	IDType         string  `json:"idType"`
	IdNumber       string  `json:"idNumber"`
	Gender         string  `json:"gender"`
	FirstName      string  `json:"firstName"`
	SecondName     string  `json:"secondName"`
	FirstLastName  string  `json:"firstLastName"`
	SecondLastName string  `json:"secondLastName"`
	Birthdate      string  `json:"birthdate"`
	Contact        Contact `json:"contact"`
}
