package schema

import (
	"github.com/cacing69/graphql/resolver"
	"github.com/graphql-go/graphql"
)

var AggregateQuery = graphql.Fields{
	"ping":  resolver.SysPing(),
	"user":  resolver.User(),
	"users": resolver.Users(),
}

var AggregateMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutation",
	Fields: graphql.Fields{
		"authToken":  resolver.AuthToken(),
		"createUser": resolver.CreateUser(),
		"updateUser": resolver.CreateUser(),
		"deleteUser": resolver.CreateUser(),
	},
})
