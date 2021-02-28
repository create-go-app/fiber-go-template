package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/create-go-app/fiber-go-template/app/queries"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries // load queries from User model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Switch types of database connection.
	switch os.Getenv("DATABASE_TYPE") {
	case "sqlite":
		// SQL lite connection logic here.
	case "postgres":
		// Open DB connection for PostgreSQL.
		db, err := sqlx.Connect("pgx", os.Getenv("POSTGRES_SERVER_URL"))
		if err != nil {
			return nil, fmt.Errorf("error connecting to PostgreSQL database, %w", err)
		}

		// Define DB connection settings.
		maxConn, _ := strconv.Atoi(os.Getenv("POSTGRES_MAX_CONNECTIONS"))
		maxIdleConn, _ := strconv.Atoi(os.Getenv("POSTGRES_MAX_IDLE_CONNECTIONS"))
		maxLifetimeConn, _ := strconv.Atoi(os.Getenv("POSTGRES_MAX_LIFETIME_CONNECTIONS"))

		// Set DB connection settings.
		db.SetMaxOpenConns(maxConn)                           // the default is 0 (unlimited)
		db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns = 2
		db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // 0, connections are reused forever

		// Try to ping DB.
		if err := db.Ping(); err != nil {
			_ = db.Close() // close database connection
			return nil, fmt.Errorf("error while send ping to PostgreSQL database, %w", err)
		}

		return &Queries{
			UserQueries: &queries.UserQueries{DB: db}, // return queries for User model
		}, nil
	}

	return nil, fmt.Errorf("error, when connecting to any database")
}
