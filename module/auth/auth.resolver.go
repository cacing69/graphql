package auth

import (
	"github.com/cacing69/graphql/app/lib"
	"github.com/cacing69/graphql/app/model"
	//"github.com/cacing69/graphql/module/user"
	"github.com/gofiber/fiber"
	"github.com/graphql-go/graphql"
	"log"
)

type Resolver struct {

}

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

//var AuthTokenType = graphql.NewObject(
//	graphql.ObjectConfig{
//		Name: "authToken",
//		Fields: graphql.Fields{
//			"token": &graphql.Field{
//				Type: graphql.String,
//			},
//			"user": &graphql.Field{
//				Type: user.Type,
//			},
//			"log": &graphql.Field{
//				Type: graphql.NewList(LogType),
//				Args: graphql.FieldConfigArgument{
//					"last": &graphql.ArgumentConfig{
//						Type: graphql.Int,
//					},
//				},
//			},
//		},
//	},
//)

func(p Resolver) Token() *graphql.Field {
	return &graphql.Field{
		Type:        TokenType,
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
			// log.Println(p.Args)

			// spew.Dump(p.Info.FieldName)

			var user model.User
			var logging interface{}

			selected, _ := lib.GetSelectedFields(p)

			if _, ok := selected["user"]; ok {
				user = model.User{
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
				"token": "this_is_token",
				"user":  user,
				"log":   logging,
			}

			return data, nil
		},
	}
}
