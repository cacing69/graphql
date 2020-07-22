package types

import "github.com/graphql-go/graphql"

var TesterType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "tester",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"key": &graphql.Field{
				Type: graphql.String,
			},
			"value": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
