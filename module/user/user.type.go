package user

import (
	"github.com/cacing69/graphql/module/tester"
	"github.com/graphql-go/graphql"
)

var Type = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			// "history": &graphql.Field{
			// 	Type: graphql.NewList(HistoryType),
			// 	Args: graphql.FieldConfigArgument{
			// 		"id": &graphql.ArgumentConfig{
			// 			Type: graphql.Int,
			// 		},
			// 	},
			// },
			"tester": &graphql.Field{
				Type: graphql.NewList(tester.Type),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"last": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
		},
	},
)
