package model

type User struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
