package model

import (
"database/sql"
)

type Merk struct {
	Id       uint64 `json:"id" gorm:"column:merk_id;primaryKey"`
	Nama string `json:"nama" gorm:"column:merk_nama"`
	Logo string `json:"logo" gorm:"column:merk_logo"`
	Alias1 string `json:"alias_1" gorm:"column:merk_alias_1"`

	//Tester   []*Tester `gorm:"foreignkey:user_id"`
	// default available column on models
	CreatedBy sql.NullInt64 `json:"created_by" gorm:"column:merk_created_by"`
	UpdatedBy sql.NullInt64 `json:"updated_by" gorm:"column:merk_updated_by"`
	DeletedBy sql.NullInt64 `json:"deleted_by" gorm:"column:merk_deleted_by"`
	CreatedAt sql.NullTime `json:"created_at" gorm:"column:merk_created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" gorm:"column:merk_updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"column:merk_deleted_at"`
}

func (Merk) TableName() string {
	return "m_merk"
}
