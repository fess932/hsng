// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Message struct {
	Text    string `json:"text"`
	Sender  *User  `json:"sender"`
	Reciver *User  `json:"reciver"`
}

type NewMessage struct {
	Text      string `json:"text"`
	ReciverID string `json:"reciverID"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
