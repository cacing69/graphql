package schema

import (
	"github.com/cacing69/graphql/resolver"
	"github.com/graphql-go/graphql"
)

var AggregateQuery = graphql.Fields{
	"ping":  resolver.SysPing(),
	// user query
	"user":  resolver.UserFind(),
	"users": resolver.UserGet(),
	// order query
	"order":  resolver.OrderFind(),
	"orders": resolver.OrderGet(),
}

var AggregateMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutation",
	Fields: graphql.Fields{
		"authToken":  resolver.AuthToken(),
		"createUser": resolver.UserCreate(),
		"updateUser": resolver.UserUpdate(),
		"deleteUser": resolver.UserDelete(),
	},
})
