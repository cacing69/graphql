package types

import "github.com/graphql-go/graphql"

var HistoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "history",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
			"desc": &graphql.Field{
				Type: graphql.String,
			},
			"viewer": &graphql.Field{
				Type: graphql.NewList(ViewerType),
				Args: graphql.FieldConfigArgument{
					"last": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
		},
	},
)
