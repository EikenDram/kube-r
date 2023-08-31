package database

import "database/sql"

type Request struct {
	IdRequest int64  `field:"ID_REQUEST"`
	IdState   int64  `field:"ID_STATE"`
	IdApp     int64  `field:"ID_APP"`
	IdUser    int64  `field:"ID_USER"`
	Path      string `field:"PATH"`
	Payload   string `field:"PAYLOAD"`
}

// Create a custom Model type which wraps the sql.DB connection pool.
type RequestModel struct {
	DB *sql.DB
}

// Use a method on the custom Model type to run the SQL query.
func (m RequestModel) All() ([]Request, error) {
	rows, err := m.DB.Query("SELECT * FROM REQUEST.REQUESTS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Request

	for rows.Next() {
		var r Request

		err := rows.Scan(
			&r.IdRequest,
			&r.IdState,
			&r.IdApp,
			&r.IdUser,
			&r.Path,
			&r.Payload,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, r)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
