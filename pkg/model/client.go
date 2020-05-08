package model

//Client ... struct that contains the client attributes
// Note además las anotaciones encerradas e m ` esto ayudará a realizar automaticamente la conversion a json
// o viceversa
type Client struct {
	ID   string `json:"ID"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
