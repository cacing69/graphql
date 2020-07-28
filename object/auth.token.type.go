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
			"message": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)