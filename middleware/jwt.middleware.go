package middleware

import (
	"errors"
	"fmt"
	"github.com/cacing69/graphql/lib"
	"github.com/dgrijalva/jwt-go"
	"strings"

	//"time"
	//
	//"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
	"github.com/graphql-go/graphql/language/visitor"
)

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
		//ctx.Next()
		bearer := ctx.Fasthttp.Request.Header.Peek("Authorization")

		if len(bearer) == 0 {
			ctx.Next(errors.New("missing or malformed token"))
			return
		}

		if !strings.Contains(string(bearer), "Bearer") {
			ctx.Next(errors.New("invalid authorization"))
			return
		}

		authorization := strings.Replace(string(bearer), "Bearer ", "", -1)

		token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {

			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")

			} else if method != lib.JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("signing method invalid")
			}
			return lib.JWT_SIGNATURE_KEY, nil
		})

		if err != nil {
			ctx.Next(errors.New("invalid token"))
			return
		}

		claims, state := token.Claims.(jwt.MapClaims)

		if !state || !token.Valid {
			ctx.Next(err)
			return
		}

		ctx.Fasthttp.SetUserValue("auth", claims)
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
