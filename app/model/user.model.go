package model

type User struct {
	Id       uint `gorm:"primarykey"`
	Name     string
	Email    string
	Password string
	Tester   []*Tester `gorm:"foreignkey:user_id"`
}

func (User) TableName() string {
	return "users"
}