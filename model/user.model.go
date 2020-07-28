package model

import (
	"database/sql"
)

type User struct {
	Id       uint `json:"id" gorm:"primaryKey"`
	Nama     string `json:"nama" gorm:"column:nama"`
	Username string `json:"username" gorm:"column:username"`
	Foto     string `json:"foto" gorm:"column:foto"`
	NoHp     string `json:"no_hp" gorm:"column:no_hp"`
	Password string `json:"password" gorm:"column:password"`
	//Tester   []*Tester `gorm:"foreignkey:user_id"`
	// default available column on models
	CreatedBy sql.NullInt64 `json:"created_by" gorm:"column:created_by"`
	UpdatedBy sql.NullInt64 `json:"updated_by" gorm:"column:updated_by"`
	DeletedBy sql.NullInt64 `json:"deleted_by" gorm:"column:deleted_by"`
	CreatedAt sql.NullTime `json:"created_at" gorm:"column:created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}