package object

import (
	//"github.com/cacing69/graphql/module/tester"
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"nama": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"foto": &graphql.Field{
				Type: graphql.String,
			},
			"avatar": &graphql.Field{
				Type: graphql.String,
			},
			"no_hp": &graphql.Field{
				Type: graphql.String,
			},
			// default type
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"deleted_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"created_by": &graphql.Field{
				Type: graphql.Int,
			},
			"updated_by": &graphql.Field{
				Type: graphql.Int,
			},
			"deleted_by": &graphql.Field{
				Type: graphql.Int,
			},
			// "history": &graphql.Field{
			// 	Type: graphql.NewList(HistoryType),
			// 	Args: graphql.FieldConfigArgument{
			// 		"id": &graphql.ArgumentConfig{
			// 			Type: graphql.Int,
			// 		},
			// 	},
			// },
			//"tester": &graphql.Field{
			//	Type: graphql.NewList(tester.Type),
			//	Args: graphql.FieldConfigArgument{
			//		"id": &graphql.ArgumentConfig{
			//			Type: graphql.Int,
			//		},
			//		"last": &graphql.ArgumentConfig{
			//			Type: graphql.Int,
			//		},
			//	},
			//},
		},
	},
)
