package database

import "database/sql"

type User struct {
	IdUser int64 `field:"ID_USER"`
}

// Create a custom Model type which wraps the sql.DB connection pool.
type UserModel struct {
	DB *sql.DB
}
