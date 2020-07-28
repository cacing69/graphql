package model


import (
"database/sql"
)

type Kategori struct {
	Id       uint64 `json:"id" gorm:"column:kategori_id;primaryKey"`
	Nama string `json:"nama" gorm:"column:kategori_nama"`
	KodeShopee string `json:"kode_shopee" gorm:"column:kategori_kode_shopee"`
	KodeTokopedia string `json:"kode_tokopedia" gorm:"column:kategori_kode_tokopedia"`

	//Tester   []*Tester `gorm:"foreignkey:user_id"`
	// default available column on models
	CreatedBy sql.NullInt64 `json:"created_by" gorm:"column:kategori_created_by"`
	UpdatedBy sql.NullInt64 `json:"updated_by" gorm:"column:kategori_updated_by"`
	DeletedBy sql.NullInt64 `json:"deleted_by" gorm:"column:kategori_deleted_by"`
	CreatedAt sql.NullTime `json:"created_at" gorm:"column:kategori_created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" gorm:"column:kategori_updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"column:kategori_deleted_at"`
}

func (Kategori) TableName() string {
	return "m_kategori"
}
