package resolvers

import (
	"errors"
	"strconv"

	"github.com/cacing69/graphql/configs"
	"github.com/cacing69/graphql/libs"
	"github.com/cacing69/graphql/models"
	"github.com/cacing69/graphql/types"
	"github.com/davecgh/go-spew/spew"
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
				o := configs.ORM()

				data := models.User{}

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

func Users() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.UserType),
		Description: "read user list",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			o := configs.ORM()
			var data []*models.User
			q := o.QueryTable("m_user")

			// qs.RelatedSel("m_tester")

			q.All(&data)

			// keys := make([]int, 0, len(mapEg))

			// keys := []int64{}

			// for _, e := range data {
			// 	keys = append(keys, e.Id)
			// }

			// spew.Dump(keys)

			fieldMap, _ := libs.GetSelectedFields(p)
			argMap, _ := libs.GetMappedArgs(p)
			// spew.Dump(argMap["tester"]["last"])
			if _, ok := fieldMap["tester"]; ok {
				// spew.Dump("TESTER_REQUESTED")
				for _, e := range data {
					// 	// o.LoadRelated(e, "Tester", true, 3) --load all relation
					q1 := o.QueryTable("m_tester")

					var child1 []*models.Tester

					arg1 := argMap["tester"].(map[string]interface{})

					if _, ok1 := arg1["last"]; ok1 {
						spew.Dump("MASUK_NIH")
						spew.Dump(arg1["last"])
						qLimit1, _ := strconv.Atoi(arg1["last"].(string))

						q1 = q1.OrderBy("-tester_id")
						q1 = q1.Limit(qLimit1)
					}

					// q1 = q1.Filter("tester_meta_id__in", keys)

					q1.All(&child1)

					// for _, c := range child1 {
					// 	for _, i := range data {
					// 		if c.User.Id == i.Id {
					// 			i.Tester = append(i.Tester, c)
					// 		}
					// 	}
					// }

					e.Tester = child1
				}
			}

			return data, nil

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

func UserDelete() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserType,
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
