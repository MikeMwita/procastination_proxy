package database

import (
	"github.com/MikeMwita/procastination_proxy/internal/core/services"
)

// Queries struct for collect all app services.
type Queries struct {
	*services.UserQueries // load services from User model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set services from models:
		UserQueries: &services.UserQueries{DB: db}, // from User model
	}, nil
}
