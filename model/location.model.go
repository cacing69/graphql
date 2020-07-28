package model


import (
	"database/sql"
)

type Location struct {
	Id       uint64 `json:"id" gorm:"column:location_id;primaryKey"`
	Name string `json:"name" gorm:"column:location_name"`
	Shortname string `json:"shortname" gorm:"column:location_shortname"`
	Postcode string `json:"postcode" gorm:"column:location_postcode"`
	ParentId uint64 `json:"parent_id" gorm:"column:location_parent_id"`
	Type uint64 `json:"type" gorm:"column:location_type"`

	//Tester   []*Tester `gorm:"foreignkey:user_id"`
	CreatedAt sql.NullTime `json:"created_at" gorm:"column:location_created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" gorm:"column:location_updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"column:location_deleted_at"`
}

func (Location) TableName() string {
	return "c_location"
}
