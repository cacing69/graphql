package database

import (
	"github.com/astaxie/beego/orm"
	"github.com/cacing69/graphql/app/model"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	datasource := "root:cacing.mysql@tcp(localhost:3306)/db_source?parseTime=true&charset=utf8"

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", datasource)

	orm.RegisterModel(new(model.User), new(model.Tester))

	orm.RunSyncdb("default", false, true)

	orm.Debug = true
}

func ORM() orm.Ormer {
	return orm.NewOrm()
}
