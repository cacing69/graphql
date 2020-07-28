package object

import "github.com/graphql-go/graphql"

var OrderEnumStatus = graphql.NewEnum(graphql.EnumConfig{
	Name: "enum_order_status",
	Values: graphql.EnumValueConfigMap{
		"SELESAI": &graphql.EnumValueConfig{Value: "selesai"},
		"BATAL": &graphql.EnumValueConfig{Value: "dibatalkan"},
		"KIRIM": &graphql.EnumValueConfig{Value: "dikirim"},
		"TERIMA": &graphql.EnumValueConfig{Value: "diterima"},
	},
})

var OrderType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "order",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"no": &graphql.Field{
				Type: graphql.String,
			},
			"via": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"status_bayar": &graphql.Field{
				Type: graphql.String,
			},
			"tanggal": &graphql.Field{
				Type: graphql.DateTime,
			},
			"total_dp": &graphql.Field{
				Type: graphql.Int,
			},
			"total_ongkir": &graphql.Field{
				Type: graphql.Int,
			},
			"sub_total": &graphql.Field{
				Type: graphql.Int,
			},
			"order_total": &graphql.Field{
				Type: graphql.Int,
			},
			"alamat": &graphql.Field{
				Type: graphql.String,
			},
			"no_resi": &graphql.Field{
				Type: graphql.String,
			},
			"no_reff": &graphql.Field{
				Type: graphql.String,
			},
			"pelanggan": &graphql.Field{
				Type: PelangganType,
			},
			"kurir": &graphql.Field{
				Type: KurirType,
			},
			"kecamatan": &graphql.Field{
				Type: LocationType,
			},
			"detail": &graphql.Field{
				Type: graphql.NewList(OrderDetailType),
			},
			"selesai_by": &graphql.Field{
				Type: graphql.Int,
			},
			"dikirim_by": &graphql.Field{
				Type: graphql.Int,
			},
			"batal_by": &graphql.Field{
				Type: graphql.Int,
			},
			"diterima_by": &graphql.Field{
				Type: graphql.Int,
			},
			"bayar_by": &graphql.Field{
				Type: graphql.Int,
			},
			"selesai_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"dikirim_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"batal_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"diterima_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"bayar_at": &graphql.Field{
				Type: graphql.DateTime,
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
