package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type PSQLConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SslMode  string
}

func Connect(config PSQLConfig) *sql.DB {
	conf := fmt.Sprintf(
		"dbname=%s host=%s port=%s user=%s password=%s sslmode=%s",
		config.Database, config.Host, config.Port, config.Username, config.Password, config.SslMode,
	)

	db, err := sql.Open("postgres", conf)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")
	return db
}
