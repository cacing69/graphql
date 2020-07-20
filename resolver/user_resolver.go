package resolver

import (
	"errors"

	"github.com/cacing69/graphql/entity"
	"github.com/davecgh/go-spew/spew"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

var mockHistory = []entity.History{
	{
		Id:     1,
		UserId: 1,
		Desc:   "lorem 1",
		Viewer: mockViewer,
	},
	{
		Id:     2,
		UserId: 2,
		Desc:   "lorem 2",
		Viewer: mockViewer,
	},
	{
		Id:     3,
		UserId: 1,
		Desc:   "lorem 3",
		Viewer: mockViewer,
	},
	{
		Id:     4,
		UserId: 1,
		Desc:   "lorem 4",
		Viewer: mockViewer,
	},
	{
		Id:     5,
		UserId: 3,
		Desc:   "lorem 5",
		Viewer: mockViewer,
	},
}

var mockViewer = []entity.Viewer{
	{
		Id:   1,
		Name: "lorem 1",
	},
	{
		Id:   2,
		Name: "lorem 2",
	},
	{
		Id:   3,
		Name: "lorem 3",
	},
	{
		Id:   4,
		Name: "lorem 4",
	},
	{
		Id:   5,
		Name: "lorem 5",
	},
}

var mockUser = []entity.User{
	{
		Id:       1,
		UserName: "cacing69",
		History:  mockHistory,
	},
	{
		Id:       2,
		UserName: "g9insane",
		History:  mockHistory,
	},
	{
		Id:       3,
		UserName: "alucrad",
		History:  mockHistory,
	},
}

var ViewrType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "viewer",
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

var ViewerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "viewer",
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

var HistoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "history",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
			"desc": &graphql.Field{
				Type: graphql.String,
			},
			"viewer": &graphql.Field{
				Type: graphql.NewList(ViewerType),
				Args: graphql.FieldConfigArgument{
					"last": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
		},
	},
)

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"userName": &graphql.Field{
				Type: graphql.String,
			},
			"history": &graphql.Field{
				Type: graphql.NewList(HistoryType),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
		},
	},
)

func User() *graphql.Field {
	return &graphql.Field{
		Type:        UserType,
		Description: "get single user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			x, _ := getMappedArgs(p)

			spew.Dump(x)

			id, ok := p.Args["id"].(int)
			if ok {
				return mockUser[id], nil
			}

			return nil, nil
		},
	}
}

func getMappedArgs(p graphql.ResolveParams) (map[string]interface{}, error) {
	fieldASTs := p.Info.FieldASTs
	return getArgsFromQuery(p, fieldASTs[0].SelectionSet.Selections)
}

func getArgsFromQuery(params graphql.ResolveParams, selections []ast.Selection) (args map[string]interface{}, err error) {
	args = map[string]interface{}{}

	for _, l1 := range selections {
		switch l1 := l1.(type) {
		case *ast.Field:

			arg := map[string]interface{}{}
			for _, c1 := range l1.Arguments {

				arg[c1.Name.Value] = c1.Value.GetValue()
				for _, l2 := range l1.SelectionSet.Selections {
					switch l2 := l2.(type) {
					case *ast.Field:
						arg2 := map[string]interface{}{}
						for _, d := range l2.Arguments {
							arg2[d.Name.Value] = d.Value.GetValue()
						}

						if _, ok := arg[l2.Name.Value]; !ok {
							arg[l2.Name.Value] = arg2
						}
					}
				}

			}

			args[l1.Name.Value] = arg
		}
	}

	return args, err
}

func Users() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(UserType),
		Description: "read user list",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return mockUser, nil
		},
	}
}

func CreateUser() *graphql.Field {
	return &graphql.Field{
		Type:        UserType,
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
		Type:        UserType,
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
		Type:        UserType,
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
