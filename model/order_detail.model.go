package model

import (
	"database/sql"
)

type OrderDetail struct {
	Id       uint64 `json:"id" gorm:"column:order_detail_id;primaryKey"`
	OrderId uint64 `json:"order_id" gorm:"column:order_detail_order_id"`
	KategoriId uint64 `json:"kategori_id" gorm:"column:order_detail_kategori_id"`
	MerkId uint64 `json:"merk_id" gorm:"column:order_detail_merk_id"`
	Harga uint `json:"harga" gorm:"column:order_detail_harga"`
	HargaModal uint `json:"harga_modal" gorm:"column:order_detail_harga_modal"`
	Foto string `json:"foto" gorm:"column:order_detail_foto"`
	Qty uint `json:"qty" gorm:"column:order_detail_qty"`
	Potongan uint `json:"potongan" gorm:"column:order_detail_potongan"`
	Merk Merk `json:"merk" gorm:"foreignKey:order_detail_merk_id"`
	Kategori Kategori `json:"kategori" gorm:"foreignKey:order_detail_kategori_id"`
	//Tester   []*Tester `gorm:"foreignkey:user_id"`
	// default available column on models
	CreatedBy sql.NullInt64 `json:"created_by" gorm:"column:order_detail_created_by"`
	UpdatedBy sql.NullInt64 `json:"updated_by" gorm:"column:order_detail_updated_by"`
	DeletedBy sql.NullInt64 `json:"deleted_by" gorm:"column:order_detail_deleted_by"`
	CreatedAt sql.NullTime `json:"created_at" gorm:"column:order_detail_created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" gorm:"column:order_detail_updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"column:order_detail_deleted_at"`
}

func (OrderDetail) TableName() string {
	return "t_order_detail"
}
