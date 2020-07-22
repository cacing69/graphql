package models

type User struct {
	Id     int64     `json:"id" orm:"column(user_id)"`
	Name   string    `json:"name" orm:"column(user_name)"`
	Email  string    `json:"email" orm:"column(user_email)"`
	Tester []*Tester `json:"tester" orm:"reverse(many)"`
	// History  []History `json:"history"`
}

func (t *User) TableName() string {
	return "m_user"
}
