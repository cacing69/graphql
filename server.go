package main

import (
	"github.com/cacing69/graphql/app/http/handler"
	"log"

	"github.com/cacing69/graphql/app/entity"
	"github.com/cacing69/graphql/app/http/middleware"
	_ "github.com/cacing69/graphql/config/database"
	"github.com/cacing69/graphql/schema"
	"github.com/gofiber/fiber"
	"github.com/graphql-go/graphql"
)

func main() {

	query := graphql.ObjectConfig{Name: "query", Fields: schema.AggregateQuery}

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    graphql.NewObject(query),
			Mutation: schema.AggregateMutations,
		},
	)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	route := fiber.New(
		&fiber.Settings{
			Prefork:       true,
			CaseSensitive: true,
			StrictRouting: true,
			ErrorHandler: func(c *fiber.Ctx, err error) {
				c.JSON(entity.GrapqlError{
					Data: nil,
					Errors: []entity.GraphqlErrorList{
						{
							Message: err.Error(),
						},
					},
				})
			},
		},
	)

	route.Use("/graphql", middleware.Jwt)

	route.Get("/graphql", func(ctx *fiber.Ctx) {
		result := handler.GraphQLExecuteQuery(ctx, schema)
		ctx.JSON(result)
	})

	route.Listen(4000)
}
