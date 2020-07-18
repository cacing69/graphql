package schema

import (
	"errors"

	"github.com/cacing69/graphql/entity"
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func UserRow() *graphql.Field {
	return &graphql.Field{
		Type:        UserType,
		Description: "get single user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			_, ok := p.Args["id"].(int)
			if ok {
				data := entity.User{
					Id:   1,
					Name: "cacing69",
				}
				return data, nil
			}
			return nil, nil
		},
	}
}

func UserRows() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(UserType),
		Description: "read user list",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			data := []entity.User{
				{
					Id:   1,
					Name: "cacing69",
				},
				{
					Id:   2,
					Name: "g9insane",
				},
				{
					Id:   3,
					Name: "alucrad",
				},
			}
			return data, nil
		},
	}
}

func UserInsert() *graphql.Field {
	return &graphql.Field{
		Type:        UserType,
		Description: "insert a new user",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
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
		Type:        UserType,
		Description: "update a new user",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
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
		Type:        UserType,
		Description: "delete a new user",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return nil, errors.New("Lorem ipsum dolor sit amet")
		},
	}
}
