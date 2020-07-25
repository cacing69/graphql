package database

import (
	//"database/sql"
	"gorm.io/driver/mysql"

	//"github.com/go-sql-driver/mysql"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/volatiletech/sqlboiler/v4/boil"
	"gorm.io/gorm"
)

//var  con *sql.DB
var DB *gorm.DB
func Connect() {
	datasource := "root:cacing.mysql@tcp(localhost:3306)/go_graphql?parseTime=true&charset=utf8"

	//db, _ := sql.Open("mysql", datasource)
	//con = db

	db, _ := gorm.Open(mysql.Open(datasource), &gorm.Config{
		//Logger:
	})

	//gorm.DB{}.Debug(true)
	//db.Debug()
	//boil.DebugMode = true
	DB = db.Debug()
	//Con.Debug()
	//con.log
}

//func Con() *sql.DB {
//	return con
//}

//func ORM() orm.Ormer {
//	return orm.NewOrm()
//}
