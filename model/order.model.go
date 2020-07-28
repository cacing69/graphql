package model

import "time"

type Order struct {
	Id       uint64 `gorm:"primarykey"`
	No     string `json:"no" gorm:"column:order_no"`
	Via    string `json:"via" gorm:"column:order_via"`
	Status string `json:"status" gorm:"column:order_status"`
	StatusBayar string `json:"statusBayar" gorm:"column:order_status_bayar"`
	Tanggal time.Time `json:"tanggal" gorm:"column:order_tanggal"`
	KurirId uint64 `json:"kurir_id" gorm:"column:order_kurir_id"`
	KecamatanId uint64 `json:"kecamatan_id" gorm:"column:order_kecamatan_id"`
	TotalDp uint `json:"total_dp" gorm:"column:order_total_dp"`
	TotalOngkir uint `json:"total_ongkir" gorm:"column:order_total_ongkir"`
	SubTotal uint `json:"sub_total" gorm:"column:order_sub_total"`
	OrderTotal uint `json:"order_total" gorm:"column:order_total"`
	Alamat string `json:"alamat" gorm:"column:order_alamat"`
	NoResi string `json:"no_resi" gorm:"column:order_no_resi"`
	NoReff string `json:"no_reff" gorm:"column:order_no_reff"`
	PelangganId uint64 `json:"pelanggan_id" gorm:"column:order_pelanggan_id"`
	SelesaiBy uint64 `json:"selesai_by" gorm:"column:order_selesai_by"`
	DikirimBy uint64 `json:"created_by" gorm:"column:order_created_by"`
	BatalBy uint64 `json:"batal_by" gorm:"column:order_batal_by"`
	DiterimaBy uint64 `json:"diterima_by" gorm:"column:order_diterima_by"`
	BayarBy uint64 `json:"bayar_by" gorm:"column:order_bayar_by"`
	BatalAt time.Time `json:"batal_at" gorm:"column:order_batal_at"`
	SelesaiAt time.Time `json:"selesai_at" gorm:"column:order_selesai_at"`
	DikirimAt time.Time `json:"dikirim_at" gorm:"column:order_dikirim_at"`
	DiterimaAt time.Time `json:"diterima_at" gorm:"column:order_diterima_at"`
	BayarAt time.Time `json:"bayar_at" gorm:"column:order_bayar_at"`
	//Tester   []*Tester `gorm:"foreignkey:user_id"`
	// default available column on models
	CreatedBy uint64 `json:"created_by" gorm:"column:order_created_by"`
	UpdatedBy uint64 `json:"updated_by" gorm:"column:order_updated_by"`
	DeletedBy uint64 `json:"deleted_by" gorm:"column:order_deleted_by"`
	CreatedAt time.Time `json:"created_at" gorm:"column:order_created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:order_updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:order_deleted_at"`
}

func (Order) TableName() string {
	return "users"
}
