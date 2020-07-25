package main

import (
	"github.com/cacing69/graphql/app/entity"
	"github.com/cacing69/graphql/app/http/handler"
	"github.com/cacing69/graphql/app/http/middleware"
	"github.com/cacing69/graphql/config/database"
	"github.com/cacing69/graphql/config/env"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	fiberMiddleware "github.com/gofiber/fiber/middleware"
	"strconv"
)

func main() {
	database.Connect()

	app := fiber.New(
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

	//app.Use(cors)
	app.Use(cors.New())
	app.Use(fiberMiddleware.Recover())

	graph := app.Group("/graphql")
	graph.Get("", middleware.AuthGraphQL, handler.GraphQLExecuteQuery)

	p := env.Get("SERVER_PORT")

	port, _ := strconv.Atoi(p)

	app.Listen(port)
}
