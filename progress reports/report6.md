Adding functions related to the items. Working on functions: delete,add and modify. They are on the main.go file. Everything appear in database mysql.
Create query like:
database.Exec("insert into productdb.Products (model, company, price) values (?, ?, ?)"
_, err := database.Exec("delete from productdb.Products where id = ?", id)
	_, err = database.Exec("update productdb.Products set model=?, company=?, price = ? where id = ?"
  
