package viewer

import "github.com/graphql-go/graphql"

var Type = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "viewer",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
