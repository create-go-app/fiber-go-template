package database

import (
	"fmt"
	"os"

	"github.com/create-go-app/fiber-go-template/app/queries"
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries // load queries from User model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Switch types of database connection.
	switch os.Getenv("DB_TYPE") {
	case "sqlite":
		// SQL lite connection logic here.
	case "mongo":
		// MongoDB connection logic here.
	case "postgres":
		// Define a new PostgreSQL connection.
		db, err := PostgreSQLConnection()
		if err != nil {
			return nil, err
		}

		return &Queries{
			// Set queries from models:
			UserQueries: &queries.UserQueries{DB: db}, // from user model
		}, nil
	}

	return nil, fmt.Errorf("error, when connecting to database (check DB_TYPE)")
}
