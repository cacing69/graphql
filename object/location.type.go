package object

import "github.com/graphql-go/graphql"


var LocationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "location",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"parent_id": &graphql.Field{
				Type: graphql.Int,
			},
			"shortname": &graphql.Field{
				Type: graphql.String,
			},
			"postcode": &graphql.Field{
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
		},
	},
)
