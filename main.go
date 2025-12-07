package main

import (
	"Godb/connections"
	"fmt"
)

func main() {
	manager, err := NewManager(connections.Sqlite{Database: "db.sqlite3"})
	if err != nil {
		fmt.Println(err)
	}

	UserTable := Tab

}
