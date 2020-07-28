package model

import "database/sql"

type Pelanggan struct {
	Id     uint64         `gorm:"primaryKey;column:pelanggan_id"`
	NoHp   string         `json:"no_hp" gorm:"column:pelanggan_no_hp"`
	Nama   string         `json:"nama" gorm:"column:pelanggan_nama"`
	Jk     string `json:"jk" gorm:"column:pelanggan_jk"`
	Alamat string `json:"alamat" gorm:"column:pelanggan_alamat"`
	Shopee string `json:"shopee" gorm:"column:pelanggan_shopee"`
	Tokopedia string `json:"tokopedia" gorm:"column:pelanggan_tokopedia"`
	Instagram string `json:"instagram" gorm:"column:pelanggan_instagram"`
	Bukalapak string `json:"bukalapak" gorm:"column:pelanggan_bukalapak"`
	KecamatanId uint64 `json:"kecamatan_id" gorm:"column:pelanggan_kecamatan_id"`
	Kecamatan Location `json:"kecamatan" gorm:"foreignKey:pelanggan_kecamatan_id"`
	//Order []Order `json:"order" gorm:"foreignKey:order_pelanggan_id"`
	// default available column on models
	CreatedBy sql.NullInt64 `json:"created_by" gorm:"column:pelanggan_created_by"`
	UpdatedBy sql.NullInt64 `json:"updated_by" gorm:"column:pelanggan_updated_by"`
	DeletedBy sql.NullInt64 `json:"deleted_by" gorm:"column:pelanggan_deleted_by"`
	CreatedAt sql.NullTime  `json:"created_at" gorm:"column:pelanggan_created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at" gorm:"column:pelanggan_updated_at"`
	DeletedAt sql.NullTime  `json:"deleted_at" gorm:"column:pelanggan_deleted_at"`
}

func (Pelanggan) TableName() string {
	return "m_pelanggan"
}
