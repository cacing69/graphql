package main

import (
	"log"

	"github.com/cacing69/graphql/entities"
	"github.com/cacing69/graphql/handlers"
	"github.com/cacing69/graphql/middlewares"
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
				c.JSON(entities.GrapqlError{
					Data: nil,
					Errors: []entities.GraphqlErrorList{
						{
							Message: err.Error(),
						},
					},
				})
			},
		},
	)

	route.Use("/v1", middlewares.Jwt)

	route.Post("/v1", func(ctx *fiber.Ctx) {
		result := handlers.GraphQLExecuteQuery(ctx, schema)
		ctx.JSON(result)
	})

	route.Listen(4000)
}
