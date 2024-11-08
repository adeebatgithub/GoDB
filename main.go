package main

import (
	"Godb/database"
	"database/sql"
	"fmt"
)

const TableName = "users"

var Users = map[string]string{
	"id":   database.PrimaryKey(),
	"name": database.TextField(false),
}

func main() {
	conf := database.PSQLConfig{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "root",
		Database: "postgres",
		SslMode:  "disable",
	}
	conn := database.Connect(conf)
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	err := database.CreateTable(conn, "users", Users)
	if err != nil {
		panic(err)
	}

	data, err := database.FetchByID(conn, TableName, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
