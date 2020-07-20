package entity

type User struct {
	Id       int       `json:"id"`
	UserName string    `json:"userName"`
	Email    string    `json:"email"`
	History  []History `json:"history"`
}
