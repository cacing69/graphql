package sys

import (
	"github.com/graphql-go/graphql"
)

type Resolver struct {
}

func (p Resolver) Ping() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "ping healthcheck",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "200 ok", nil
		},
	}
}