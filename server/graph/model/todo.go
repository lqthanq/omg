package model

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}
