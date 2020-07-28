package resolver

import (
	"errors"
	. "github.com/cacing69/graphql/config/database"
	"github.com/cacing69/graphql/lib"
	"github.com/cacing69/graphql/model"
	"github.com/cacing69/graphql/object"
	"github.com/gofiber/fiber"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

func AuthToken() *graphql.Field {
	return &graphql.Field{
		Type:        object.AuthTokenType,
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
			var data fiber.Map
			var user model.User

			auth, user := authenticate(p.Args["username"].(string), p.Args["password"].(string))

			if !auth {
				data = fiber.Map{
					"token": nil,
					"message": "invalid username/password",
				}
			} else {
				signedToken, err := lib.ClaimToken(user)

				if err != nil {
					return err, nil
				}

				data = fiber.Map{
					"token": signedToken,
					"message": "success generated token",
				}
			}


			return data, nil
		},
	}
}

func authenticate(username string, password string) (bool, model.User) {
	// query : get user from db
	var query model.User

	db := DB

	result := db.First(&query, "username = ?", username)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, model.User{}
	}

	if password == "password" {
		return true, model.User{
			Id:   1,
		}
	} else {
		return false, model.User{}
	}
}
