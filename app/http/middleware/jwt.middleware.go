package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
	"github.com/graphql-go/graphql/language/visitor"
)

var LOGIN_EXPIRATION_DURATION = time.Duration(24) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("secret")
var APPLICATION_NAME = "api-source"

var pulicQuery = map[string]bool{
	"ping":      true,
	"authToken": true,
	"register":  true,
}

func AuthGraphQL(ctx *fiber.Ctx) {

	query := ctx.Query("query")

	typeQuery, _ := getType(query)

	// spew.Dump(typeQuery)

	if pulicQuery[typeQuery] {
		// this is public query
		ctx.Next()
	} else {
		// this is protected query
		// spew.Dump(pulicQuery[typeQuery])
		//ctx.Next(errors.New("missing or malformed token"))
		//ctx.Next(errors.New("Invalid or expired token"))
		ctx.Next()
	}
}

func getType(query string) (string, error) {
	var fields []string

	field, err := parser.Parse(parser.ParseParams{Source: query})

	if err != nil {
		return "", err
	}

	v := &visitor.VisitorOptions{
		Enter: func(p visitor.VisitFuncParams) (string, interface{}) {
			if node, ok := p.Node.(*ast.Field); ok {
				fields = append(fields, node.Name.Value)
			}
			return visitor.ActionNoChange, nil
		},
	}

	visitor.Visit(field, v, nil)
	if len(fields) > 0 {
		return fields[0], nil
	} else {
		return "", nil
	}
}
