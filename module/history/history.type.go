package history

import (
	"github.com/cacing69/graphql/module/viewer"
	"github.com/graphql-go/graphql"
)

var Type = graphql.NewObject(
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
				Type: graphql.NewList(viewer.Type),
				Args: graphql.FieldConfigArgument{
					"last": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
		},
	},
)
