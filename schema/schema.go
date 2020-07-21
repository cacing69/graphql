package schema

import (
	"github.com/cacing69/graphql/resolvers"
	"github.com/graphql-go/graphql"
)

var AggregateQuery = graphql.Fields{
	"ping":  resolvers.SysPing(),
	"user":  resolvers.User(),
	"users": resolvers.Users(),
}

var AggregateMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutation",
	Fields: graphql.Fields{
		"authToken":  resolvers.AuthToken(),
		"createUser": resolvers.CreateUser(),
		"updateUser": resolvers.CreateUser(),
		"deleteUser": resolvers.CreateUser(),
	},
})
