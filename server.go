package main

import (
	"log"

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

	route := fiber.New()

	route.Get("/v1", func(ctx *fiber.Ctx) {
		result := executeQuery(ctx.Query("query"), schema)
		ctx.JSON(result)
	})

	route.Listen(4000)
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	return result
}
