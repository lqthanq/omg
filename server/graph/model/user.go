package model

type User struct {
	Model
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Age      *int    `json:"age"`
	Gender   *bool   `json:"gender"`
	Phone    *string `json:"phone"`
	Address  *string `json:"address"`
}
