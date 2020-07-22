package handler

import (
	"github.com/gofiber/fiber"
	"github.com/graphql-go/graphql"
)

func GraphQLExecuteQuery(ctx *fiber.Ctx, schema graphql.Schema) *graphql.Result {
	token := ctx.Fasthttp.Request.Header.Peek("Authorization")

	ctx.Fasthttp.SetUserValue("authorization", token)

	// spew.Dump(ctx.Query("query"))

	query := ctx.Query("query")

	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
		Context:       ctx.Context(),
	}

	result := graphql.Do(params)

	return result
}
