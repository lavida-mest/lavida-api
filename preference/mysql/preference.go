package mysql

import (
	"database/sql"
	"log"

	"github.com/muathendirangu/lavida-api/preference"
)

type repository struct {
	conn *sql.DB
}

//New creates a new instance of preferences mysql repository
func New(conn *sql.DB) preference.Repository {
	return &repository{
		conn: conn,
	}
}

func (r *repository) Add(preference *preference.Preference) error {
	query := `INSERT INTO preference (trip_id, tour_guide_id) VALUES(?,?)`
	stmt, err := r.conn.Prepare(query)
	if err != nil {
		log.Fatalf("an error occurred while preparing the query: %v", err)
	}
	res, err := stmt.Exec(preference.Trip, preference.Guide)
	if err != nil {
		log.Fatalf("An error occured while trying to insert preferences to database %v", err)
	}
	lastPreference, err := res.LastInsertId()
	preference.ID = lastPreference
	return nil
}
