package order

import "github.com/graphql-go/graphql"

var Type = graphql.NewObject(
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
			"kurir_id": &graphql.Field{
				Type: graphql.Int,
			},
			"kecamatan_id": &graphql.Field{
				Type: graphql.Int,
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
			"no_Reff": &graphql.Field{
				Type: graphql.String,
			},
			"pelanggan_id": &graphql.Field{
				Type: graphql.Int,
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
