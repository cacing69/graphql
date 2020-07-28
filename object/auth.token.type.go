package object

import (
	"github.com/graphql-go/graphql"
)

var AuthTokenType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "authToken",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"user": &graphql.Field{
				Type: UserType,
			},
			//"log": &graphql.Field{
			//	Type: graphql.NewList(auth.LogType),
			//	Args: graphql.FieldConfigArgument{
			//		"last": &graphql.ArgumentConfig{
			//			Type: graphql.Int,
			//		},
			//	},
			//},
		},
	},
)