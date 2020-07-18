package main

import (
	"fmt"
	"log"

	"github.com/cacing69/graphql/schema"
	"github.com/gofiber/fiber"
	"github.com/graphql-go/graphql"
)

var aggregateSchema = graphql.Fields{
	"user":  schema.UserRow(),
	"users": schema.UserRows(),
}

var aggregateMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutation",
	Fields: graphql.Fields{
		"user_insert": schema.UserInsert(),
		"user_update": schema.UserInsert(),
		"user_delete": schema.UserInsert(),
	},
})

func main() {
	query := graphql.ObjectConfig{Name: "query", Fields: aggregateSchema}

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    graphql.NewObject(query),
			Mutation: aggregateMutations,
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

	route.Listen(8000)
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Println(result.Errors)
	}

	return result
}
