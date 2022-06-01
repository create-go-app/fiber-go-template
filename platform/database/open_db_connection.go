package database

import (
	"os"

	"github.com/create-go-app/fiber-go-template/app/queries"
	"github.com/jmoiron/sqlx"
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries // load queries from User model
	*queries.BookQueries // load queries from Book model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new Database connection.
	connect := os.Getenv("DB_TYPE")

	var db *sqlx.DB
	var err error
	if connect == "pgx" {
		db, err = PostgreSQLConnection()
	}

	if connect == "mysql" {
		db, err = MysqlConnection()
	}

	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		UserQueries: &queries.UserQueries{DB: db}, // from User model
		BookQueries: &queries.BookQueries{DB: db}, // from Book model
	}, nil
}
