package main

import (
	"fmt"

	"github.com/adeebatgithub/biscut/connections"
)

func main() {
	manager, err := NewManager(connections.Sqlite{Database: "db.sqlite3"})
	if err != nil {
		fmt.Println(err)
	}

	UserTable := Table{
		TableName: "User",
		Fields: map[string]string{
			"Id":       manager.Dialect.PrimaryKey(),
			"username": manager.Dialect.VarCharField(25, true, true),
		},
	}

	err = manager.CreateTable(&UserTable)
	if err != nil {
		fmt.Println(err)
	}

	//insert into db
	//data := map[string]string{
	//	"username": "root",
	//}
	//err = manager.Insert("User", data)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//get all data => [map[]...]
	data, err := manager.FetchAll(UserTable.TableName, "", false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	//get data by where condition => [map[]...]
	//where := map[string]string{
	//	"Username": "root",
	//}
	//data, err := manager.FetchWhere(UserTable.TableName, where, "", false)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(data)

	//get all values from a col => []
	//data, err := manager.FetchCol(UserTable.TableName, "username", "", false)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(data)

	//get row by id => [map[]]
	//data, err := manager.FetchByID(UserTable.TableName, "1")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(data)

	//update row
	//where := map[string]string{
	//	"id": "1",
	//}
	//cols := map[string]string{
	//	"username": "superuser",
	//}
	//err = manager.Update(UserTable.TableName, cols, where)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//delete row
	//where := map[string]string{
	//	"id": "2",
	//}
	//err = manager.Delete(UserTable.TableName, where)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
