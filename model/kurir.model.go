package model


import (
	"database/sql"
)

type Kurir struct {
	Id       uint64 `json:"id" gorm:"column:kurir_id;primaryKey"`
	Nama string `json:"nama" gorm:"column:kurir_nama"`

	//Tester   []*Tester `gorm:"foreignkey:user_id"`
	// default available column on models
	CreatedBy sql.NullInt64 `json:"created_by" gorm:"column:kurir_created_by"`
	UpdatedBy sql.NullInt64 `json:"updated_by" gorm:"column:kurir_updated_by"`
	DeletedBy sql.NullInt64 `json:"deleted_by" gorm:"column:kurir_deleted_by"`
	CreatedAt sql.NullTime `json:"created_at" gorm:"column:kurir_created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" gorm:"column:kurir_updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"column:kurir_deleted_at"`
}

func (Kurir) TableName() string {
	return "m_kurir"
}
