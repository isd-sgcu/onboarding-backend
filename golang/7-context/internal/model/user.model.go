package model

type User struct {
	Base
	Email    string `json:"email"`
	Password string `json:"password"`
}
