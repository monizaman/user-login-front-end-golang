package models

type User struct {
	ID        uint    `json:"id"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Fullname  string `json:"fullname"`
	Telephone string  `json:"telephone"`
}
