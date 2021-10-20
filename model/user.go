package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"first_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

type Response struct {
	Data User
}
