package schema

import (
	"github.com/cacing69/graphql/resolver"
	"github.com/graphql-go/graphql"
)

var AggregateQuery = graphql.Fields{
	"ping":  resolver.SysPing(),
	"user":  resolver.UserRow(),
	"users": resolver.UserRows(),
}

var AggregateMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutation",
	Fields: graphql.Fields{
		"authToken":  resolver.AuthToken(),
		"userInsert": resolver.UserInsert(),
		"userUpdate": resolver.UserInsert(),
		"userDelete": resolver.UserInsert(),
	},
})


