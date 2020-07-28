package model

import "time"

type User struct {
	Id       uint `gorm:"primarykey"`
	Nama     string `json:"nama" gorm:"column:nama"`
	Username string `json:"username" gorm:"column:username"`
	Foto     string `json:"foto" gorm:"column:foto"`
	NoHp     string `json:"no_hp" gorm:"column:no_hp"`
	Password string `json:"password" gorm:"column:password"`
	//Tester   []*Tester `gorm:"foreignkey:user_id"`
	// default available column on models
	CreatedBy uint64 `json:"created_by" gorm:"column:created_by"`
	UpdatedBy uint64 `json:"updated_by" gorm:"column:updated_by"`
	DeletedBy uint64 `json:"deleted_by" gorm:"column:deleted_by"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}