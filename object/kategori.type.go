package object

import "github.com/graphql-go/graphql"


var KategoriType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "kategori",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"nama": &graphql.Field{
				Type: graphql.String,
			},
			"kode_shopee": &graphql.Field{
				Type: graphql.String,
			},
			"kode_tokopedia": &graphql.Field{
				Type: graphql.String,
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
