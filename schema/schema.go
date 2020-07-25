package schema

import (
	"github.com/cacing69/graphql/module/auth"
	"github.com/cacing69/graphql/module/sys"
	"github.com/cacing69/graphql/module/user"
	"github.com/graphql-go/graphql"
)

var AggregateQuery = graphql.Fields{
	"ping":  sys.Resolver{}.Ping(),
	"user":  user.Resolver{}.Find(),
	"users": user.Resolver{}.Get(),
}

var AggregateMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutation",
	Fields: graphql.Fields{
		"auth_token":  auth.Resolver{}.Token(),
		"create_user": user.Resolver{}.Create(),
		"update_user": user.Resolver{}.Update(),
		"delete_user": user.Resolver{}.Delete(),
	},
})
