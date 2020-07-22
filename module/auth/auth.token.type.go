package auth

import (
	"github.com/cacing69/graphql/module/user"
	"github.com/graphql-go/graphql"
)

var TokenType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "authToken",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"user": &graphql.Field{
				Type: user.Type,
			},
			"log": &graphql.Field{
				Type: graphql.NewList(LogType),
				Args: graphql.FieldConfigArgument{
					"last": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
		},
	},
)