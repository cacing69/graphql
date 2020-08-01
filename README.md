go run github.com/volatiletech/sqlboiler/v4 mysql 

### sub query
`
SELECT * FROM 
(
  SELECT foo.*
  FROM foo
  WHERE bar = "baz"
) AS t1
WHERE t1.id = 1;`

`db.Where("amount > ?", DB.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").QueryExpr()).Find(&orders)`

Which generate SQL
`SELECT * FROM "orders" WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders" WHERE (state = 'paid')));`

### raw
var result Result
`db.Raw("SELECT id, name, age FROM users WHERE name = ?", 3).Scan(&result)`

`db.Raw("SELECT id, name, age FROM users WHERE name = ?", 3).Scan(&result)`

`var age int
DB.Raw("select sum(age) from users where role = ?", "admin").Scan(&age)`

`db.Exec("UPDATE orders SET shipped_at=? WHERE id IN ?", time.Now(), []int64{1,2,3})`