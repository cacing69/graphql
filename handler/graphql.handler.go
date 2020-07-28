package handler

import (
	"github.com/cacing69/graphql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber"
	"github.com/graphql-go/graphql"
	"log"
)



func GraphQLExecuteQuery(ctx *fiber.Ctx) {
	//datasource := "root:cacing.mysql@tcp(localhost:3306)/db_source?parseTime=true&charset=utf8"
	//
	//conn, err := sql.Open("mysql", datasource)

	//boil.SetDB(conn)

	//testx , _ := db.MUsers().All(ctx.Context(), db.Con())
	//
	//spew.Dump(testx)


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

	token := ctx.Fasthttp.Request.Header.Peek("Authorization")

	ctx.Fasthttp.SetUserValue("authorization", token)

	// spew.Dump(ctx.Query("query"))

	queryParam := ctx.Query("query")

	params := graphql.Params{
		Schema:        schema,
		RequestString: queryParam,
		Context:       ctx.Context(),
	}

	result := graphql.Do(params)

	ctx.JSON(result)
}
