package resolver

import (
	. "github.com/cacing69/graphql/config/database"
	"github.com/cacing69/graphql/lib"
	"github.com/cacing69/graphql/model"
	"github.com/cacing69/graphql/object"
	"github.com/graphql-go/graphql"
)

func OrderFind() *graphql.Field {
	return &graphql.Field{
		Type:        object.OrderType,
		Description: "get single order",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, _ := p.Args["id"].(int)

			var res model.Order

			db := DB

			db = db.Preload("Pelanggan")

			db.First(&res, "order_id = ?", id)

			return res, nil
		},
	}
}

func OrderGet() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.OrderType),
		Description: "read order list",
		Args: graphql.FieldConfigArgument{
			"status": &graphql.ArgumentConfig{
				Type: graphql.NewList(object.OrderEnumStatus),
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var res []model.Order

			db := DB

			status := p.Args["status"]

			if status != nil {
				db = db.Where("order_status = ?", p.Args["status"].(int))
			}

			if lib.IsFieldExist("pelanggan", p) {
				db = db.Preload("Pelanggan")
				//db = db.Preload("Pelanggan", func(db *gorm.DB) *gorm.DB {
				//	// use custom parameter here
				//	return db
				//})
			}

			if lib.IsFieldExist("kurir", p) {
				db = db.Preload("Kurir")
			}

			if lib.IsFieldExist("kecamatan", p) {
				db = db.Preload("Kecamatan")
			}

			if lib.IsFieldExist("detail", p) {

				db = db.Preload("Detail")

				if lib.IsFieldExist("detail.merk", p) {
					db = db.Preload("Detail.Merk")
				}

				if lib.IsFieldExist("detail.merk", p) {
					db = db.Preload("Detail.Kategori")
				}
			}

			db = db.Order("order_id desc")

			page := 0

			if p.Args["page"] != nil {
				if p.Args["page"].(int) < 0 {
					page = p.Args["page"].(int) - 1
				}
			}

			limit := 20


			if p.Args["limit"] != nil {
				if p.Args["limit"].(int) < 0 {
					limit = p.Args["limit"].(int)
				}
			}

			db = db.Limit(limit).Offset(page * limit)

			db.Find(&res)

			return res, nil
		},
	}
}
