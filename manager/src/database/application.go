package database

import "database/sql"

type Application struct {
	IdApp int64 `field:"ID_APP"`
}

// Create a custom Model type which wraps the sql.DB connection pool.
type ApplicationModel struct {
	DB *sql.DB
}
