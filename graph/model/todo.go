package model

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user"`
}

type Message struct {
	Text      string `json:"text" bson:"text"`
	SenderID  string `json:"sender" bson:"sender_id"`
	ReciverID string `json:"reciver" bson:"reciver_id"`
}
