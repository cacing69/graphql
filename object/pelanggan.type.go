package object

import "github.com/graphql-go/graphql"

var PelangganType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "pelanggan",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"nama": &graphql.Field{
				Type: graphql.String,
			},
			"no_hp": &graphql.Field{
				Type: graphql.String,
			},
			"jk": &graphql.Field{
				Type: graphql.String,
			},
			"alamat": &graphql.Field{
				Type: graphql.String,
			},
			"shopee": &graphql.Field{
				Type: graphql.String,
			},
			"instagram": &graphql.Field{
				Type: graphql.String,
			},
			"tokopedia": &graphql.Field{
				Type: graphql.String,
			},
			"bukalapak": &graphql.Field{
				Type: graphql.String,
			},
			"kecamatan": &graphql.Field{
				Type: LocationType,
			},
			// default type
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"deleted_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"created_by": &graphql.Field{
				Type: graphql.Int,
			},
			"updated_by": &graphql.Field{
				Type: graphql.Int,
			},
			"deleted_by": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
