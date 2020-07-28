package object

import "github.com/graphql-go/graphql"


var KurirType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "kurir",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"nama": &graphql.Field{
				Type: graphql.String,
			},
			"logo": &graphql.Field{
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
