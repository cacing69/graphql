package schema

import (
	"fmt"
	"log"

	"github.com/cacing69/graphql/entity"
	"github.com/gofiber/fiber"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

var LogType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "log",
		Fields: graphql.Fields{
			"key": &graphql.Field{
				Type: graphql.String,
			},
			"value": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var AuthTokenType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "authToken",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"user": &graphql.Field{
				Type: UserType,
			},
			"log": &graphql.Field{
				Type: graphql.NewList(LogType),
				Args: graphql.FieldConfigArgument{
					"last": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
		},
	},
)

func AuthToken() *graphql.Field {
	return &graphql.Field{
		Type:        AuthTokenType,
		Description: "get token auth",
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// .

			log.Println(p.Args)

			var user entity.User
			var logging interface{}

			// if ok {
			selected, _ := getSelectedFields(p)
			// log.Printf("", selected)
			if _, ok := selected["user"]; ok {
				user = entity.User{
					Name:  "Cacing69",
					Email: "cacingworm69@gmail.com",
				}
				log.Println("user_hit")
			}

			if _, ok := selected["log"]; ok {
				logging = []fiber.Map{
					fiber.Map{
						"key":   "key 1",
						"value": "value 1",
					},
					fiber.Map{
						"key":   "key 1",
						"value": "value 1",
					},
					fiber.Map{
						"key":   "key 1",
						"value": "value 1",
					},
				}
				log.Println("log_hit")
			}

			data := fiber.Map{
				"token": "AssadadsadasdlAKSdjsadasd.asdhasdbsadsdsadsad.sad.sad.d.sad.ad.ad.a.dsa",
				"user":  user,
				"log":   logging,
			}

			return data, nil
		},
	}
}

func getSelectedFields(params graphql.ResolveParams) (map[string]interface{}, error) {
	fieldASTs := params.Info.FieldASTs
	if len(fieldASTs) == 0 {
		return nil, fmt.Errorf("getSelectedFields: ResolveParams has no fields")
	}
	return selectedFieldsFromSelections(params, fieldASTs[0].SelectionSet.Selections)
}

func selectedFieldsFromSelections(params graphql.ResolveParams, selections []ast.Selection) (selected map[string]interface{}, err error) {
	selected = map[string]interface{}{}

	for _, s := range selections {
		switch s := s.(type) {
		case *ast.Field:
			if s.SelectionSet == nil {
				selected[s.Name.Value] = true
			} else {
				selected[s.Name.Value], err = selectedFieldsFromSelections(params, s.SelectionSet.Selections)
				if err != nil {
					return
				}
			}
		case *ast.FragmentSpread:
			n := s.Name.Value
			frag, ok := params.Info.Fragments[n]
			if !ok {
				err = fmt.Errorf("getSelectedFields: no fragment found with name %v", n)

				return
			}
			selected[s.Name.Value], err = selectedFieldsFromSelections(params, frag.GetSelectionSet().Selections)
			if err != nil {
				return
			}
		default:
			err = fmt.Errorf("getSelectedFields: found unexpected selection type %v", s)

			return
		}
	}

	return
}
