package resolvers

import (
	"errors"

	"github.com/cacing69/graphql/mocks"
	"github.com/cacing69/graphql/types"
	"github.com/graphql-go/graphql"
)

func User() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserType,
		Description: "get single user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// x, _ := libs.GetMappedArgs(p)
			// spew.Dump(p.Info.FieldName)
			// spew.Dump(p.Info.RootValue)

			// spew.Dump(p.Context.Value("authorization"))

			id, ok := p.Args["id"].(int)
			if ok {
				return mocks.UserMock[id], nil
			}

			return nil, nil
		},
	}
}

func Users() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.UserType),
		Description: "read user list",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return mocks.UserMock, nil
		},
	}
}

func CreateUser() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserType,
		Description: "insert a new user",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"userName": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return nil, errors.New("Lorem ipsum dolor sit amet")
		},
	}
}

func UserUpdate() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserType,
		Description: "update a new user",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"userName": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return nil, errors.New("Lorem ipsum dolor sit amet")
		},
	}
}

func UserDelete() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserType,
		Description: "delete a new user",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"userName": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return nil, errors.New("Lorem ipsum dolor sit amet")
		},
	}
}
