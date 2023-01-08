package model

type User struct {
	Model
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Age      *int    `json:"age"`
	Gender   *bool   `json:"gender" gorm:"default:false"`
	Phone    *string `json:"phone"`
	Address  *string `json:"address"`
}

type UserConnection struct {
	Total int64   `json:"total"`
	Nodes []*User `json:"nodes"`
}
