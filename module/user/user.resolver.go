package user

import (
	"errors"
	"strconv"

	"github.com/cacing69/graphql/app/lib"
	"github.com/cacing69/graphql/app/model"
	"github.com/cacing69/graphql/config/database"
	"github.com/davecgh/go-spew/spew"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
}

func (p Resolver) Find() *graphql.Field {
	return &graphql.Field{
		Type:        Type,
		Description: "get single user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// x, _ := lib.GetMappedArgs(p)
			// spew.Dump(p.Info.FieldName)
			// spew.Dump(p.Info.RootValue)

			// spew.Dump(p.Context.Value("authorization"))

			id, ok := p.Args["id"].(int)

			if ok {
				o := database.ORM()

				data := model.User{}

				o.QueryTable("m_user").Filter("user_id", id).One(&data)

				// spew.Dump(data)

				// err := o.Read(&data)

				// if err != nil {
				// 	return nil, err
				// }
				return data, nil
			}
			return nil, nil
		},
	}
}

func (p Resolver) Get() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(Type),
		Description: "read user list",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			o := database.ORM()
			var data []*model.User
			q := o.QueryTable("m_user")

			q.All(&data)

			fieldMap, _ := lib.GetSelectedFields(p)
			argMap, _ := lib.GetMappedArgs(p)
			if _, ok := fieldMap["tester"]; ok {
				for _, e := range data {
					// 	o.LoadRelated(e, "Tester", true, 3) --load all relation
					q1 := o.QueryTable("m_tester")

					var child1 []*model.Tester

					arg1 := argMap["tester"].(map[string]interface{})

					if _, ok1 := arg1["last"]; ok1 {
						spew.Dump("MASUK_NIH")
						spew.Dump(arg1["last"])
						qLimit1, _ := strconv.Atoi(arg1["last"].(string))

						q1 = q1.OrderBy("-tester_id")
						q1 = q1.Limit(qLimit1)
					}

					q1.All(&child1)
					e.Tester = child1
				}
			}

			return data, nil

		},
	}
}

func (p Resolver) Create() *graphql.Field {
	return &graphql.Field{
		Type:        Type,
		Description: "insert a new user",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return nil, errors.New("Lorem ipsum dolor sit amet")
		},
	}
}

func (p Resolver) Update() *graphql.Field {
	return &graphql.Field{
		Type:        Type,
		Description: "update user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return nil, errors.New("Lorem ipsum dolor sit amet")
		},
	}
}

func (p Resolver) Delete() *graphql.Field {
	return &graphql.Field{
		Type:        Type,
		Description: "delete a new user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return nil, errors.New("Lorem ipsum dolor sit amet")
		},
	}
}
