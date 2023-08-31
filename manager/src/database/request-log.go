package database

import (
	"database/sql"
	"time"
)

type RequestLog struct {
	Id          int64     `field:"ID"`
	IdRequest   int64     `field:"ID_REQUEST"`
	IdOperation int64     `field:"ID_OPERATION"`
	Stamp       time.Time `field:"STAMP"`
	Message     string    `field:"MESSAGE"`
}

// Create a custom Model type which wraps the sql.DB connection pool.
type RequestLogModel struct {
	DB *sql.DB
}
