package resolver

import (
	"errors"
	. "github.com/cacing69/graphql/config/database"
	"github.com/cacing69/graphql/model"
	"github.com/cacing69/graphql/object"

	//"github.com/cacing69/graphql/app/db"
	//"github.com/volatiletech/sqlboiler/v4/queries/qm"
	//"strconv"

	//"github.com/cacing69/graphql/config/db"
	"github.com/davecgh/go-spew/spew"
	"github.com/graphql-go/graphql"
)

//type User struct {
//}

func UserFind() *graphql.Field {
	return &graphql.Field{
		Type:        object.UserType,
		Description: "get single user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			spew.Dump("user_resolver_find")

			id, _ := p.Args["id"].(int)

			//mod := []qm.QueryMod{}


			var res model.User
			db := DB

			//if lib.IsFieldExist("tester", p) {
			//	//mod = append(mod, qm.Load("Testers"))
			//	db = db.Preload("Tester")
			//}

			db.Find(&res, id)

			//mod = append(mod, qm.Where("id = ?", id))

			//res, _ := boiler.Users(mod...).
			//	One(p.Context, database.Con())
			//
			//mapperTester := []tester.Mapper{}

			//if lib.IsFieldExist("tester", p) {
				//for i, e := range res.R.Testers {
				//	mapperTester[i].ID = e.ID
				//	mapperTester[i].Key = e.Key.String
				//	mapperTester[i].Value = e.Value.String
				//}
			//}

			//mapper := Mapper{
			//	ID:     res.ID,
			//	Name:   res.Name,
			//	Email:  res.Email,
			//	Tester: mapperTester,
			//}

			//y := make([]b)

			//x := Mapper{
			//	ID: user.ID,
			//	Name: user.Name,
			//	Email: user.Email,
			//	Tester: y
			//}
			//for _, e := range user {
			//	spew.Dump(e.Name)
			//	for _, e2 := range e.R.Testers {
			//		spew.Dump(e2.Key.String, e2.Value.String)
			//	}
			//}

			//Model.u
			//Bind(p.Context, database.Con(), &x)

			//data.SetTesters()

			//fieldMap, _ := lib.GetSelectedFields(p)
			//
			//if _, ok := fieldMap["testers"]; ok {
			//	user.Testers().All(p.Context, database.Con())
			//	_ = testers
			//	spew.Dump(user)

			//user.R.Testers = tester
			//user.AddTesters(p.Context, database.Con(), false, tester[0])

			//spew.Dump(child_1)
			//}

			//if true {
			//	// BEEGO ORM
			//	//o := db.ORM()
			//	//
			//	//data := model.User{}
			//	//
			//	//fieldMap, _ := lib.GetSelectedFields(p)
			//	//argMap, _ := lib.GetMappedArgs(p)
			//	//
			//	//if _, ok := fieldMap["tester"]; ok {
			//	//	o.QueryTable("m_user").Filter("user_id", id).One(&data)
			//	//
			//	//	q1 := o.QueryTable("m_tester")
			//	//
			//	//	var child1 []*model.Tester
			//	//
			//	//	arg1 := argMap["tester"].(map[string]interface{})
			//	//
			//	//	if _, ok1 := arg1["last"]; ok1 {
			//	//		qLimit1, _ := strconv.Atoi(arg1["last"].(string))
			//	//
			//	//		q1 = q1.OrderBy("-tester_id")
			//	//		q1 = q1.Limit(qLimit1)
			//	//	}
			//	//
			//	//	q1.All(&child1)
			//	//
			//	//	data.Tester = child1
			//	//}
			//	//
			//	//return data, nil
			//
			//	// SQL BOILER
			//	//data , _ := db.MUsers(qm.Where("user_id=?", id)).One(p.Context, db.Con())
			//	//spew.Dump(data)
			//
			//	return nil, nil
			//}
			return res, nil
		},
	}
}

func UserGet() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(object.UserType),
		Description: "read user list",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			//mod := []qm.QueryMod{}

			db := DB

			//if lib.IsFieldExist("tester", p) {
			//	db = db.Preload("Tester")
			//	//mod = append(mod, qm.Load("Testers"))
			//}

			var res []model.User

			db.Find(&res)
			//mod = append(mod, qm.Where("id = ?", id))

			//var x UserAndTester
			//res, _ := boiler.Users(mod...).
			//	All(p.Context, database.Con())
			//
			//users := make([]Mapper, len(res))
			//for i, e := range res {
			//	mapperTester := []tester.Mapper{}
			//
			//	if lib.IsFieldExist("tester", p) {
			//		mapperTester = make([]tester.Mapper, len(e.R.Testers))
			//
			//		for i2, e2 := range e.R.Testers {
			//			mapperTester[i2].ID = e2.ID
			//			mapperTester[i2].Key = e2.Key.String
			//			mapperTester[i2].Value = e2.Value.String
			//		}

					//for i2, e2 = range e.R.Testers {
					//	mapperTester[i2].ID = e2.ID
					//	mapperTester[i2].Key = e2.Key.String
					//	mapperTester[i2].Value = e2.Value.String
					//}
			//	}
			//
			//	users[i].ID = e.ID
			//	users[i].Name = e.Name
			//	users[i].Email = e.Email
			//	users[i].Tester = mapperTester
			//}

			//var data []*model.User
			//q := db.ORM().QueryTable("m_user")
			//
			//q.All(&data)
			//
			//fieldMap, _ := lib.GetSelectedFields(p)
			//argMap, _ := lib.GetMappedArgs(p)
			//if _, ok := fieldMap["tester"]; ok {
			//	for _, e := range data {
			//		q1 := db.ORM().QueryTable("m_tester")
			//
			//		var child1 []*model.Tester
			//
			//		arg1 := argMap["tester"].(map[string]interface{})
			//
			//		if _, ok1 := arg1["last"]; ok1 {
			//			qLimit1, _ := strconv.Atoi(arg1["last"].(string))
			//
			//			q1 = q1.OrderBy("-tester_id")
			//			q1 = q1.Limit(qLimit1)
			//		}
			//
			//		q1.All(&child1)
			//		e.Tester = child1
			//	}
			//}

			return res, nil

		},
	}
}

func UserCreate() *graphql.Field {
	return &graphql.Field{
		Type:        object.UserType,
		Description: "insert a new user",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			//var data model.User

			//data.Name = p.Args["name"].(string)
			//data.Email = p.Args["email"].(string)
			//data.Password, _ = lib.HashPassword(p.Args["password"].(string))
			//
			//id, err := db.ORM().Insert(&data)
			//
			//if err != nil {
			//	return nil, err
			//} else {
			//	data.Id = id
			return nil, nil
			//}
		},
	}
}

func UserUpdate() *graphql.Field {
	return &graphql.Field{
		Type:        object.UserType,
		Description: "update user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return nil, errors.New("Lorem ipsum dolor sit amet")
		},
	}
}

func UserDelete() *graphql.Field {
	return &graphql.Field{
		Type:        object.UserType,
		Description: "delete a new user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return nil, errors.New("Lorem ipsum dolor sit amet")
		},
	}
}
