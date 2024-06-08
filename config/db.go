package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToPostgresDB() *sql.DB {
	const (
		host     = "localhost"
		port     = 4321
		user     = "postgres"
		password = "secret"
		dbname   = "bob"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %s", err)
	}

	fmt.Println("Successfully connected to the database!")
	return db
}
