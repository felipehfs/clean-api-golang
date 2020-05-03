package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	database = os.Getenv("PG_DATABASE")
	user     = os.Getenv("PG_USER")
	password = os.Getenv("PG_PASSWORD")
	host     = os.Getenv("PG_HOST")
	port     = os.Getenv("PG_PORT")
)

var conn *sql.DB

// GetPostgresInstance instantiates the database
func GetPostgresInstance() (*sql.DB, error) {
	if conn == nil {
		uri := fmt.Sprintf("host=%v port=%v user=%v "+
			"password=%v dbname=%v sslmode=disable",
			host, port, user, password, database)

		db, err := sql.Open("postgres", uri)
		if err != nil {
			return nil, err
		}

		conn = db

		return conn, nil
	}

	return conn, nil
}
