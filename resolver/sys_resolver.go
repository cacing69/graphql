package resolver

import (
	"github.com/graphql-go/graphql"
)

func SysPing() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "ping healthcheck",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "200 ok", nil
		},
	}
}
