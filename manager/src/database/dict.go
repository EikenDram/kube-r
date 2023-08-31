package database

import "database/sql"

type RequestOperation struct {
	IdOperation int64  `field:"ID_OPERATION"`
	Operation   string `field:"OPERATION"`
	Description string `field:"DESC"`
}

type RequestOperationModel struct {
	DB *sql.DB
}

type RequestState struct {
	IdState     int64  `field:"ID_STATE"`
	State       string `field:"STATE"`
	Description string `field:"DESC"`
}

type RequestStateModel struct {
	DB *sql.DB
}
