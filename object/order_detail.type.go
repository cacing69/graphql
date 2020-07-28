package object

import "github.com/graphql-go/graphql"

//var OrderEnumStatus = graphql.NewEnum(graphql.EnumConfig{
//	Name: "enum_order_status",
//	Values: graphql.EnumValueConfigMap{
//		"SELESAI": &graphql.EnumValueConfig{Value: "selesai"},
//		"BATAL": &graphql.EnumValueConfig{Value: "dibatalkan"},
//		"KIRIM": &graphql.EnumValueConfig{Value: "dikirim"},
//		"TERIMA": &graphql.EnumValueConfig{Value: "diterima"},
//	},
//})

var OrderDetailType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "order_detail",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"merk": &graphql.Field{
				Type: MerkType,
			},
			"kategori": &graphql.Field{
				Type: KategoriType,
			},
			"harga": &graphql.Field{
				Type: graphql.Int,
			},
			"foto": &graphql.Field{
				Type: graphql.String,
			},
			"qty": &graphql.Field{
				Type: graphql.Int,
			},
			"harga_modal": &graphql.Field{
				Type: graphql.Int,
			},
			"potongan": &graphql.Field{
				Type: graphql.Int,
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
		},
	},
)
