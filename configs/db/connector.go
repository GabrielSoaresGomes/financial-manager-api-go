package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	connection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = connection.Ping()
	if err != nil {
		connectionCloseError := connection.Close()
		if connectionCloseError != nil {
			return nil, err
		}
		return nil, err
	}

	return connection, nil
}
